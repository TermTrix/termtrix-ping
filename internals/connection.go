package internals

func SockConn() {

	fd, err := unix.Socket(
		unix.AF_INET,
		unix.SOCK_RAW,
		unix.IPPROTO_ICMP,
	)

}
