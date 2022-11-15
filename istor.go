package istor

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
)

// IsExitNode returns true if an IP is a Tor exit node
func IsExitNode(ctx context.Context, ip string) (bool, error) {
	// Reverse IP
	split := strings.Split(ip, ".")
	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}

	// Launch DNS query
	dns := fmt.Sprintf("%s.dnsel.torproject.org", strings.Join(split, "."))
	ips, err := net.DefaultResolver.LookupIPAddr(ctx, dns)
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	// If it returns address 127.0.0.2, then it is a Tor exit node.
	for _, addr := range ips {
		if addr.IP.Equal(net.IPv4(127, 0, 0, 2)) {
			return true, nil
		}
	}
	return false, nil
}
