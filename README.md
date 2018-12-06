# Minimal Go RTMP server

This is a very minimal RTMP server that I have made for my own personal use. I would not recommend using it if you do not understand the source code.

## Build

[Go](https://golang.org) is required for building.

```bash
go get github.com/catsocks/go-rtmp-server
```

## Usage

```bash
go-rtmp-server -key changeMe
```

The stream key is set with the `key` flag, a viewer's password can be set with the `pass` flag, and the server address can be set with `addr`. By default the server will listen at the RTMP port **1935**.

You may use [OBS](https://OBSproject.com) to stream using the URL ```rtmp://localhost/?key=changeMe``` (leaving the _Stream key_ field empty), and [VLC](https://www.videolan.org/vlc/index.html) to watch it by going over to the menu, clicking Media > Open Network Stream..., and pasting ```rtmp://localhost/?pass=abcd``` in the URL field (omit the _pass_ query string if you did not set a password).
