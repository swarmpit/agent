all: build clean
build: compile 
	docker build -t swarmpit/event-collector .
compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
clean:
	rm -rf event-collector
