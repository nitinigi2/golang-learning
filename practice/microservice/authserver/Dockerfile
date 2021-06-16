# base image for our application
# this wil pull docker image from docker hub
FROM golang

# copy root dir to /app/auth in docker image
COPY . /app/auth

# setting work directory, so all the subsequent commands can be run on this dir
WORKDIR /app/auth

# build authserver code and set it's output binary in dir as /out/auth
RUN go build -o ./out/auth .

# command to execute binaries
CMD ["./out/auth"]