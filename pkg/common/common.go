package common

type config struct {
	TCPServerConfig tcpServerConfig
	UDPServerConfig udpServerConfig
}

type tcpServerConfig struct {
	Enabled bool
	Addr    string
}

type udpServerConfig struct {
	Enabled bool
	Addr    string
}

var _config config

func Config() *config {
	return &_config
}
