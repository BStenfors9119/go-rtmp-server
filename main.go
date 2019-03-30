package main

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format"
	"github.com/nareix/joy4/format/rtmp"
)

var addr = flag.String("addr", ":1935", "Address of the stream server")
var pass = flag.String("pass", "", "Password for watching the stream")
var key = flag.String("key", "", "Key for streaming to the server")

func main() {
	flag.Parse()

	format.RegisterAll()

	var que *pubsub.Queue

	server := &rtmp.Server{
		Addr: *addr,
		HandlePlay: func(conn *rtmp.Conn) {
			defer conn.Close()
			if que == nil {
				return
			}

			if *pass != "" && *pass != strings.TrimPrefix(conn.URL.Path, "/") {
				fmt.Println("The wrong password was used to watch the stream")
				return
			}

			if err := avutil.CopyFile(conn, que.Latest()); err != nil && err != io.EOF {
				fmt.Println("Unable to serve stream:", err)
			}
		},
		HandlePublish: func(conn *rtmp.Conn) {
			defer conn.Close()

			if *key != "" && *key != strings.TrimPrefix(conn.URL.Path, "/") {
				fmt.Println("The wrong stream key was used")
				return
			}

			if que != nil {
				que.Close()
			}
			que = pubsub.NewQueue()

			streams, err := conn.Streams()
			if err != nil {
				fmt.Println("Unable to stream:", err)
				return
			}
			que.WriteHeader(streams)

			fmt.Println("Starting to stream")

			if err := avutil.CopyPackets(que, conn); err == io.EOF {
				fmt.Println("Stopped streaming")
			} else if err != nil {
				fmt.Println("Unable to stream: ", err)
			}
		},
	}

	fmt.Println("Starting the streaming server at", *addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Unable run the stream server:", err)
	}
}
