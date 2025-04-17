FROM debian:bookworm-slim

WORKDIR /app
COPY app /app/app

ENTRYPOINT ["/app/app"]