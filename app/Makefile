include .env

export GO111MODULE=on

init: .env

.env:
	cp ./$@.example ./$@

build:
	# env GOOS=linux GOARCH=amd64 go build main.go
	# for raspberry pi
	env GOOS=linux GOARCH=arm go build main.go

gomod:
	go mod init github.com/mizuho1998/grafana_monitor/app
