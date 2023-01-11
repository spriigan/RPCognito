FROM alpine:latest

WORKDIR /app

COPY rpcognito /

CMD [ "/rpcognito" ]