FROM golang:latest
FROM scratch
MAINTAINER Fajarhide "fajarhide@gmail.com"
ADD ca-certificates.crt /etc/ssl/certs/
ADD url /
ADD .env /
ADD app /
CMD ["/app"]
