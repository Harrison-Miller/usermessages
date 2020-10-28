# Using a multi-stage build so that we can get rid of the bloat from the source + golang image after we have the binary
FROM golang:1.12-alpine AS build

# A cert bundle is needed so the server can verify https://reqres.in/api/users
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD . .

RUN GO111MODULE=on GOARCH=386 CGO_ENABLED=0 GOOS=linux go build -o usermessages .

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/usermessages /

EXPOSE 8080

ENV DATA_DIR="/data"

CMD ["/usermessages"]