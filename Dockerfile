FROM golang:1.23-alpine

WORKDIR /user/src/app

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

# pre-copy/cache go.mod for better layer caching for subsequent builds
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

CMD ["app"]