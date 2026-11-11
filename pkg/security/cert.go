package security

// HTTPS certificate

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"math/big"
	"net"
	"time"
)

// novoCertificado devolve o certificado pronto e o hash SHA-256 dele (em hex).
func novoCertificado(ips []net.IP) (tls.Certificate, string, error) {
	chave, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return tls.Certificate{}, "", err
	}

	modelo := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject:      pkix.Name{CommonName: "train"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(13 * 24 * time.Hour), // menos de 14 dias
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  ips,
		DNSNames:     []string{"localhost"},
	}

	der, err := x509.CreateCertificate(rand.Reader, modelo, modelo, &chave.PublicKey, chave)
	if err != nil {
		return tls.Certificate{}, "", err
	}

	hash := sha256.Sum256(der)
	cert := tls.Certificate{Certificate: [][]byte{der}, PrivateKey: chave}
	return cert, hex.EncodeToString(hash[:]), nil
}

// ipsLocais lista os IPv4 das placas de redes.
// para inserir no certificado
func ipsLocais() []net.IP {
	ips := []net.IP{net.ParseIP("127.0.0.1")}
	addrs, _ := net.InterfaceAddrs()

	for _, a := range addrs {
		if n, ok := a.(*net.IPNet); ok && n.IP.To4() != nil && !n.IP.IsLoopback() {
			ips = append(ips, n.IP)
		}
	}

	return ips
}
