package user

import (
	"context"
	"sync"
	"fmt"

	//	"golang.org/x/crypto/bcrypt"

	"github.com/Fishwaldo/mouthpiece/pkg/dbdriver"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"


	"github.com/mitchellh/mapstructure"
	"github.com/go-logr/logr"
)

type User struct {
	dbUser *ent.DbUser
	lock   sync.RWMutex
	log    logr.Logger
}

func newUser(ctx context.Context, logger logr.Logger, email string, name string) (user *User, err error) {
	user = &User{
		log: logger.WithName("User").WithValues("Email", email),
	}
	user.dbUser, err = dbdriver.DbClient.DbUser.Create().
		SetEmail(email).
		SetName(name).
		Save(ctx)
	if err != nil {
		user.log.Error(err, "Error creating User")
		return nil, mperror.FilterErrors(err)
	}
	if err := user.init(ctx); err != nil {
		user.log.Error(err, "Error initializing User")
		return nil, mperror.FilterErrors(err)
	}
	return user, nil
}

func (u *User) init(ctx context.Context) error {
	return nil
}

func (u *User) Load(ctx context.Context, logger logr.Logger, dbuser any) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	var ok bool
	u.dbUser, ok = dbuser.(*ent.DbUser)
	if !ok {
		logger.Error(mperror.ErrInvalidType, "Invalid Type", "Type", fmt.Sprintf("%T", dbuser))
		return mperror.ErrInvalidType
	}
	u.log = logger.WithName("User").WithValues("Email", u.dbUser.Email)
	return mperror.FilterErrors(u.init(ctx))
}
func (u *User) Save(ctx context.Context) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()
	u.dbUser, err = u.dbUser.Update().Save(ctx)
	if err != nil {
		u.log.Error(err, "Error Saving User")
		return mperror.FilterErrors(err)
	}
	return nil
}

func (u *User) GetEmail() string {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return u.dbUser.Email
}

func (u *User) SetEmail(ctx context.Context, email string) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()
	dbtmp, err := u.dbUser.Update().SetEmail(email).Save(ctx)
	if err != nil {
		u.log.Error(err, "Error Saving User")
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbtmp
	return nil
}

func (u *User) GetID() int {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return u.dbUser.ID
}

func (u *User) GetName() string {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return u.dbUser.Name
}

func (u *User) SetName(ctx context.Context, name string) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()
	dbtmp, err := u.dbUser.Update().SetName(name).Save(ctx)
	if err != nil {
		u.log.Error(err, "Error Saving User")
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbtmp
	return nil
}

func (u *User) GetDescription() string {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return u.dbUser.Description
}

func (u *User) SetDescription(ctx context.Context, description string) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	dbtmp, err := u.dbUser.Update().SetDescription(description).Save(ctx)
	if err != nil {
		u.log.Error(err, "Error Saving User")
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbtmp
	return nil
}

func (u *User) SetFields(ctx context.Context, fields map[string]string) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()
	tx, err := dbdriver.DbClient.Tx(ctx)

	for k, v := range fields {
		if _, err := tx.DbUserMetaData.Create().
			SetUser(u.dbUser).
			SetName(k).
			SetValue(v).
			Save(ctx); err != nil {
			u.log.Error(err, "Error Saving User MetaData", "User", u.dbUser.Email, "Field", k, "Value", v)
			tx.Rollback()
			return mperror.FilterErrors(err)
		}
	}
	tx.Commit()
	u.dbUser = dbdriver.DbClient.DbUser.Query().WithMetadata().Where(dbuser.ID(u.dbUser.ID)).FirstX(ctx)
	return nil
}

func (u *User) GetFields(ctx context.Context) (flds map[string]string, err error) {
	u.lock.RLock()
	defer u.lock.RUnlock()

	var metadata []*ent.DbUserMetaData
	if metadata, err = u.dbUser.Edges.MetadataOrErr(); err != nil {
		metadata, err = u.dbUser.QueryMetadata().All(ctx)
		if err != nil {
			u.log.Error(err, "Error Loading User MetaData")
			return nil, mperror.FilterErrors(err)
		}
	}
	flds = make(map[string]string)
	for _, f := range metadata {
		flds[f.Name] = f.Value
	}
	return flds, nil
}

func (u *User) GetField(ctx context.Context, key string) (value string, err error) {
	u.lock.RLock()
	defer u.lock.RUnlock()

	var metadata []*ent.DbUserMetaData
	if metadata, err = u.dbUser.Edges.MetadataOrErr(); err != nil {
		metadata, err = u.dbUser.QueryMetadata().All(ctx)
		if err != nil {
			u.log.Error(err, "Error Loading User MetaData")
			return "", mperror.FilterErrors(err)
		}
	}
	for _, f := range metadata {
		if f.Name == key {
			return f.Value, nil
		}
	}
	return "", mperror.ErrMsgFieldNotFound
}

func (u *User) SetField(ctx context.Context, key string, value string) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	if _, err := dbdriver.DbClient.DbUserMetaData.Create().
		SetUser(u.dbUser).
		SetName(key).
		SetValue(value).
		Save(ctx); err != nil {
		u.log.Error(err, "Error Saving MetaData", "Field", key, "Value", value)
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbdriver.DbClient.DbUser.Query().WithMetadata().Where(dbuser.ID(u.dbUser.ID)).FirstX(ctx)
	return nil
}

func (u *User) AddFilter(ctx context.Context, filter interfaces.FilterI) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	dbtmp, err := u.dbUser.Update().AddFilterIDs(filter.GetID()).Save(ctx)
	if err != nil {
		u.log.Error(err, "Error adding filter to User", "Filter", filter.GetName())
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbtmp
	return nil
}

func (u *User) DelFilter(ctx context.Context, filter interfaces.FilterI) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	dbtmp, err := u.dbUser.Update().RemoveFilterIDs(filter.GetID()).Save(ctx)
	if err != nil {
		u.log.Error(err, "Error removing filter from User", "Filter", filter.GetName())
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbtmp
	return nil

}
func (u *User) GetFilters(ctx context.Context) (flts []interfaces.FilterI, err error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	// make sure our Filters are loaded
	var dbflts []*ent.DbFilter
	if dbflts, err = u.dbUser.Edges.FiltersOrErr(); err != nil {
		if dbflts, err = u.dbUser.QueryFilters().All(ctx); err != nil {
			u.log.Error(err, "Error loading Filters for User")
			return nil, mperror.FilterErrors(err)
		}
	}

	for _, f := range dbflts {
		flt, err := interfaces.GetFilterService(ctx).Load(ctx, f)
		if err != nil {
			log.Log.Error(err, "Error loading Filter", "Filter", f.Name)
			continue
		}
		flts = append(flts, flt)
	}
	return flts, nil
}

func (u *User) GetTransportRecipients(ctx context.Context) []interfaces.TransportRecipient {
	u.lock.RLock()
	defer u.lock.RUnlock()

	if tpr, err := u.dbUser.QueryTransportRecipients().All(ctx); err != nil {
		u.log.Error(err, "Error getting TransportRecipients for User")
		return nil
	} else {
		tprsvc := interfaces.GetTransportService(ctx)
		var ret []interfaces.TransportRecipient
		for _, t := range tpr {
			if tpr, err := tprsvc.Load(ctx, t); err != nil {
				u.log.Error(err, "Error Loading TransportRecipient for User", "tpr", t.ID)
			} else {
				ret = append(ret, tpr)
			}
		}
		return ret
	}
}

func (u *User) AddTransportRecipient(ctx context.Context, tid interfaces.TransportRecipient) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	u.log.Info("Adding Transport to User", "tid", tid)
	dbtmp, err := u.dbUser.Update().AddTransportRecipientIDs(tid.GetID()).Save(ctx)
	if err != nil {
		u.log.Error(err, "Error adding TransportRecipient to User", "tid", tid)
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbtmp
	return nil
}

func (u *User) DelTransportRecipient(ctx context.Context, tid interfaces.TransportRecipient) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	u.log.Info("Deleting Transport from User", "tid", tid)
	dbtmp, err := u.dbUser.Update().RemoveTransportRecipientIDs(tid.GetID()).Save(ctx)
	if err != nil {
		u.log.Error(err, "Error removing TransportRecipient from User", "tid", tid)
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbtmp
	return nil
}

func (u *User) ProcessMessage(ctx context.Context, msg interfaces.MessageI) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()
	/* add User Fields to Message */
	msg.SetMetadata(ctx, interfaces.MD_UserName, u.dbUser.Name)
	msg.SetMetadata(ctx, interfaces.MD_UserEmail, u.dbUser.Email)
	u.log.V(1).Info("User Processing Message", "MessageID", msg.GetID())

	// make sure our Filters are loaded
	var flts []*ent.DbFilter
	if flts, err = u.dbUser.Edges.FiltersOrErr(); err != nil {
		if flts, err = u.dbUser.QueryFilters().All(ctx); err != nil {
			u.log.Error(err, "Error loading Filters for User")
			return mperror.ErrInternalError
		}
	}

	u.log.V(1).Info("User Filters", "Filters", flts)
	for _, userfilter := range flts {
		u.log.V(1).Info("User Processing Message with Filter", "Filter", userfilter.Name)
		flt, err := interfaces.GetFilterService(ctx).Load(ctx, userfilter)
		if err != nil {
			u.log.Error(err, "Error loading Filter", "Filter", userfilter.Name)
			continue
		}
		if ok, err := flt.ProcessMessage(ctx, msg); err != nil {
			u.log.Error(err, "Error processing message with user filter", "Filter", flt.GetName())
			continue
		} else if ok == interfaces.FilterMatch {
			u.log.Info("User Filter Matched Message", "Filter", flt.GetName(), "Message", msg.GetID())
			break
		} else if ok == interfaces.FilterNoMatch {
			u.log.Info("User Filter did not match Message", "Filter", flt.GetName(), "Message", msg.GetID())
			return nil
		} else if ok == interfaces.FilterPass {
			u.log.Info("User Filter passed Message", "Filter", flt.GetName(), "Message", msg.GetID())
			continue
		}
	}

	var tpr []*ent.DbTransportRecipients
	if tpr, err = u.dbUser.Edges.TransportRecipientsOrErr(); err != nil {
		if tpr, err = u.dbUser.QueryTransportRecipients().All(ctx); err != nil {
			u.log.Error(err, "Error loading TransportRecipients for User")
			return mperror.ErrInternalError
		}
	}
	for _, tr := range tpr {
		u.log.V(1).Info("User Processing Message with TransportRecipient", "TransportRecipient", tr.Name)
		trobj, err := interfaces.GetTransportService(ctx).Load(ctx, tr)
		if err != nil {
			u.log.Error(err, "Error loading TransportRecipient", "TransportRecipient", tr.Name)
			continue
		}
		u.lock.Unlock()
		if err := trobj.ProcessMessage(ctx, msg); err != nil {
			u.log.Error(err, "Error processing message with transport recipient", "TransportRecipient", trobj.GetName())
			u.lock.Lock()
			continue
		}
		u.lock.Lock()
	}
	return nil
}

func (u *User) GetAppData(ctx context.Context, name string, data any) (err error) {
	u.lock.RLock()
	defer u.lock.RUnlock()
	var ok bool
	newdata, ok := u.dbUser.AppData.Data[name]
	if !ok {
		return mperror.ErrAppDataNotFound
	}
	err = mapstructure.Decode(newdata, &data)
	if err != nil {
		u.log.Error(err, "Error decoding AppData", "name", name)
	}

	return nil
}

func (u *User) SetAppData(ctx context.Context, name string, data any) (err error) {
	u.lock.Lock()
	defer u.lock.Unlock()
	appdata := u.dbUser.AppData
	if appdata.Data == nil {
		appdata.Data = make(map[string]any)
	}
	appdata.Data[name] = data
	dbtmp, err := u.dbUser.Update().SetAppData(appdata).Save(ctx)
	if err != nil {
		u.log.Error(err, "Error setting app data on User", "name", name)
		return mperror.FilterErrors(err)
	}
	u.dbUser = dbtmp
	return nil
}



var _ interfaces.UserI = (*User)(nil)

// func (u User) AddRoleToUser(ctx context.Context, role string) bool {
// 	//TODO: Move to Internal
// 	// _, err := auth.AuthService.AuthEnforcer.AddRoleForUser(u.Email, fmt.Sprintf("role:%s", role))
// 	// if err != nil {
// 	// 	log.Log.Info("Failed to add role for user", "email", u.Email, "role", role, "error", err)
// 	// 	return false
// 	// }
// 	return true
// }

// func (u User) CheckPassword(ctx context.Context, password string) bool {
// 	log.Log.Info("Checking Password", "email", u.Email)
// 	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
// 	if err != nil {
// 		log.Log.Info("Password Check Failed", "Error", err)
// 		return false
// 	}
// 	return true
// }

// func (u User) SetPassword(ctx context.Context, password string) error {
// 	log.Log.Info("Setting Password", "Email", u.Email)
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Log.Info("Error Generating SetPassword Hash", "Error", err)
// 		return err
// 	}
// 	if tx := db.Db.WithContext(ctx).Model(&u).Update("password", string(hashedPassword)); tx.Error != nil {
// 		log.Log.Info("Error Setting Password", "Error", tx.Error)
// 		return tx.Error
// 	}
// 	return nil
// }
