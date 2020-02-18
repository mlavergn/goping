package ping

import (
	"log"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// Ping export
func Ping(hostname string) bool {
	// resovle the host/ip to an ipaddr
	ipaddr, err := net.ResolveIPAddr("ip4", hostname)
	if err != nil {
		log.Println("failed to resolve", err)
		return false
	}

	// establishing a listener for ICMP typically requires root access,
	// but UDP offers an unprivileged option on macos and linux
	var netAddr net.Addr = ipaddr
	var proto = "ip4:icmp"
	useUDP := true
	if useUDP {
		proto = "udp4"
		netAddr = &net.UDPAddr{
			IP:   ipaddr.IP,
			Zone: ipaddr.Zone,
		}
	}

	// establish a listener
	conn, err := icmp.ListenPacket(proto, "0.0.0.0")
	if err != nil {
		log.Println("failed to listen", err)
		return false
	}
	defer conn.Close()

	// generate an ICMP message
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   1,
			Seq:  1,
			Data: []byte(""),
		},
	}

	// get the message byte representation
	bytes, err := msg.Marshal(nil)
	if err != nil {
		log.Println("failed to marshal", err)
		return false
	}

	// send the message
	wlen, err := conn.WriteTo(bytes, netAddr)
	if err != nil {
		log.Println("failed to write", err)
		return false
	} else if wlen != len(bytes) {
		log.Println("failed to write all", err)
		return false
	}

	// Wait for a reply
	buffer := make([]byte, 1500)
	err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		log.Println("read deadline expired", err)
		return false
	}

	rlen, _, err := conn.ReadFrom(buffer)
	if err != nil {
		// normal failure case for a down / non-existing host
		return false
	}

	// parse the reply
	ICMPProtocol := 1
	message, err := icmp.ParseMessage(ICMPProtocol, buffer[:rlen])
	if err != nil {
		log.Println("failed to parse", err)
		return false
	}

	// detemine if we received an echo
	switch message.Type {
	case ipv4.ICMPTypeEchoReply:
		// normal success case for an up host
		return true
	default:
		log.Println("non-echo reply", err)
		return false
	}
}
