package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	defaultPort = "1599"
	defaultHost = "localhost"
)

var (
	conn net.Conn
)

func sendPin1111() error {
	_, err := conn.Write([]byte("LO_SERVER_CLIENT_PAIR\n"))
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte("toms remote-control\n"))
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte("1111\n\n"))
	if err != nil {
		return err
	}

	return nil
}

// https://github.com/LibreOffice/impress_remote/blob/master/android/sdremote/mobile/src/main/java/org/libreoffice/impressremote/communication/Protocol.java
// in impress -> shift-esc, then search for remote
// https://wiki.documentfoundation.org/Development/Impress_Remote_Protocol
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: impress-cli-remote <command>")
		fmt.Println("Available commands: pair, transition_next, transition_previous, goto_slide:<number>, presentation_start, presentation_stop")
		os.Exit(1)
	}

	var err error
	conn, err = net.Dial("tcp", defaultHost+":"+defaultPort)
	if err != nil {
		fmt.Printf("error connecting to server: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	err = sendPin1111()
	if err != nil {
		fmt.Printf("error connecting to server: %v", err)
		os.Exit(1)
	}

	command := os.Args[1]
	if len(os.Args) == 2 {
		if command != "pair" { //only send the pin when pairing, then exit
			command = command + "\n\n"
			_, err = conn.Write([]byte(command))
			if err != nil {
				fmt.Printf("error connecting to server: %v", err)
				os.Exit(1)
			}
		}
	} else if command == "goto_slide" && len(os.Args) == 3 {
		command = command + ":" + os.Args[2] + "\n\n"
		_, err = conn.Write([]byte(command))
		if err != nil {
			fmt.Printf("error connecting to server: %v", err)
			os.Exit(1)
		}
	}
	//If this is not here, libreoffice crashes hard. No idea why
	time.Sleep(time.Millisecond * 500)
}
