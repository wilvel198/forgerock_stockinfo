# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV FR_Stock_API_Key=C227WD9W3LUVKVV9
ENV FR_Stock_Symbol=IBM
ENV FR_Stock_Days=20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /forgerock_stockinfo

EXPOSE 8080

CMD [ "/forgerock_stockinfo" ]