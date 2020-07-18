package _468_validate_ip_address

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	parts := strings.Split(IP, ".")
	if len(parts) == 4 {
		if validIPv4(parts) {
			return "IPv4"
		}
	}
	parts = strings.Split(IP, ":")
	if len(parts) == 8 {
		if validIPv6(parts) {
			return "IPv6"
		}
	}
	return "Neither"
}

func validIPv4(parts []string) bool {
	for _, v := range parts {
		if len(v) == 0 || len(v) > 3 {
			return false
		}
		if len(v) > 1 && v[0] == '0' {
			return false
		}
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil || i > 255 || i < 0 {
			return false
		}
	}
	return true
}

func validIPv6(parts []string) bool {
	for _, v := range parts {
		if len(v) == 0 || len(v) > 4 {
			return false
		}
		v = strings.ToUpper(v)
		i, err := strconv.ParseInt(v, 16, 64)
		if err != nil {
			return false
		}
		if i < 0 || i > 0xFFFF {
			return false
		}
	}
	return true
}
