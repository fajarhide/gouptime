# gouptime
Monitoring uptime golang notify to Telegram

### Requirement

* go 1.8 or latest
* Docker
* scratch (https://docs.docker.com/samples/library/scratch/)
* SSL Certificates (https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt)
  - `curl -o ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt`

### Dependency Module
`go get gopkg.in/tucnak/telebot.v2`

### Usage
1. `bash start.sh {{name_service}}`
2. `/start` Starting Monitoring
3. `/stop` Stopping Monitoring for checking issue

### Maintainer
fajarhide (https://github.com/fajarhide)
