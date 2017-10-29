all: build clean
build: compile 
	docker build -t swarmpit/swarmpit-ec .
compile:
	GOOS=linux go build .
clean:
	rm -rf swarmpit-ec
