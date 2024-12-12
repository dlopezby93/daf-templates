FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .

RUN (cd cmd && GOOS=linux GOARCH=arm64 go build -o /bin/bootstrap main.go)

FROM scratch
WORKDIR /app

COPY --from=build /bin/bootstrap /go/bin/bootstrap
COPY --from=build /go/src/app/cmd/.env /app

ENTRYPOINT ["/go/bin/bootstrap"]