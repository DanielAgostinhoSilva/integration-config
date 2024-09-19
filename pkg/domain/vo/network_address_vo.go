package vo

import (
	"context"
	"github.com/DanielAgostinhoSilva/integration-config/pkg/domain/errors"
	"net"
	"regexp"
)

const ipPattern = `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`

var (
	ErrInvalidNetworkAddress  = errors.NewBusinessError("invalid network address")
	ErrNoIpAssociatedWithHost = errors.NewBusinessError("no IPs associated with the hostname")
)

type NetworkAddress struct {
	value string
}

func (n NetworkAddress) Value() string {
	return n.value
}

// NewNetworkAddress creates a new NetworkAddress instance and validates the host
func NewNetworkAddress(ctx context.Context, host string) (*NetworkAddress, error) {
	if isValidIP(host) {
		return &NetworkAddress{value: host}, nil
	}

	ips, err := net.DefaultResolver.LookupHost(ctx, host)
	if err != nil {
		return nil, ErrInvalidNetworkAddress
	}
	if len(ips) == 0 {
		return nil, ErrNoIpAssociatedWithHost
	}

	return &NetworkAddress{value: host}, nil
}

// isValidIP checks if the given string is a valid IP address
func isValidIP(ip string) bool {
	re := regexp.MustCompile(ipPattern)
	return re.MatchString(ip)
}
