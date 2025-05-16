package setup

import "context"

type SystemSetup struct {
	ctx context.Context
}

// NewSetup initializes a new SystemSetup instance with the provided context and returns it.
func NewSetup(r context.Context) *SystemSetup {
	return &SystemSetup{
		ctx: r,
	}
}
