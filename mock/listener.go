package mock

import (
	"net"

	"github.com/go-gost/core/listener"
	md "github.com/go-gost/core/metadata"
)

type fakeListener struct {
}

func NewListener(opts ...listener.Option) listener.Listener {
	return &fakeListener{}
}

func (c *fakeListener) Init(_ md.Metadata) (err error) {
	return nil
}
func (c *fakeListener) Accept() (net.Conn, error) {
	return nil, nil
}

func (c *fakeListener) Addr() net.Addr {
	return nil
}

func (c *fakeListener) Close() error {
	return nil
}
