package mperror

import (
	"errors"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"
)

var (
	ErrInternalError = errors.New("Internal Error")
	ErrInvalidType   = errors.New("Invalid Type")

	ErrAppExists   = errors.New("App Already Exists")
	ErrAppNotFound = errors.New("App Not Found")

	ErrFilterNotFound = errors.New("Filter Not Found")
	ErrFilterConfigNotFound = errors.New("Filter Config Not Found")
	ErrFilterConfigInvalid = errors.New("Filter Config Invalid")

	ErrGroupExists   = errors.New("Group Already Exists")
	ErrGroupNotFound = errors.New("Group Does Not Exist")

	ErrMsgNotInitialized   = errors.New("Message Not Initialized")
	ErrMsgNoAppOwner       = errors.New("Message Has No App Owner")
	ErrMsgMetadataNotFound = errors.New("Message Metadata Not Found")

	ErrMsgFieldNotFound = errors.New("Message Field Not Found")

	ErrTransportInstanceNotFound = errors.New("Transport Instance Not Found")
	ErrTransportInstanceExists   = errors.New("Transport Instance Already Exists")
	ErrTransportConfigInvalid    = errors.New("Transport Config Invalid")

	ErrTransportProviderNotFound = errors.New("Transport Provider Not Found")

	ErrTransportReciepientNotFound = errors.New("Transport Recipient Not Found")
	ErrTransportReciptientExists   = errors.New("Transport Recipient Already Exists")
	ErrTransportRecipientGroupSet  = errors.New("Transport Recipient Group Already Set")
	ErrTransportRecipientUserSet  = errors.New("Transport Recipient User Already Set")
	ErrTransportRecipientGroupOrUserNotSet  = errors.New("Transport Recipient Group or User Not Set")

	ErrUserExists   = errors.New("User Already Exists")
	ErrUserNotFound = errors.New("User Not Found")
)

func FilterErrors(err error) error {
	if ent.IsNotFound(err) {
		return err
	}
	return ErrInternalError
}
