package utils

import "fmt"

// ToEndpoint formats a given host and port into a socket address `<host>:<port>`.
func ToEndpoint(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
