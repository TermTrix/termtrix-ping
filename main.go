package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

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
	targetIP := net.ParseIP(*target).To4()

	if targetIP == nil {
		fmt.Println("Target is not a valid one....")
		return
	}

	if len(*message) < 1 {
		*message = "TERMTRIX"
	}

	osType := runtime.GOOS

	if osType != "linux" {
		fmt.Println("Please continue with linux distro....")
		return
	}

	input := internals.UserInputs{
		TargetIP: targetIP,
		Message:  *message,
	}

	packet := internals.BuildPacket(input)

	buff := internals.SockConn(targetIP, packet)

	VERSION, IHL, TOTAL_BYTES := internals.Find_version_ihl(buff)
	icmp := internals.ICMP{
		Version:       VERSION,
		Ihl:           IHL,
		TotalBytes:    TOTAL_BYTES,
		SourceIP:      net.IP(buff[12:16]),
		DestinationIP: net.IP(buff[16:20]),
	}

	icmp_echo_reply_packets := buff[icmp.TotalBytes:]

	internals.ParseEchoReply(icmp_echo_reply_packets, &icmp)

	internals.PrintICMP(icmp)
}
