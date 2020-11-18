SHELL := /bin/bash

run:
	go run app/sales-api/main.go

tidy:
	go mod tidy
	go mod vendor

hey:
	hey -m GET -c 100 -n 100000 "http://localhost:3000/readiness"

expvarmon:
	expvarmon -ports="4000"