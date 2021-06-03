package auth

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
)

// Status represents authorization status.
type Status struct {
	// Authorized is true if client is authorized.
	Authorized bool
	// User is current User object.
	User *tg.User
}

// Status gets authorization status of client.
func (c *Client) Status(ctx context.Context) (*Status, error) {
	u, err := c.self(ctx)
	if IsKeyUnregistered(err) {
		return &Status{}, nil
	}
	if err != nil {
		return nil, err
	}

	return &Status{
		Authorized: true,
		User:       u,
	}, nil
}

// IfNecessary runs given auth flow if current session is not authorized.
func (c *Client) IfNecessary(ctx context.Context, flow Flow) error {
	auth, err := c.Status(ctx)
	if err != nil {
		return xerrors.Errorf("get auth status: %w", err)
	}
	if auth.Authorized {
		return nil
	}
	if err := flow.Run(ctx, c); err != nil {
		return xerrors.Errorf("auth flow: %w", err)
	}
	return nil
}

// Test creates and runs auth flow using Test authenticator
// if current session is not authorized.
func (c *Client) Test(ctx context.Context, dc int) error {
	return c.IfNecessary(ctx, NewFlow(Test(c.rand, dc), SendCodeOptions{}))
}

// TestUser creates and runs auth flow using TestUser authenticator
// if current session is not authorized.
func (c *Client) TestUser(ctx context.Context, phone string, dc int) error {
	return c.IfNecessary(ctx, NewFlow(TestUser(phone, dc), SendCodeOptions{}))
}
