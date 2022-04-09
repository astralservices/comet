FROM golang:1.18-alpine

ENV COMET_PORT=8080
ENV COMET_ROOT=/files

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /comet

EXPOSE $COMET_PORT

CMD [ "/comet" ]