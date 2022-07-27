package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format"
	"github.com/nareix/joy4/format/rtmp"
)

var (
	addr = flag.String("addr", ":1935", "server address")
	key  = flag.String("key", "", "stream key for streaming to the server")
	pass = flag.String("pass", "", "password for watching the stream")
)

var que *pubsub.Queue

func main() {
	flag.Parse()

	format.RegisterAll()

	server := &rtmp.Server{
		Addr:          *addr,
		HandlePublish: HandlePublish,
		HandlePlay:    handlePlay,
	}

	if *key == "" {
		fmt.Println("Warning: A stream key was not set and anyone can publish a stream to this server.")
	} else {
		fmt.Printf("Info: Your stream key is %q. Don't let anyone see it!\n", *key)
	}
	if *pass == "" {
		fmt.Println("Warning: A viewer's password was not set and anyone can watch the stream.")
	} else {
		fmt.Println("Info: The viewer's password should be added to the end of the URL for this server like so:",
			"rtmp://127.0.0.1/"+url.PathEscape(*pass))
	}

	fmt.Println("Info: Starting the stream server at", *addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, "Fatal: Couldn't run the stream server:", err)
	}
}

func HandlePublish(conn *rtmp.Conn) {
	defer conn.Close()

	// Require the streamer to append the stream key to the end of the
	// URL for the server, if the stream key isn't empty.
	// The popular program OBS Studio will append the value of the Stream Key
	// to the end of the URL provided in the Server field.
	fmt.Sprintf("Info: Conn URL Path %s", conn.URL.Path)
	if *key != "" && *key != strings.TrimPrefix(conn.URL.Path, "/") {
		fmt.Println("Info: The wrong stream key was used to stream to the server.")
		return
	}

	if que != nil {
		que.Close()
	}
	que = pubsub.NewQueue()

	streams, err := conn.Streams()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Couldn't stream:", err)
		return
	}
	que.WriteHeader(streams)

	fmt.Println("Info: The server has started streaming.")

	if err := avutil.CopyPackets(que, conn); err == io.EOF {
		fmt.Println("Info: The server has stopped streaming.")
	} else if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Couldn't stream:", err)
	}
}

func handlePlay(conn *rtmp.Conn) {
	defer conn.Close()
	if que == nil {
		return
	}

	// Require the viewer to append the viewer's password to the end of the URL
	// for the server, if the viewer's password isn't empty.
	if *pass != "" {
		p := strings.TrimPrefix(conn.URL.Path, "/")
		if p == "" {
			fmt.Println("Info: A viewer tried to watch the stream providing no password.")
			return
		}

		if *pass != p {
			fmt.Println("Info: A viewer tried to watch the stream with the wrong password.")
			return
		}
	}

	if err := avutil.CopyFile(conn, que.Latest()); err != nil && err != io.EOF {
		fmt.Printf("%+v\n", err)
		fmt.Println("Info: Couldn't serve the stream to a viewer:", err)
	}
}
