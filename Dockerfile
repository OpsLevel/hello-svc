FROM golang

WORKDIR /app
COPY . .

RUN go build -o /usr/local/bin/hello-svc .

EXPOSE 8080
CMD ["hello-svc"]
