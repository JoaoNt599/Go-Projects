# Search IPs and Servers 

This project is a command line application (CLI) to search for IPs and Servers on the internet using the Go language.  

## Features

- Search IPs from the internet
- Search Server from the internet

## Run the project

1. Install dependencies:

    go mod tidy

2. Search IPs:

    go run cmd/main.go ip --host <site>

3. Search Server:

    go run cmd/main.go servers --host <site>