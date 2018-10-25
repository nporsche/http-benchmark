package main

import (
	"crypto/tls"
)

func cipherMap() map[string]uint16 {
	cipherMap := make(map[string]uint16)

	cipherMap["TLS_RSA_WITH_RC4_128_SHA"] = tls.TLS_RSA_WITH_RC4_128_SHA
	cipherMap["TLS_RSA_WITH_3DES_EDE_CBC_SHA"] = tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA
	cipherMap["TLS_RSA_WITH_AES_128_CBC_SHA"] = tls.TLS_RSA_WITH_AES_128_CBC_SHA
	cipherMap["TLS_RSA_WITH_AES_256_CBC_SHA"] = tls.TLS_RSA_WITH_AES_256_CBC_SHA
	cipherMap["TLS_RSA_WITH_AES_128_CBC_SHA256"] = tls.TLS_RSA_WITH_AES_128_CBC_SHA256
	cipherMap["TLS_RSA_WITH_AES_128_GCM_SHA256"] = tls.TLS_RSA_WITH_AES_128_GCM_SHA256
	cipherMap["TLS_RSA_WITH_AES_256_GCM_SHA384"] = tls.TLS_RSA_WITH_AES_256_GCM_SHA384
	cipherMap["TLS_ECDHE_ECDSA_WITH_RC4_128_SHA"] = tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA
	cipherMap["TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA"] = tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA
	cipherMap["TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA"] = tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA
	cipherMap["TLS_ECDHE_RSA_WITH_RC4_128_SHA"] = tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA
	cipherMap["TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA"] = tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA
	cipherMap["TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA"] = tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA
	cipherMap["TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA"] = tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA
	cipherMap["TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"] = tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256
	cipherMap["TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256"] = tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256
	cipherMap["TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"] = tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
	cipherMap["TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"] = tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
	cipherMap["TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"] = tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
	cipherMap["TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"] = tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
	cipherMap["TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305"] = tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305
	cipherMap["TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305"] = tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305

	return cipherMap
}