# FIRST STAGE
# BUILD STAGE
#  AS is the first stage for multistage dockerfile is to build a binary file
#  this is to reduce the size
FROM golang:1.20-alpine3.18 AS builder

ENV GOPATH /go
#  to declare the current working directory inside the image
WORKDIR /app
#  copy all the necessary files, the first dot copies everything
#  from the current folder where we run the docker build command
#  to build the image
#  the second dot is the current working directory inside the image
#  where the files and folder are being copied to, in our case as we
#  specified before it will be the /app it will be the place we store data
COPY . .

# we build our app to a single binanry executable file
#  o stands for output 
RUN go build -o main main.go
# we run it to download and extract the migrate binary

RUN apk add --no-cache wget
RUN wget -q https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz
RUN tar -xzf migrate.linux-amd64.tar.gz
RUN mkdir -p ~/bin
RUN mv migrate.linux-amd64 ~/bin/migrate
RUN rm migrate.linux-amd64.tar.gz
# Set executable permissions for the migrate binary
RUN chmod +x ~/bin/migrate

# SECOND STAGE
# RUN STAGE
FROM alpine:3.18
WORKDIR /app
# we use the from argument to tell docker where to copy the file from
# then the path to the file we want to copy
# the dot represents the WORKDIR that we set above
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
# we will copy from the builder the downloaded migrate bianry to the final image
COPY --from=builder /root/bin/migrate ./migrate
COPY .env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration
# doesn't actually publish the port it only functions as a documentation
# between the person who builds the image and the person who runs the container
# about which ports to be published
EXPOSE 3000
# define the default command when the container starts
# the CMD is an array of command-line arguments, in this instance
# we just need to run the executable file which we built in the previous step
# /app/main
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]