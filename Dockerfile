FROM golang:1

WORKDIR app/aristotle

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

EXPOSE 5000
CMD ["aristotle"]
