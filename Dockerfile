FROM golang:1.19.3-alpine AS builder

RUN go version
RUN apk add git

COPY ./ /github.com/riderufa/qrq_etp
WORKDIR /github.com/riderufa/qrq_etp

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/etp/etp.go

#lightweight docker container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/riderufa/qrq_etp/.bin/app .
#COPY --from=0 /github.com/riderufa/qrq_etp/pkg/config/ ./config/
#COPY --from=0 /github.com/riderufa/qrq_etp/templates/ ./templates/

EXPOSE 8005

CMD [ "./app"]
