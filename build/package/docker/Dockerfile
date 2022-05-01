FROM golang:1.18-alpine as builder

WORKDIR /build

# Git must be installed because go mod requires it to download dependencies
# After its used we may remove it to reduce docker image size
RUN apk --no-cache add git tzdata gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -tags musl -o b2w-explorer cmd/b2w-explorer/main.go

# Execution container
FROM alpine
RUN apk --no-cache add tzdata
WORKDIR /app
USER 1000
COPY --from=builder /build/b2w-explorer /app/
EXPOSE 8080
CMD ["/app/b2w-explorer"]


# docker build -t b2w-explorer  -f  ./build/package/docker/Dockerfile .
# docker tag b2w-explorer igorralexsander/b2w-explorer
# docker push igorralexsander/b2w-explorer
