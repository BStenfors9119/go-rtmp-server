# Minimal Go RTMP server

This is a very minimal RTMP server that I have made for my own personal use. I would not recommend using it if you do not understand the source code.

## Build & Install

[Go](https://golang.org) is required for building and installing.

```bash
go get github.com/catsocks/go-rtmp-server
```

## Usage

```bash
go-rtmp-server -key changeMe
```

A stream key is optional (not recommended), and if no address is provided with the `-addr` flag the server will listen at all addresses in the default RTMP **1935** port.

You may use [OBS](https://OBSproject.com) to stream using e.g the URL ```rtmp://localhost/?key=changeMe``` (leaving the _Stream key_ field empty), and [VLC](https://www.videolan.org/vlc/index.html) to watch it by going over to the menu, clicking Media > Open Network Stream..., and pasting e.g ```rtmp://localhost``` in the URL field.
