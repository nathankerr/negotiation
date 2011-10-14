# Negotiation

Currently the client address has to be set matching the protocol used. This is because I don't know how to get all the protocols to work on the same port.

Client support for TLS is broken (due to an API change that I haven't tracked down yet.

## Client

To compile:

    8g client.go && 8l -o client client.8

## Server

To compile:

    8g server.go server-model.go arith.go
    8l -o server server.8
