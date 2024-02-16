FROM golang:alpine as build 
RUN apk add --no-cache ca-certificates 

WORKDIR /go/src/github.com/RapidCodeLab/vast-provider
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags '-extldflags "-static"' -o app ./main

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt \
     /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/src/github.com/RapidCodeLab/vast-provider/app /app


ENTRYPOINT ["/app"]
