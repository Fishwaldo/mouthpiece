package mock_interfaces

import (
)

//go:generate mockgen -destination ./message.go github.com/Fishwaldo/mouthpiece/pkg/interfaces MessageI

//go:generate mockgen -destination ./app.go github.com/Fishwaldo/mouthpiece/pkg/interfaces AppI
