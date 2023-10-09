FROM golang:1.20.4-alpine3.17

WORKDIR /collabora/wopi

RUN go install github.com/cosmtrek/air@latest

COPY wopi/go.mod wopi/.air.toml ./

RUN go mod download

CMD [ "air", "-c", ".air.toml" ]
