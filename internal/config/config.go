package config

// TransportType defines the type of transport.
type TransportType string

const (
	TCP    TransportType = "tcp"
	TCPMUX TransportType = "tcpmux"
	WS     TransportType = "ws"
)

// ServerConfig represents the configuration for the server.
type ServerConfig struct {
	BindAddr       string        `toml:"bind_addr"`
	Transport      TransportType `toml:"transport"`
	Token          string        `toml:"token"`
	Nodelay        bool          `toml:"nodelay"`
	Keepalive      int           `toml:"keepalive_period"`
	ChannelSize    int           `toml:"channel_size"`
	LogLevel       string        `toml:"log_level"`
	ConnectionPool int           `toml:"connection_pool"`
	Ports          []string      `toml:"ports"`
	PPROF          bool          `toml:"pprof"`
	// smux config
	MuxSession       int `toml:"mux_session"`
	MuxVersion       int `toml:"mux_version"`
	MaxFrameSize     int `toml:"mux_framesize"`
	MaxReceiveBuffer int `toml:"mux_recievebuffer"`
	MaxStreamBuffer  int `toml:"mux_streambuffer"`
}

// ClientConfig represents the configuration for the client.
type ClientConfig struct {
	RemoteAddr    string        `toml:"remote_addr"`
	Transport     TransportType `toml:"transport"`
	Token         string        `toml:"token"`
	RetryInterval int           `toml:"retry_interval"`
	Nodelay       bool          `toml:"nodelay"`
	Keepalive     int           `toml:"keepalive_period"`
	LogLevel      string        `toml:"log_level"`
	Forwarder     []string      `toml:"forwarder"`
	PPROF         bool          `toml:"pprof"`
	// smux config
	MuxSession       int `toml:"mux_session"`
	MuxVersion       int `toml:"mux_version"`
	MaxFrameSize     int `toml:"mux_framesize"`
	MaxReceiveBuffer int `toml:"mux_recievebuffer"`
	MaxStreamBuffer  int `toml:"mux_streambuffer"`
}

// Config represents the complete configuration, including both server and client settings.
type Config struct {
	Server ServerConfig `toml:"server"`
	Client ClientConfig `toml:"client"`
}
