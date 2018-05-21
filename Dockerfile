FROM scratch
MAINTAINER Fajarhide "fajarhide@gmail.com"
ADD ca-certificates.crt /etc/ssl/certs/
ADD url.txt /
ADD app /
CMD ["/app"]
