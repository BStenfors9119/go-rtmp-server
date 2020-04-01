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

To publish your stream to the server, the cross-platform [OBS Studio](https://OBSproject.com) can be used by pasting "rtmp://127.0.0.1" in the Server field of the Stream tab in the settings window, and the stream key in the Stream Key field, assuming you are running the server in the same computer that OBS is running and using the default address.

To play the stream, the cross-plaform [VLC media player](https://www.videolan.org/vlc) can be used by pasting "rtmp://127.0.0.1/Really%20strong%20password", or just "rtmp://127.0.0.1" if using no password, in the network URL field under the Network tab of the Open Media window that can be accessed by clicking Media > Open Network Stream in the menu bar.

Note that you may have to forward the RTMP port in your router to the computer which the server is running so that others outside your local network may be able to connect and watch the stream. And don't rely on the viewer's password alone if streaming sensitive information, consider sharing a virtual private network with who you want to stream to, passing your address in the private network to the `-addr` to avoid the server listening at other addresses.

## Install

With [Go](https://golang.org) >= 1.6 installed you can build and install with the following command:

```$ go get github.com/catsocks/go-rtmp-server```

Alternatively, executables are available under the releases tab of this project's Github repository page.
