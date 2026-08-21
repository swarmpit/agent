package task

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/docker/docker/client"
)

// statsServer fakes the docker daemon's container stats endpoint.
func statsServer(t *testing.T, body string) (*client.Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/stats") {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Ostype", "linux")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	cli, err := client.NewClientWithOpts(client.WithHost(srv.URL), client.WithVersion("1.40"))
	if err != nil {
		t.Fatalf("cannot build docker client: %s", err)
	}
	return cli, srv.Close
}

// A container that exits while its stats are being read yields a body that does
// not decode into StatsJSON. That used to leave a nil *types.StatsJSON and panic
// the whole agent on the next dereference. See issue #736.
func TestContainerUsageDoesNotPanicOnUndecodableBody(t *testing.T) {
	for name, body := range map[string]string{
		"empty":           "",
		"null":            "null",
		"truncated":       `{"cpu_stats": `,
		"not an object":   `"gone"`,
		"empty object":    `{}`,
		"unexpected html": `<html>error</html>`,
	} {
		t.Run(name, func(t *testing.T) {
			cli, stop := statsServer(t, body)
			defer stop()

			status, ok := ContainerUsage(cli, "deadbeef")

			if name == "empty object" {
				if !ok {
					t.Fatalf("an empty stats object should still be reported")
				}
				return
			}
			if ok {
				t.Fatalf("expected ok=false for %q body, got status %+v", name, status)
			}
			if status.ID != "" || status.Name != "" {
				t.Fatalf("expected a zero status for %q body, got %+v", name, status)
			}
		})
	}
}

func TestContainerUsageReadsStats(t *testing.T) {
	body := `{"name":"/web","id":"deadbeef",
	          "cpu_stats":{"cpu_usage":{"total_usage":200,"percpu_usage":[100,100]},"system_cpu_usage":2000,"online_cpus":2},
	          "precpu_stats":{"cpu_usage":{"total_usage":100},"system_cpu_usage":1000},
	          "memory_stats":{"usage":2048,"limit":4096,"stats":{"cache":0}}}`
	cli, stop := statsServer(t, body)
	defer stop()

	status, ok := ContainerUsage(cli, "deadbeef")
	if !ok {
		t.Fatal("expected the stats to be read")
	}
	if status.Name != "/web" || status.ID != "deadbeef" {
		t.Fatalf("unexpected identity: %+v", status)
	}
	if status.MemoryLimit != 4096 {
		t.Fatalf("expected memory limit 4096, got %v", status.MemoryLimit)
	}
	if status.CPUPercentage <= 0 {
		t.Fatalf("expected a positive cpu percentage, got %v", status.CPUPercentage)
	}
}

// One bad container must not take down collection for the others, and must not
// kill the process: the goroutine in ContainersUsage now recovers.
func TestContainersUsageSurvivesABadContainer(t *testing.T) {
	stats := `{"name":"/ok","id":"aaa","cpu_stats":{"cpu_usage":{"total_usage":200},"system_cpu_usage":2000,"online_cpus":1},"precpu_stats":{"cpu_usage":{"total_usage":100},"system_cpu_usage":1000},"memory_stats":{"usage":1,"limit":2}}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/containers/json"):
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{"Id":"aaa"},{"Id":"bbb"}]`))
		case strings.Contains(r.URL.Path, "aaa/stats"):
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Ostype", "linux")
			w.Write([]byte(stats))
		default:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Ostype", "linux")
			w.Write([]byte("null"))
		}
	}))
	defer srv.Close()

	cli, err := client.NewClientWithOpts(client.WithHost(srv.URL), client.WithVersion("1.40"))
	if err != nil {
		t.Fatalf("cannot build docker client: %s", err)
	}

	got := ContainersUsage(cli)
	if len(got) != 1 {
		t.Fatalf("expected only the healthy container to be reported, got %+v", got)
	}
	if got[0].ID != "aaa" {
		t.Fatalf("expected container aaa, got %+v", got[0])
	}
}
