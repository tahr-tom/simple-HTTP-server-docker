FROM golang:latest

# create a build directory
WORKDIR /build

# get the dependency
RUN go get -u github.com/go-chi/chi

# copy the files to container
COPY . .

# build the program
RUN go build -o main .

# create a working directory
WORKDIR /out

# copy the binary to /out
RUN cp /build/main .

EXPOSE 80

CMD ["/out/main"]