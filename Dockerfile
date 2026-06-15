FROM golang:1.27.0-alpine3.24 AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd cmd
COPY internal internal
RUN go build -o /bin/notes_server cmd/server/main.go

FROM base AS test
CMD [ "go", "test", "./..." ]

FROM alpine:3.24 AS prod
COPY --from=base /bin/notes_server /bin/notes_server
CMD ["/bin/notes_server"]
