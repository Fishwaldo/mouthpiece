package msg

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/copier"

	"github.com/Fishwaldo/mouthpiece/pkg/dbdriver"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessage"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

type Message struct {
	dbEntry  *ent.DbMessage `copier:"-"`
	lock     sync.RWMutex   `copier:"-"`
	log      logr.Logger
	metaData map[string]any `copier:"must"`
	cloned   bool
}

func NewMessage(ctx context.Context, message string, app interfaces.AppI) *Message {
	msg := &Message{
		log:      log.Log.WithName("Message"),
		metaData: make(map[string]any),
		cloned:   false,
	}
	if err := msg.init(); err != nil {
		msg.log.Error(err, "Error Initializing Message")
		return nil
	}
	var err error
	msg.dbEntry, err = dbdriver.DbClient.DbMessage.Create().
		SetMessage(message).
		SetAppID(app.GetID()).
		Save(ctx)
	if err != nil {
		msg.log.Error(err, "Error Saving Message")
		return nil
	}
	msg.log = log.Log.WithName("Message").WithValues("MessageID", msg.dbEntry.ID)
	return msg
}

func LoadFromDB(ctx context.Context, id uuid.UUID) (*Message, error) {
	msg := &Message{
		log:      log.Log.WithName("Message"),
		metaData: make(map[string]any),
	}
	if newmsg, err := dbdriver.DbClient.DbMessage.Query().WithFields().WithApp().Where(dbmessage.ID(id)).First(ctx); err != nil {
		msg.log.Error(err, "Error Loading Message", "Message", msg)
		return nil, mperror.FilterErrors(err)
	} else {
		msg.dbEntry = newmsg
		if err := msg.init(); err != nil {
			msg.log.Error(err, "Error Initializing Message", "Message", msg)
			return nil, mperror.FilterErrors(err)
		}
		msg.log = log.Log.WithName("Message").WithValues("MessageID", msg.dbEntry.ID)
	}
	return msg, nil
}

func (msg *Message) init() error {

	return nil
}

func (msg *Message) Load(ctx context.Context, newmsg any) error {
	msg.lock.Lock()
	defer msg.lock.Unlock()
	var ok bool
	if msg.dbEntry, ok = newmsg.(*ent.DbMessage); !ok {
		msg.log.Error(mperror.ErrInvalidType, "Invalid Type", "Type", fmt.Sprintf("%T", newmsg))
		return mperror.ErrInvalidType
	}
	msg.log = log.Log.WithName("Message").WithValues("MessageID", msg.dbEntry.ID)
	return msg.init()
}

// XXX TODO: Change to Mouthpeice App
func (msg *Message) Save(ctx context.Context, app interfaces.AppI) error {
	msg.lock.Lock()
	defer msg.lock.Unlock()
	if _, err := dbdriver.DbClient.DbMessage.Create().
		SetDbMessageFromStruct(msg.dbEntry).
		SetAppID(app.GetID()).
		Save(ctx); err != nil {
		msg.log.Error(err, "Error Saving Message")
		return mperror.FilterErrors(err)
	}
	msg.dbEntry = dbdriver.DbClient.DbMessage.Query().WithFields().WithApp().Where(dbmessage.ID(msg.dbEntry.ID)).FirstX(ctx)
	msg.log = log.Log.WithName("Message").WithValues("MessageID", msg.dbEntry.ID)
	return nil
}

func (msg *Message) GetID() uuid.UUID {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	return msg.dbEntry.ID
}

func (msg *Message) GetMessage() string {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	return msg.dbEntry.Message
}

func (msg *Message) SetMessage(ctx context.Context, newmessage string) (err error) {
	msg.lock.Lock()
	defer msg.lock.Unlock()
	if msg.cloned {
		msg.log.Error(mperror.ErrMsgLocked, "Message is cloned")
		return mperror.ErrMsgLocked
	}

	dbtmp, err := msg.dbEntry.Update().SetMessage(newmessage).Save(ctx)
	if err != nil {
		msg.log.Error(err, "Error Updating Message")
		return mperror.FilterErrors(err)
	}
	msg.dbEntry = dbtmp
	return nil
}

func (msg *Message) GetShortMsg() *string {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	return msg.dbEntry.ShortMsg
}

func (msg *Message) SetShortMsg(ctx context.Context, shortmsg string) (err error) {
	msg.lock.Lock()
	defer msg.lock.Unlock()

	if msg.cloned {
		msg.log.Error(mperror.ErrMsgLocked, "Message is cloned")
		return mperror.ErrMsgLocked
	}

	dbtmp, err := msg.dbEntry.Update().SetShortMsg(shortmsg).Save(ctx)
	if err != nil {
		msg.log.Error(err, "Error Updating Message")
		return mperror.FilterErrors(err)
	}
	msg.dbEntry = dbtmp
	return nil
}

func (msg *Message) GetTopic() *string {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	return msg.dbEntry.Topic
}

func (msg *Message) SetTopic(ctx context.Context, topic string) (err error) {
	msg.lock.Lock()
	defer msg.lock.Unlock()

	if msg.cloned {
		msg.log.Error(mperror.ErrMsgLocked, "Message is cloned")
		return mperror.ErrMsgLocked
	}

	dbtmp, err := msg.dbEntry.Update().SetTopic(topic).Save(ctx)
	if err != nil {
		msg.log.Error(err, "Error Updating Message")
		return mperror.FilterErrors(err)
	}
	msg.dbEntry = dbtmp
	return nil
}

func (msg *Message) GetSeverity() int {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	return msg.dbEntry.Severity
}

func (msg *Message) SetSeverity(ctx context.Context, sev int) (err error) {
	msg.lock.Lock()
	defer msg.lock.Unlock()

	if msg.cloned {
		msg.log.Error(mperror.ErrMsgLocked, "Message is cloned")
		return mperror.ErrMsgLocked
	}

	dbtmp, err := msg.dbEntry.Update().SetSeverity(sev).Save(ctx)
	if err != nil {
		msg.log.Error(err, "Error Updating Message")
		return mperror.FilterErrors(err)
	}
	msg.dbEntry = dbtmp
	return nil
}

func (msg *Message) GetTimestamp() time.Time {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	return msg.dbEntry.Timestamp
}

func (msg *Message) SetTimestamp(ctx context.Context, ts time.Time) (err error) {
	msg.lock.Lock()
	defer msg.lock.Unlock()

	if msg.cloned {
		msg.log.Error(mperror.ErrMsgLocked, "Message is cloned")
		return mperror.ErrMsgLocked
	}

	dbtmp, err := msg.dbEntry.Update().SetTimestamp(ts).Save(ctx)
	if err != nil {
		msg.log.Error(err, "Error Updating Message")
		return mperror.FilterErrors(err)
	}
	msg.dbEntry = dbtmp
	return nil
}

func (msg *Message) SetFields(ctx context.Context, fields map[string]string) error {
	msg.lock.Lock()
	defer msg.lock.Unlock()

	if msg.cloned {
		msg.log.Error(mperror.ErrMsgLocked, "Message is cloned")
		return mperror.ErrMsgLocked
	}

	tx, err := dbdriver.DbClient.Tx(ctx)
	if err != nil {
		msg.log.Error(err, "Error Starting Transaction", "Message", msg)
		return mperror.FilterErrors(err)
	}

	for k, v := range fields {
		if err := tx.DbMessageFields.Create().
			SetOwner(msg.dbEntry).
			SetName(k).
			SetValue(v).
			OnConflict().
			UpdateNewValues().
			Exec(ctx); err != nil {
			msg.log.Error(err, "Error Saving Message Field", "Message", msg, "Field", k, "Value", v)
			tx.Rollback()
			return mperror.FilterErrors(err)
		}
	}
	tx.Commit()
	msg.dbEntry = dbdriver.DbClient.DbMessage.Query().WithFields().WithApp().Where(dbmessage.ID(msg.dbEntry.ID)).FirstX(ctx)
	return nil
}

func (msg *Message) GetFields(ctx context.Context) (flds map[string]string, err error) {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	var fields []*ent.DbMessageFields
	if fields, err = msg.dbEntry.Edges.FieldsOrErr(); err != nil {
		fields, err = msg.dbEntry.QueryFields().All(ctx)
		if err != nil {
			msg.log.Error(err, "Error Getting Message Fields", "Message", msg)
			return nil, mperror.ErrInternalError
		}
	}
	flds = make(map[string]string)
	for _, f := range fields {
		flds[f.Name] = f.Value
	}
	return flds, nil
}

func (msg *Message) GetField(ctx context.Context, key string) (value string, err error) {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	var fields []*ent.DbMessageFields
	if fields, err = msg.dbEntry.Edges.FieldsOrErr(); err != nil {
		fields, err = msg.dbEntry.QueryFields().All(ctx)
		if err != nil {
			msg.log.Error(err, "Error Getting Message Fields", "Message", msg)
			return "", mperror.ErrInternalError
		}
	}
	for _, f := range fields {
		if f.Name == key {
			return f.Value, nil
		}
	}
	return "", mperror.ErrMsgFieldNotFound
}

func (msg *Message) SetField(ctx context.Context, key string, value string) (err error) {
	msg.lock.Lock()
	defer msg.lock.Unlock()

	if msg.cloned {
		msg.log.Error(mperror.ErrMsgLocked, "Message is cloned")
		return mperror.ErrMsgLocked
	}

	if err := dbdriver.DbClient.DbMessageFields.Create().
		SetOwner(msg.dbEntry).
		SetName(key).
		SetValue(value).
		OnConflict().
		UpdateNewValues().
		Exec(ctx); err != nil {
		msg.log.Error(err, "Error Saving Message Field", "Message", msg, "Field", key, "Value", value)
		return mperror.FilterErrors(err)
	}
	msg.dbEntry = dbdriver.DbClient.DbMessage.Query().WithFields().WithApp().Where(dbmessage.ID(msg.dbEntry.ID)).FirstX(ctx)
	return nil
}

func (msg *Message) GetApp(ctx context.Context) (interfaces.AppI, error) {
	msg.lock.RLock()
	defer msg.lock.RUnlock()
	dbapp, err := msg.dbEntry.QueryApp().Only(ctx)
	if err != nil {
		msg.log.Error(err, "Error Getting Message App", "Message", msg)
		return nil, mperror.FilterErrors(err)
	}
	return interfaces.GetAppService(ctx).Load(ctx, dbapp)
}

func (msg *Message) ProcessMessage(ctx context.Context) (err error) {
	msg.lock.Lock()
	defer msg.lock.Unlock()

	msg.cloned = true

	if msg.dbEntry == nil {
		return mperror.ErrMsgNotInitialized
	}
	if _, err := msg.dbEntry.Edges.AppOrErr(); err != nil {
		if _, err = msg.dbEntry.QueryApp().OnlyID(ctx); err != nil {
			return mperror.ErrMsgNoAppOwner
		}
	}
	log.Log.V(1).Info("Processing Message", "Message", msg)
	// appsvc := interfaces.GetAppService(ctx)
	// if appsvc == nil {
	// 	log.Log.Error(nil, "No App Service Found On Context")
	// 	return mperror.ErrMsgNotInitialized
	// }
	// if app, err := appsvc.GetApp(ctx, ID); err != nil {
	// 	log.Log.Error(err, "Error Getting App", "AppID", ID)
	// 	return err
	// } else {
	// 	if err := app.ProcessMessage(ctx, msg); err != nil {
	// 		log.Log.Error(err, "Error Processing Message", "Message", msg)
	// 		return err
	// 	}
	// }
	return nil
}

func (msg *Message) String() string {
	//	msg.lock.RLock()
	//	defer msg.lock.RUnlock()
	return msg.dbEntry.String()
}

func (msg *Message) SetMetadata(ctx context.Context, key string, value any) error {
	msg.lock.Lock()
	defer msg.lock.Unlock()
	msg.metaData[key] = value
	return nil
}

func (msg *Message) GetMetadata(ctx context.Context, key string) (any, error) {
	msg.lock.RLock()
	defer msg.lock.RUnlock()
	if val, ok := msg.metaData[key]; ok {
		return val, nil
	}
	return nil, mperror.ErrMsgMetadataNotFound
}

func (msg *Message) GetMetadataFields(ctx context.Context) (map[string]any, error) {
	msg.lock.RLock()
	defer msg.lock.RUnlock()
	metadata := make(map[string]any)
	for k, v := range msg.metaData {
		metadata[k] = v
	}

	return metadata, nil
}

func (msg *Message) Clone() interfaces.MessageI {
	msg.lock.RLock()
	defer msg.lock.RUnlock()

	msg.cloned = true

	newmsg := Message{}
	if err := copier.CopyWithOption(&newmsg, msg, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		msg.log.Error(err, "Error Cloning Message", "Message", msg)
		return nil
	}
	newmsg.dbEntry = msg.dbEntry
	newmsg.metaData = make(map[string]any)
	for k, v := range msg.metaData {
		newmsg.metaData[k] = v
	}

	/* the Copied message will be locked, so unlock it */
	newmsg.lock.RUnlock()
	return &newmsg
}


var _ interfaces.MessageI = (*Message)(nil)