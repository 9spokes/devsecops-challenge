FROM golang:latest AS build
WORKDIR /9spokes
COPY hello.go .
RUN go mod init hello
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello .

FROM scratch
COPY --from=build /9spokes/hello .
EXPOSE 8080
CMD ["./hello"]
