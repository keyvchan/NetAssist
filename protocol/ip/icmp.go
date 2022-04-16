package ip

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/internal"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func ICMPRequest() {
	address := internal.GetArg(3)
	dstAddress, err := net.ResolveIPAddr("ip", address)
	if err != nil {
		log.Fatal(err)
	}
	// linux darwin only
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	body := icmp.Echo{
		ID:   os.Getpid() & 0xffff,
		Seq:  1,
		Data: []byte("HELLO-R-U-THERE"),
	}
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &body,
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.WriteTo(wb, &net.UDPAddr{IP: dstAddress.IP})
	if err != nil {
		log.Fatal(err)
	}
	rb := make([]byte, 1500)
	_, addr, err := conn.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
