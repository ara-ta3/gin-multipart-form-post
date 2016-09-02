.PHONY: gin/main simple/main

install:
	go get "github.com/gin-gonic/gin"

run-gin: gin/main
	./$<

run-simple: simple/main
	./$<

gin/main:
	go build -o $@ gin/main.go

simple/main:
	go build -o $@ simple/main.go
