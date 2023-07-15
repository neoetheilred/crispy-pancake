FROM golang:alpine as build

WORKDIR /app/
RUN apk add --no-cache ca-certificates
COPY . .
RUN go build -o main ./cmd/app/main.go

FROM alpine:latest
COPY --from=build /app/main ./
COPY --from=build /app/assets/pages/* ./assets/pages/

CMD [ "./main" ]