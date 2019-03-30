# Minimal Go RTMP server

This is a minimal RTMP server that I have made for my own personal use as a local stream server.

## Usage

```bash
go-rtmp-server -addr :1935 -key Muffins -pass Mittens
```

You may set a stream key with the `key` flag, password protect the stream with the `pass` flag, and choose the address the server will listen at with the `addr` flag. By default the server will listen at all hosts at the default RTMP port 1935.

In [OBS](https://OBSproject.com) you may stream by pasting ```rtmp://127.0.0.1``` in the Server field in the Output tab in settings and pasting your stream key in the Stream key field.

In [VLC](https://www.videolan.org/vlc) you may watch the stream by going over to the menu and clicking Media > Open Network Stream and pasting `rtmp://127.0.0.1/` plus your your password in the network URL field.

## Build

Install [Go](https://golang.org) and run `go get github.com/catsocks/go-rtmp-server` to download the project to your GOPATH, retrieve all the dependencies, build, and install it.

## TODO

* GUI
* Encryption