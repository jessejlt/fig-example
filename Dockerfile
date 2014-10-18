# Official Golang image
FROM golang:1.3.3

# Install Godep for dependency management
RUN go get github.com/tools/godep

# Add Golang compatible Redis network
ENV REDIS_NETWORK "redis_1:6379"

# Set app path
ENV APP_PATH /go/src/app

# This image adds our code to /go/src/app
WORKDIR $APP_PATH

# Add source
ADD . $APP_PATH

# Buld our service
RUN godep go build -v

# App will listen on 8080
EXPOSE 8080

# Start our service
CMD ["./app"]