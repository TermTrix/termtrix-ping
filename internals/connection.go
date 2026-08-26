package internals

import (
	"fmt"
	"net"

	"golang.org/x/sys/unix"
)

func SockConn(target net.IP, packet []byte) {

	fd, err := unix.Socket(
		unix.AF_INET,
		unix.SOCK_RAW,
		unix.IPPROTO_ICMP,
	)

	if err != nil {
		fmt.Println("Failed to connect socket conn")
		return
	}

	addr := unix.SockaddrInet4{
		Addr: [4]byte(target),
	}

	unix.Sendto(fd, packet, 0, &addr)

	buff := make([]byte, 1024)

	unix.Recvfrom(fd, buff, 0)

	fmt.Println(buff)

}
