FROM scratch
MAINTAINER Pavol Noha <pavol.noha@gmail.com>
WORKDIR /
COPY event-collector /
ENTRYPOINT ["./event-collector"]
