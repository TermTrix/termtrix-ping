package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/termtrix-ping/internals"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Plase follow this [PATTERN] -> (entrypoint :) -t 192.268.2.20")
		return
	}
	target := flag.String("t", "", "Enter your target :)")
	message := flag.String("m", "", "Your message!!")
	flag.Parse()
	targetIP := net.ParseIP(*target)

	if targetIP == nil {
		fmt.Println("Target is not a valid one....")
		return
	}

	if len(*message) < 1 {
		*message = "TERMTRIX"
	}

	fmt.Println("TARGET IP :", targetIP, *message)

	input := internals.UserInputs{
		TargetIP: targetIP,
		Message:  *message,
	}

	packet := internals.BuildPacket(input)

	fmt.Println("PACKET : ", packet)

	internals.SockConn(targetIP, packet)

}
