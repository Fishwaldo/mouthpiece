package mperror

import (
	"errors"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"
)

var (
	ErrUnsupportedDBType = errors.New("Unsupported Database Type")
	ErrInternalError = errors.New("Internal Error")
	ErrInvalidType   = errors.New("Invalid Type")

	ErrAppExists   = errors.New("App Already Exists")
	ErrAppNotFound = errors.New("App Not Found")

	ErrFilterNotFound       = errors.New("Filter Not Found")
	ErrFilterConfigNotFound = errors.New("Filter Config Not Found")
	ErrFilterConfigInvalid  = errors.New("Filter Config Invalid")
	ErrFilterImplNotFound   = errors.New("Filter Implementation Not Found")

	ErrGroupExists   = errors.New("Group Already Exists")
	ErrGroupNotFound = errors.New("Group Does Not Exist")

	ErrMsgNotInitialized   = errors.New("Message Not Initialized")
	ErrMsgNoAppOwner       = errors.New("Message Has No App Owner")
	ErrMsgMetadataNotFound = errors.New("Message Metadata Not Found")
	ErrMsgLocked           = errors.New("Message Locked")

	ErrMsgFieldNotFound = errors.New("Message Field Not Found")

	ErrTransportInstanceNotFound = errors.New("Transport Instance Not Found")
	ErrTransportInstanceExists   = errors.New("Transport Instance Already Exists")
	ErrTransportConfigInvalid    = errors.New("Transport Config Invalid")

	ErrTransportProviderNotFound = errors.New("Transport Provider Not Found")

	ErrTransportReciepientNotFound         = errors.New("Transport Recipient Not Found")
	ErrTransportReciptientExists           = errors.New("Transport Recipient Already Exists")
	ErrTransportRecipientSet               = errors.New("Transport Recipient Group or User Already Set")
	ErrTransportRecipientGroupOrUserNotSet = errors.New("Transport Recipient Group or User Not Set")

	ErrUserExists   = errors.New("User Already Exists")
	ErrUserNotFound = errors.New("User Not Found")

	ErrValidationError = errors.New("Validation Error")
	ErrAppDataNotFound = errors.New("App Data Not Found")
)

func FilterErrors(err error) error {
	if err == nil {
		return nil
	}
	if ent.IsNotFound(err) {
		return err
	}
	if ent.IsValidationError(err) {
		return ErrValidationError
	}
	return ErrInternalError
}
