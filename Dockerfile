FROM golang:1.20 AS gobuilder
WORKDIR /app
COPY pkg/ pkg/
COPY cmd/cartaman/main.go /app/main.go
COPY .env /app/.env
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN CGO_ENABLED=0 GOOS=linux go build -o bin main.go

FROM alpine:3.15.1 as cartaman-app
COPY --from=gobuilder /app/bin bin
COPY --from=gobuilder /app/pkg/adapters/repository/pg/migrations/ migrations/
COPY --from=gobuilder /app/.env .env
CMD ["bin"]