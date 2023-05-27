FROM golang:1.20-alpine
LABEL authors="shonayevshyngys"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build

EXPOSE 8080

CMD [ "./rating_service" ]