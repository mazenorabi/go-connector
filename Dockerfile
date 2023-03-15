FROM golang:1.20-bullseye

COPY app /app/

WORKDIR /app

#ENV GOPATH=$GOPATH:/app

RUN go get github.com/gin-gonic/gin gorm.io/gorm

RUN go mod download

RUN go build -v -o /app/bin/htg-connector
#CMD ["app"]

EXPOSE 8080

CMD ["/app/bin/htg-connector"]