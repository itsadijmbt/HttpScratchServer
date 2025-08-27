package main

import "net"

const address string = "http://127.0.0.1:42069"

func main() {

	addr, err := net.ResolveUDPAddr(address)

	updConn, err := net.DialUDP(address,addr,)

}
