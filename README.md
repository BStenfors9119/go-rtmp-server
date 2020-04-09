# Go RTMP server

A RTMP server with a minimal command-line interface written in Go that I made for my own personal use as a local streaming server to privately stream to my friends.

## Usage

```
Usage of go-rtmp-server:
  -addr string
        server address (default ":1935")
  -key string
        stream key for streaming to the server
  -pass string
        password for watching the stream
```

To start the stream server at the default RTMP port 1935 with no stream key or password:

```
$ go-rtmp-server
Warning: A stream key was not set and anyone can publish a stream to this server.
Warning: A viewer's password was not set and anyone can watch the stream.
Info: Starting the stream server at :1935
```

To start the stream server with a stream key and password (recommended):

```
$ go-rtmp-server -key "Really strong key" -pass "Really strong password"
Info: Your stream key is "Really strong key". Don't let anyone see it!
Info: The viewer's password should be added to the end of the URL for this server like so: rtmp://127.0.0.1/Really%20strong%20password
Info: Starting the stream server at :1935
```

The stream key and password are passed as arguments for the sake of ease of use but at the cost of security, so don't use them anywhere else, and because there isn't a limit to the attempts someone can make to publish or play the stream with the wrong stream key and password, make sure you make them strong.

### Publish your stream

The cross-platform [OBS Studio](https://OBSproject.com) can be used to publish a stream by pasting rtmp://127.0.0.1 in the Server field of the Stream tab in the settings window, and the stream key in the Stream Key field, assuming the server is running using the default address in the same computer.

### Play the stream

The cross-plaform [VLC media player](https://www.videolan.org/vlc) can be used to watch the stream by pasting rtmp://127.0.0.1/Really%20strong%20password, or just rtmp://127.0.0.1 if using no password, in the network URL field under the Network tab of the Open Media window that can be accessed by clicking Media > Open Network Stream in the menu bar, assuming the server is running using the default address in the same computer.

The RTMP port in your router may have to be forwarded to the computer which the server is running so that connections can be made from outside your local network for publishing a stream or watching the stream.

## Install

You can download executables for Windows and GNU/Linux under the "releases" tab of this project's Github repository page.

Alternatively, with [Go](https://golang.org) >= 1.6 installed you can build and install with the following command:

```$ go get github.com/catsocks/go-rtmp-server```
