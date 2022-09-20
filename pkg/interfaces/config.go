package interfaces

import (

)

var (
	Config ConfigS
)

type ConfigS struct {
	// SeedDB Should we seed the DB with some default data
	SeedDB bool
}

func init() {
	Config = ConfigS{
		SeedDB: true,
	}
}

