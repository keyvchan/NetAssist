# README

## Introduction

Easily setup server and client for multiple protocols.

## Usage

### General

```shell
$ NetAssist <type> <protocol> <address:port>
```

### Send data as bytes in TCP

```shell
$ NetAssist <type> <protocol> <address:port> --binary
```

Above will send data as bytes, remove `--binary` to send as string.

## Protocol

- TCP (multiple clients, broadcast to all clients)
- UDP
- Unix Domain Socket (except for `unixpacket`)
