#FROM golang:alpine

#RUN mkdir /build
#WORKDIR /build

#RUN export GO111MODULE=on
#RUN go get "github.com/google/uuid"
#RUN go get "github.com/gorilla/mux"
#RUN go get https://github.com/Carson-McCombs/Fetch-BackEnd-Test/HTTP-SERVER/main

#RUN cd /build && cd clone https://github.com/Carson-McCombs/Fetch-BackEnd-Test.git


#RUN cd/build/HTTP-SERVER/main && go build
#EXPOSE 8080

#ENTRYPOINT [ "/build/HTTP-SERVER/main" ]

############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /build
WORKDIR /build
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o /HTTP-SERVER/main
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /HTTP-SERVER/main /HTTP-SERVER/main
# Run the hello binary.

EXPOSE 8080
ENTRYPOINT ["/HTTP-SERVER/main"]