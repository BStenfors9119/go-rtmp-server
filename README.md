# Go RTMP server

A RTMP server with a minimal command-line interface written in Go that I made
for my own personal use as a local streaming server to privately stream to my
friends.

## Usage

    Usage of go-rtmp-server:
      -addr string
            server address (default ":1935")
      -key string
            stream key for streaming to the server
      -pass string
            password for watching the stream

Start the server at the default RTMP port 1935 with a stream key and password by
running the server with the -key and -pass flags like so:

    $ go-rtmp-server -key "Really strong key" -pass "Really strong password"
    Info: Your stream key is "Really strong key". Don't let anyone see it!
    Info: The viewer's password should be added to the end of the URL for this server like so: rtmp://127.0.0.1/Really%20strong%20password
    Info: Starting the stream server at :1935

You can omit the -key and -pass flags to run without a stream key and password:

    $ go-rtmp-server
    Warning: A stream key was not set and anyone can publish a stream to this server.
    Warning: A viewer's password was not set and anyone can watch the stream.
    Info: Starting the stream server at :1935

The stream key and password are passed to the server as flags for ease of use,
but that can unintentionally expose them if your shell saves your commands to a
file, for example. Additionally, there is no protection against brute-forcing a
stream key or password.

## Install

You can download the executable from the Releases tab in the GitHub repository
page of this project.

And you can download and install with [Go](https://golang.org) >= 1.6 using the
_go get_ command as follows:

```sh
$ go get github.com/catsocks/go-rtmp-server
```

## Publish your stream

You can publish your stream to the server using
[OBS Studio](https://OBSproject.com) by opening the _Settings_ window, clicking
the _Stream_ tab, and entering _rtmp://_ proceeded by the address of the
computer which the server is running on in the _Server_ field. This field can be
set to _rtmp://127.0.0.1_ if the server is running on the same computer as OBS
Studio.

## Watch the stream

You can watch the stream using [VLC media player](https://www.videolan.org/vlc)
by heading over to the menu bar and clicking _Media_ and _Open Network Stream_,
then pasting the URL containing the URL-encoded password that the server printed
on startup in the URL field if you are watching the stream from the same
computer the server is running on, otherwise you must replace the part of the
URL containing the address with the appropriate address, such as your public
IP.

Note that you may have to configure your router to forward the port that the
server uses to the computer which the server will be run on, so that computers
from outside your local network may reach it.
