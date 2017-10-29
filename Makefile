all: build clean
build: compile 
	docker build -t swarmpit/swarmpit-ec .
compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
clean:
	rm -rf swarmpit-ec
