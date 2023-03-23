# Specifies the base image we're extending
FROM golang:1.20
# create a directory for the app

# Set the Current Working Directory inside the container
WORKDIR /lms

# copy the source code as the last step

COPY . .

# download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# build the binary

RUN go build -o lms .

# our start command which kicks off

EXPOSE 8080

CMD [ "/lms" ]
