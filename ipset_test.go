package ipset

import (
	"encoding/binary"
	"net"
	"testing"
)

func BenchmarkIpsetAdd(b *testing.B) {
	var ip1 uint32
	ip1 = 16777216

	for n := 0; n < b.N; n++ {
		ip := int2ip(ip1 + uint32(n))
		err := Add("blacklist", ip.String())
		if err != nil {
			b.Fatalf("got err: %s", err)
		}

	}
}

func BenchmarkIpsetDel(b *testing.B) {
	var ip1 uint32
	ip1 = 16777216

	for n := 0; n < b.N; n++ {
		ip := int2ip(ip1 + uint32(n))
		err := Del("blacklist", ip.String())
		if err != nil {
			b.Fatalf("got err: %s", err)
		}

	}
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
