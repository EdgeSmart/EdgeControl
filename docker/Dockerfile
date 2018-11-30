FROM debian:buster-slim
ADD EdgeControl /
ADD conf/* /conf
ADD templates/* /templates
ENTRYPOINT ["/EdgeControl"]
EXPOSE 8000
