FROM golang:scratch

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get "github.com/google/uuid"
RUN go get "github.com/gorilla/mux"
RUN go get https://github.com/Carson-McCombs/Fetch-BackEnd-Test/main

RUN cd /build && cd clone https://github.com/Carson-McCombs/Fetch-BackEnd-Test.git


#RUN cd/build/HTTP-SERVER/main && go build
EXPOSE 8080

ENTRYPOINT [ "/build/HTTP-SERVER/main" ]