FROM golang:latest

# RUN mkdir /app
# ADD ./server /app
# WORKDIR /app

# RUN go get -d -v 
# RUN go build -v
# RUN echo $PATH

# CMD ["/app/main"]


WORKDIR /app
COPY ./server/go.mod .
COPY ./server/go.sum .
RUN go mod download
COPY ./server .
RUN go build

CMD ["./server"]




# RUN go mod download
# RUN go build -v
# RUN echo $PATH
# RUN ls
# RUN pwd

# CMD ["./docker-golang-example"]




