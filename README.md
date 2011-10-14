# Negotiation

## Client

To compile:

    8g client.go && 8l -o client client.8

## Server

To compile:

    8g server.go server-model.go arith.go
    8l -o server server.8
