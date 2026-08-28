package internals

import (
	"fmt"
	"net"

	"golang.org/x/sys/unix"
)

func SockConn(target net.IP, packet []byte) []byte {

	fd, err := unix.Socket(
		unix.AF_INET,
		unix.SOCK_RAW,
		unix.IPPROTO_ICMP,
	)

	if err != nil {
		fmt.Println("Failed to connect socket conn")
		return nil
	}

	addr := unix.SockaddrInet4{
		Port: 0,
	}

	copy(addr.Addr[:], target)

	unix.Sendto(fd, packet, 0, &addr)

	buff := make([]byte, 2048)

	unix.Recvfrom(fd, buff, 0)

	return buff
}
