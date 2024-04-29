# syntax=docker/dockerfile:1

FROM golang:1.22.2-alpine3.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

RUN addgroup -g 10014 choreo && \
    adduser  --disabled-password --uid 10014 --ingroup choreo choreouser

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-sample-app

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

USER 10014

# Run
CMD ["ALL_PROXY=socks5://localhost:1055/ HTTP_PROXY=socks5://127.0.0.1:1080/ /docker-sample-app"]
