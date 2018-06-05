# GOuptime
Health checking downtime get notified to Telegram

### HTTP Status (http://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml)
`1xx: Informational` - Request received, continuing process
`2xx: Success` - The action was successfully received, understood, and accepted
`3xx: Redirection` - Further action must be taken in order to complete the request
`4xx: Client Error` - The request contains bad syntax or cannot be fulfilled
`5xx: Server Error` - The server failed to fulfill an apparently valid request

### Requirement
* go 1.8 or latest
* Docker
* scratch (https://docs.docker.com/samples/library/scratch/)
* SSL Certificates (https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt)
  - `curl -o ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt`

### Usage
1. Custom update file `.env.example` to `.env`
2. Change file `url.example` to `url`
3. `bash start.sh {{name_service}}`

### Maintainer
fajarhide (https://github.com/fajarhide)
