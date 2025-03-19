package constant

import "time"

const (
	Expiry10Minutes = 10 * time.Minute
	Expiry5Minutes  = 5 * time.Minute
)

type ModelParams struct {
	Database     string
	Table        string
	DisableCache bool
	LockKey      string
	ModeType     string
}
