# builder image
FROM golang:1.16-alpine as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o /golang-gethello cmd/*.go

# generate clean, final image for end users
FROM alpine:3.11.3
WORKDIR /

COPY --from=builder /golang-gethello .

# executable
EXPOSE 8777
CMD [ "./golang-gethello" ]