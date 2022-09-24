package mock_interfaces

import (
)

//go:generate mockgen -destination ./message.go github.com/Fishwaldo/mouthpiece/pkg/interfaces MessageI

//go:generate mockgen -destination ./transportrecipient.go github.com/Fishwaldo/mouthpiece/pkg/interfaces TransportRecipient

//go:generate mockgen -destination ./app.go github.com/Fishwaldo/mouthpiece/pkg/interfaces AppI

//go:generate mockgen -destination ./user.go github.com/Fishwaldo/mouthpiece/pkg/interfaces UserI

//go:generate mockgen -destination ./group.go github.com/Fishwaldo/mouthpiece/pkg/interfaces GroupI
