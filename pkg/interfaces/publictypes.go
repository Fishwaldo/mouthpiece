package interfaces

//CtxUserValue Context Key to get token.User value from Context
type CtxUserValue struct{}

type AppDetails struct {
	ID          uint   `doc:"App ID" gorm:"primary_key"`
	AppName     string `doc:"Application Name" pattern:"^[a-z0-9]+$" gorm:"unique;uniqueIndex" validate:"required,max=255,alphanum"`
	Status      string `doc:"Status of Application" enum:"Enabled,Disabled" default:"Enabled" validate:"required,oneof=Enabled Disabled"`
	Description string `doc:"Description of Application" validate:"required,max=255"`
	Icon        string `doc:"Icon of Application" validate:"url"`
	URL         string `doc:"URL of Application" validate:"url"`
}

type UserDetails struct {
	ID        uint   `doc:"User ID" gorm:"primary_key"`
	Email     string `doc:"Email" validate:"required,email"`
	FirstName string `doc:"First Name" validate:"required"`
	LastName  string `doc:"Last Name" validate:"required"`
	Password  string `doc:"Password" json:"-" writeOnly:"true" validate:"required"`
}
