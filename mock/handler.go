package mock

import (
	"context"
	"net"

	"github.com/go-gost/core/handler"
	md "github.com/go-gost/core/metadata"
)

type fakeHandler struct {
}

func NewHandler(opts ...handler.Option) handler.Handler {
	return &fakeHandler{}
}

func (c *fakeHandler) Init(md md.Metadata) (err error) {
	return nil
}

func (c *fakeHandler) Handle(_ context.Context, _ net.Conn, opts ...handler.HandleOption) (err error) {
	return nil
}
