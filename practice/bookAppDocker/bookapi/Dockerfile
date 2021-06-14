FROM golang
COPY . /app/bookapi
WORKDIR /app/bookapi
RUN go build -o ./out/bookapi .
CMD ["./out/bookapi"]