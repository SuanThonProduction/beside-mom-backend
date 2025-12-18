FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM gcr.io/distroless/base-debian12
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=admin
ENV DB_NAME=besidemom
ENV SSL_MODE=disable
ENV APP_HOST=0.0.0.0
ENV APP_PORT=5000
ENV CORS=https://www.besidemom.com,https://besidemom.com,http://localhost:3000
ENV BUCKET_NAME=Beside-Mom
ENV EMAIL_HOST=smtp.gmail.com
ENV EMAIL_PORT=587
ENV EMAIL_USER=kasianbot66@gmail.com
COPY --from=builder /app/main /app/main
USER nonroot:nonroot
CMD ["/app/main"]