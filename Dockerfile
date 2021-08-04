FROM golang:latest

LABEL maintainer="Alireza Saffar <a.khorasany@gmail.com>"

RUN mkdir /app
WORKDIR /app
COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN go mod download
COPY ./backend .

EXPOSE 9006
RUN go build

CMD ["./backend"]