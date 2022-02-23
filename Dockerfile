# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

# COPY requirements.txt requirements.txt
# RUN pip3 install -r requirements.txt
COPY . .
CMD ["go", "run", "main.go"]