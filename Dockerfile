FROM golang:1.18-alpine as build

WORKDIR /kielba

#COPY go.mod go.sum ./
#RUN go mod download

COPY . .

RUN go build -o kielba *.go


FROM alpine:latest as production
RUN apk add --no-cache ca-certificates

COPY --from=build kielba .

EXPOSE 80
CMD ["./kielba"]