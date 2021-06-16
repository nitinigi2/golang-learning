FROM golang
COPY . /app/auth
WORKDIR /app/auth
RUN go build -o ./out/auth .
CMD ["./out/auth"]