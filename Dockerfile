FROM golang

WORKDIR /go/src/app
COPY ./app /go/src/app

RUN go get ./
RUN go build
	
EXPOSE 9911