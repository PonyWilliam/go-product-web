FROM alpine
ADD ProductWeb-service /ProductWeb-service
ENTRYPOINT [ "/ProductWeb-service" ]
