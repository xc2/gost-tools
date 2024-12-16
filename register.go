package main

import (
	"github.com/go-gost/x/registry"
	"github.com/xc2/gost-tools/mock"
)

// import (
// 	// Register connectors
// 	"fmt"

// 	_ "github.com/go-gost/x/connector/direct"
// 	_ "github.com/go-gost/x/connector/forward"
// 	_ "github.com/go-gost/x/connector/http"
// 	_ "github.com/go-gost/x/connector/http2"
// 	_ "github.com/go-gost/x/connector/relay"
// 	_ "github.com/go-gost/x/connector/serial"
// 	_ "github.com/go-gost/x/connector/sni"
// 	_ "github.com/go-gost/x/connector/socks/v4"
// 	_ "github.com/go-gost/x/connector/socks/v5"
// 	_ "github.com/go-gost/x/connector/ss"
// 	_ "github.com/go-gost/x/connector/ss/udp"
// 	_ "github.com/go-gost/x/connector/sshd"
// 	_ "github.com/go-gost/x/connector/tcp"
// 	_ "github.com/go-gost/x/connector/tunnel"
// 	_ "github.com/go-gost/x/connector/unix"
// 	"github.com/go-gost/x/registry"

// 	// Register dialers
// 	_ "github.com/go-gost/x/dialer/direct"
// 	_ "github.com/go-gost/x/dialer/dtls"
// 	_ "github.com/go-gost/x/dialer/ftcp"
// 	_ "github.com/go-gost/x/dialer/grpc"
// 	_ "github.com/go-gost/x/dialer/http2"
// 	_ "github.com/go-gost/x/dialer/http2/h2"
// 	_ "github.com/go-gost/x/dialer/http3"
// 	_ "github.com/go-gost/x/dialer/http3/wt"
// 	_ "github.com/go-gost/x/dialer/icmp"
// 	_ "github.com/go-gost/x/dialer/kcp"
// 	_ "github.com/go-gost/x/dialer/mtcp"
// 	_ "github.com/go-gost/x/dialer/mtls"
// 	_ "github.com/go-gost/x/dialer/mws"
// 	_ "github.com/go-gost/x/dialer/obfs/http"
// 	_ "github.com/go-gost/x/dialer/obfs/tls"
// 	_ "github.com/go-gost/x/dialer/pht"
// 	_ "github.com/go-gost/x/dialer/quic"
// 	_ "github.com/go-gost/x/dialer/serial"
// 	_ "github.com/go-gost/x/dialer/ssh"
// 	_ "github.com/go-gost/x/dialer/sshd"
// 	_ "github.com/go-gost/x/dialer/tcp"
// 	_ "github.com/go-gost/x/dialer/tls"
// 	_ "github.com/go-gost/x/dialer/udp"
// 	_ "github.com/go-gost/x/dialer/unix"
// 	_ "github.com/go-gost/x/dialer/ws"

// 	// Register handlers
// 	_ "github.com/go-gost/x/handler/auto"
// 	_ "github.com/go-gost/x/handler/dns"
// 	_ "github.com/go-gost/x/handler/file"
// 	_ "github.com/go-gost/x/handler/forward/local"
// 	_ "github.com/go-gost/x/handler/forward/remote"
// 	_ "github.com/go-gost/x/handler/http"
// 	_ "github.com/go-gost/x/handler/http2"
// 	_ "github.com/go-gost/x/handler/http3"
// 	_ "github.com/go-gost/x/handler/metrics"
// 	_ "github.com/go-gost/x/handler/redirect/tcp"
// 	_ "github.com/go-gost/x/handler/redirect/udp"
// 	_ "github.com/go-gost/x/handler/relay"
// 	_ "github.com/go-gost/x/handler/serial"
// 	_ "github.com/go-gost/x/handler/sni"
// 	_ "github.com/go-gost/x/handler/socks/v4"
// 	_ "github.com/go-gost/x/handler/socks/v5"
// 	_ "github.com/go-gost/x/handler/ss"
// 	_ "github.com/go-gost/x/handler/ss/udp"
// 	_ "github.com/go-gost/x/handler/sshd"
// 	_ "github.com/go-gost/x/handler/tap"
// 	_ "github.com/go-gost/x/handler/tun"
// 	_ "github.com/go-gost/x/handler/tunnel"
// 	_ "github.com/go-gost/x/handler/unix"

// 	// Register listeners
// 	_ "github.com/go-gost/x/listener/dns"
// 	_ "github.com/go-gost/x/listener/dtls"
// 	_ "github.com/go-gost/x/listener/ftcp"
// 	_ "github.com/go-gost/x/listener/grpc"
// 	_ "github.com/go-gost/x/listener/http2"
// 	_ "github.com/go-gost/x/listener/http2/h2"
// 	_ "github.com/go-gost/x/listener/http3"
// 	_ "github.com/go-gost/x/listener/http3/h3"
// 	_ "github.com/go-gost/x/listener/http3/wt"
// 	_ "github.com/go-gost/x/listener/icmp"
// 	_ "github.com/go-gost/x/listener/kcp"
// 	_ "github.com/go-gost/x/listener/mtcp"
// 	_ "github.com/go-gost/x/listener/mtls"
// 	_ "github.com/go-gost/x/listener/mws"
// 	_ "github.com/go-gost/x/listener/obfs/http"
// 	_ "github.com/go-gost/x/listener/obfs/tls"
// 	_ "github.com/go-gost/x/listener/pht"
// 	_ "github.com/go-gost/x/listener/quic"
// 	_ "github.com/go-gost/x/listener/redirect/tcp"
// 	_ "github.com/go-gost/x/listener/redirect/udp"
// 	_ "github.com/go-gost/x/listener/rtcp"
// 	_ "github.com/go-gost/x/listener/rudp"
// 	_ "github.com/go-gost/x/listener/serial"
// 	_ "github.com/go-gost/x/listener/ssh"
// 	_ "github.com/go-gost/x/listener/sshd"
// 	_ "github.com/go-gost/x/listener/tap"
// 	_ "github.com/go-gost/x/listener/tcp"
// 	_ "github.com/go-gost/x/listener/tls"
// 	_ "github.com/go-gost/x/listener/tun"
// 	_ "github.com/go-gost/x/listener/udp"
// 	_ "github.com/go-gost/x/listener/unix"
// 	_ "github.com/go-gost/x/listener/ws"
// )

var (
	handlers = []string{
		"serial", "http2", "red", "sni", "tcp", "ssu", "socks5", "dns", "forward", "relay", "redirect", "ss", "tunnel", "socks", "tun", "metrics", "http3", "http", "socks4a", "redu", "udp", "rudp", "file", "sshd", "unix", "redir", "socks4", "auto", "tap", "rtcp",
	}
	dialers = []string{
		"sshd", "ohttp", "ssh", "icmp", "http2", "serial", "http3", "grpc", "mtcp", "tcp", "dtls", "quic", "wt", "h3", "tls", "mtls", "direct", "mwss", "phts", "unix", "mws", "h2c", "otls", "kcp", "ws", "pht", "wss", "virtual", "ohttps", "icmp6", "udp", "h2", "ftcp",
	}
	connectors = []string{
		"ssu", "tcp", "sshd", "http", "sni", "ss", "socks", "tunnel", "socks4", "relay", "direct", "serial", "virtual", "http2", "socks4a", "socks5", "unix", "forward",
	}
	listeners = []string{
		"kcp", "mtcp", "ws", "udp", "ssh", "h2", "redir", "tun", "otls", "h3", "icmp6", "grpc", "red", "dtls", "rtcp", "http3", "redirect", "serial", "quic", "pht", "mws", "phts", "mwss", "mtls", "sshd", "tls", "rudp", "tcp", "http2", "unix", "tap", "redu", "dns", "h2c", "ftcp", "wss", "icmp", "ohttp", "wt",
	}
)

func init() {
	for _, h := range connectors {
		registry.ConnectorRegistry().Register(h, mock.NewConnector)
	}
	for _, h := range dialers {
		registry.DialerRegistry().Register(h, mock.NewDialer)
	}
	for _, h := range listeners {
		registry.ListenerRegistry().Register(h, mock.NewListener)
	}
	for _, h := range handlers {
		registry.HandlerRegistry().Register(h, mock.NewHandler)
	}
}

// func init() {
// 	fmt.Println("var (")
// 	fmt.Println("handlers = []string{")
// 	for r := range registry.HandlerRegistry().GetAll() {
// 		fmt.Printf("\"%s\",", r)
// 	}
// 	fmt.Println("\n}")
// 	fmt.Println("dialers = []string{")
// 	for r := range registry.DialerRegistry().GetAll() {
// 		fmt.Printf("\"%s\",", r)
// 	}
// 	fmt.Println("\n}")
// 	fmt.Println("connectors = []string{")
// 	for r := range registry.ConnectorRegistry().GetAll() {
// 		fmt.Printf("\"%s\",", r)
// 	}
// 	fmt.Println("\n}")
// 	fmt.Println("listeners = []string{")
// 	for r := range registry.ListenerRegistry().GetAll() {
// 		fmt.Printf("\"%s\",", r)
// 	}
// 	fmt.Println("\n}")
// 	fmt.Println(")")
// }
