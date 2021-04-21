FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD ProductWeb /app/ProductWeb

CMD ["./ProductWeb"]