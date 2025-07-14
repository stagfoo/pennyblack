#!/bin/bash

go build -o mock/bin/handler_screen handler_screen/main.go;
sleep 1
go build -o mock/bin/handler_toml handler_toml/main.go;
sleep 1
go build -o mock/bin/app app/*.go;

