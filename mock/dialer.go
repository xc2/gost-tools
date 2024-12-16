package mock

import (
	"context"
	"net"

	"github.com/go-gost/core/dialer"
	md "github.com/go-gost/core/metadata"
)

type fakeDialer struct {
}

func NewDialer(opts ...dialer.Option) dialer.Dialer {
	return &fakeDialer{}
}

func (c *fakeDialer) Init(md md.Metadata) (err error) {
	return nil
}

func (c *fakeDialer) Dial(ctx context.Context, address string, opts ...dialer.DialOption) (net.Conn, error) {
	return nil, nil
}
