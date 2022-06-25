package checker

import (
	"crypto/tls"
	"fmt"
	"time"
)

func Check(domain string) {
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}

	err = conn.VerifyHostname(domain)
	if err != nil {
		panic("Hostname doesn't match with certificate: " + err.Error())
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	fmt.Printf("Your cert going to expire on  %v\n", expiry.Format(time.RFC850))
}
