package rootCA

import "time"

var (
	rootCAKeyPath     = "ssl/private/rootCA.key"
	rootCACertPath    = "ssl/root/rootCA.crt"
	rootCAKeyPassword = ""
	serialFile        = "ssl/serial"
	revokeListPath    = "ssl/revokelist"
	crlLifetime       = 3600 * time.Second
)
