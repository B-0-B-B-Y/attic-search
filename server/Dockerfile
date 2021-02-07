FROM golang AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY ./src ./

# Download dependencies
RUN GOBIN=$GOPATH/bin go get -v

# Build the application
RUN go build -o server .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/server .

# Build a small image
FROM scratch

COPY --from=builder /dist/server /
COPY ./src/database/data.json /database/data.json

# Command to run
ENTRYPOINT ["/server"]