FROM golang:1.23-alpine as build
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /go/bin/app cmd/surl/*.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /
ENTRYPOINT [ "/app" ]
