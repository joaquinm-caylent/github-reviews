FROM golang:1.16.2-alpine AS build
WORKDIR /src
COPY . .
RUN go mod tidy
RUN go build

FROM alpine:3.14
COPY --from=build /src/github-reviews /src/github-reviews
CMD ["/src/github-reviews"]
