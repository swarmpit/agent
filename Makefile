all: build clean
build: compile 
	docker build -t swarmpit/agent .
compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
clean:
	rm -rf agent
