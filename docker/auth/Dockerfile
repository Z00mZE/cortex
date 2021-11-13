################# step 1 #################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /prj
COPY . .
RUN go mod download && go build -o app ./main.go
################# step 2 #################
FROM scratch
COPY --from=builder /prj/app /app
ENTRYPOINT ["/app"]