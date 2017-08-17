FROM golang:onbuild

RUN mkdir /app
ADD . /app 
WORKDIR /app
RUN go test -v
RUN go build -o yadroexe . 
#installing net-tools to run ifconfig command inside docker container
RUN apt-get update
RUN apt-get install net-tools
ENTRYPOINT ["./yadroexe"]

