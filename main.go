package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format"
	"github.com/nareix/joy4/format/rtmp"
)

var addr = flag.String("addr", ":1935", "Address")
var pass = flag.String("pass", "", "Password")
var key = flag.String("key", "changeMe", "Streaming key")

func main() {
	flag.Parse()
	format.RegisterAll()

	server := &rtmp.Server{Addr: *addr}

	var que *pubsub.Queue

	server.HandlePlay = func(conn *rtmp.Conn) {
		if que != nil {
			if *pass == conn.URL.Query().Get("pass") {
				avutil.CopyFile(conn, que.Latest())
			} else {
				fmt.Println("Wrong password.")
			}
		}
		conn.Close()
	}

	server.HandlePublish = func(conn *rtmp.Conn) {
		defer conn.Close()

		streams, err := conn.Streams()
		if err != nil {
			panic(err)
		}

		if *key != conn.URL.Query().Get("key") {
			fmt.Println("The stream was denied due to a stream key mismatch.")
			return
		}

		if que != nil {
			que.Close()
		}
		que = pubsub.NewQueue()
		que.WriteHeader(streams)
		avutil.CopyPackets(que, conn)
	}

	log.Fatal(server.ListenAndServe())
}
