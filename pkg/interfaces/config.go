package interfaces

import "time"

var (
	Config ConfigS
)

type ConfigS struct {
	// SeedDB Should we seed the DB with some default data
	SeedDB bool
	ExpireFilters time.Duration
	DebugSQL bool
}

func init() {
	Config = ConfigS{
		SeedDB: true,
		ExpireFilters: 1 * time.Minute,
		DebugSQL: false,
	}
}

