# OSNAME := $(shell uname -o)
OutName := "culaccino.exe"

all: service

service:
	go build -o $(OutName) ./src/main.go