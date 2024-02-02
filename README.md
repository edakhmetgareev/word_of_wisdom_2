# "Word of wisdom" tcp server 

## Requirements:

Test task for Server engineer (Go)

Design and implement “Word of Wisdom” tcp server.

• TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.

• The choice of the POW algorithm should be explained.

• After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.

• Docker file should be provided both for the server and for the client that solves the POW challenge


## How to start:
`make build` - build docker containers

`make start` - run client

## Some notes:
I didn't write any tests just to save the time. If it is necessary - I can do where it possible.

### Pow algorithm has been chosen as sha256 because:
1) It is used by bitcoin 
2) more secure 
3) more time-consuming (we have to protect our api from ddos, because of that we are interested in time-consuming process of pow)


