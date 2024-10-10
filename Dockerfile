# build binary in a different image
FROM golang:1.23 AS build

COPY main.go go.mod /src/

WORKDIR /src
ENV CGO_ENABLED=0
RUN go mod tidy && go build -o /http-echo-headers .

# create the final image
FROM gcr.io/distroless/static-debian12:nonroot
ENV GOTRACEBACK=single
LABEL name="http-echo-headers" \
    maintainer="abhinav1107" \
    summary="A weserver in Golang that prints all headers that the server gets"

COPY --from=build /http-echo-headers /bin/http-echo-headers
ENTRYPOINT ["http-echo-headers"]
