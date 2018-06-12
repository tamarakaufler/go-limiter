# limiter
Implementation of request rate limiting in Go

## Synopsis
Limiter is a command line implementation exploring rate limiting
with allowed bursting at regular intervals.

Limiting is achieved by blocking on a channel (BurstChan). The channel
is buffered to allow bursting. The channel is refilled at regular intervals,
to provide repeated bursts.

## Input
The CLI accepts the following flags:

  - requests ...... number of requests
  - limit ......... number of miliseconds to limit the requests
  - burst ......... number of requests allowed not to be limited
  - burstrepeat ... how often to repeat the burst (in seconds)

## Usage

go run main.go -requests=40 -limit=400 -burst=4 -burstrepeat=3

docker run \
--name=limiter \
--rm \
quay.io/tamarakaufler/limiter:v1alpha1 \
-requests=40 \
-limit=400 \
-burst=4 \
-burstrepeat=3
