FROM golang:1.16.2-alpine AS build
WORKDIR /src
COPY . .
RUN go mod tidy
RUN go build

FROM scratch
COPY --from=build /src/github-reviewsp /src/github-reviews
CMD ["/src/github-reviews"]
