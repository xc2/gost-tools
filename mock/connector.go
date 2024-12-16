package mock

import (
	"context"
	"net"

	"github.com/go-gost/core/connector"
	md "github.com/go-gost/core/metadata"
)

type fakeConnector struct {
}

func NewConnector(opts ...connector.Option) connector.Connector {
	return &fakeConnector{}
}

func (c *fakeConnector) Init(md md.Metadata) (err error) {
	return nil
}

func (c *fakeConnector) Connect(ctx context.Context, _ net.Conn, network, address string, opts ...connector.ConnectOption) (net.Conn, error) {
	return nil, nil
}
