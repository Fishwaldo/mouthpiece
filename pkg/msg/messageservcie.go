package msg

import (
	"context"
//	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/dbdriver"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessage"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
)



type MessageService struct {
	log logr.Logger
}

func NewMsgService(ctx context.Context, log logr.Logger) *MessageService {
	newlog := log.WithName("MsgService")
	newlog.V(1).Info("New Msg Service")
	return &MessageService{
		log: newlog,
	}
}

func (a *MessageService) Start(ctx context.Context) error {
	return nil
}

func (a *MessageService) Get(ctx context.Context, id uuid.UUID) (msg interfaces.MessageI, err error) {
	db_app, err := dbdriver.DbClient.DbMessage.Query().WithFields().Where(dbmessage.ID(id)).Only(ctx)
	if err != nil {
		a.log.V(1).Error(err, "Error getting message", "id", id)
		return nil, mperror.FilterErrors(err)
	}
	return a.Load(ctx, db_app)
}

func (a *MessageService) Load(ctx context.Context, db_msg any) (msg interfaces.MessageI, err error) {
	entMsg, ok := db_msg.(*ent.DbMessage)
	if !ok {
		a.log.Error(nil, "Error loading message")
		return nil, mperror.ErrInvalidType
	}
	newmsg := &Message{}
	if err := newmsg.Load(ctx, entMsg); err != nil {
		a.log.Error(err, "Error loading msg", "UUID", entMsg.ID)
		return nil, mperror.ErrInvalidType
	}
	return newmsg, nil
}

func (a *MessageService) GetMessages(ctx context.Context, options ...func(query any) any) ([]interfaces.MessageI, error) {
	query := dbdriver.DbClient.DbMessage.Query().WithFields().WithApp()
	for _, option := range options {
		query = option(query).(*ent.DbMessageQuery)
	}
	
	db_msgs, err := query.All(ctx)
	if err != nil {
		a.log.V(1).Error(err, "Error getting messages")
		return nil, mperror.FilterErrors(err)
	}
	var msgs []interfaces.MessageI
	for _, db_msg := range db_msgs {
		msg, err := a.Load(ctx, db_msg)
		if err != nil {
			a.log.Error(err, "Error loading message", "UUID", db_msg.ID)
			continue
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

func (a *MessageService) GetMessageCount(ctx context.Context) (int, error) {
	count, err := dbdriver.DbClient.DbMessage.Query().Count(ctx)
	if err != nil {
		a.log.V(1).Error(err, "Error getting message count")
		return 0, mperror.FilterErrors(err)
	}
	return count, nil
}

func WithPaginate(page int, pagesize int) func(db any) any {
	return func(db any) any {
		msgdb, ok := db.(*ent.DbMessageQuery)
		if !ok {
			return db
		}
		offset := (page - 1) * pagesize
		return msgdb.Offset(offset).Limit(pagesize)
	}
}

func WithSort(field string, dir string) func(db any) any {
	return func(db any) any {
		
		switch field {
		case "TimeStamp":
			field = dbmessage.FieldTimestamp
		case "Severity":
			field = dbmessage.FieldSeverity
		case "Message":
			field = dbmessage.FieldMessage
		case "Topic":
			field = dbmessage.FieldTopic
		default:
			return db
		}

		msgdb, ok := db.(*ent.DbMessageQuery)
		if !ok {
			return db
		}
		if dir == "asc" {
			return msgdb.Order(ent.Asc(field))
		}
		return msgdb.Order(ent.Desc(field))
	}
}


var _ interfaces.MessageServiceI = (*MessageService)(nil)