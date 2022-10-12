/*
	MIT License

	Copyright (c) 2021 Justin Hammond

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/

// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"
	"time"

	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbapp"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbfilter"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbgroup"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessage"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessagefields"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportinstances"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportrecipients"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbusermetadata"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/schema"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	dbappMixin := schema.DbApp{}.Mixin()
	dbapp.Policy = privacy.NewPolicies(dbappMixin[0], dbappMixin[1], schema.DbApp{})
	dbapp.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbapp.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbappMixinHooks1 := dbappMixin[1].Hooks()

	dbapp.Hooks[1] = dbappMixinHooks1[0]

	dbapp.Hooks[2] = dbappMixinHooks1[1]
	dbappFields := schema.DbApp{}.Fields()
	_ = dbappFields
	// dbappDescName is the schema descriptor for Name field.
	dbappDescName := dbappFields[0].Descriptor()
	// dbapp.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	dbapp.NameValidator = func() func(string) error {
		validators := dbappDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Name string) error {
			for _, fn := range fns {
				if err := fn(_Name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// dbappDescDescription is the schema descriptor for Description field.
	dbappDescDescription := dbappFields[2].Descriptor()
	// dbapp.DescriptionValidator is a validator for the "Description" field. It is called by the builders before save.
	dbapp.DescriptionValidator = func() func(string) error {
		validators := dbappDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Description string) error {
			for _, fn := range fns {
				if err := fn(_Description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// dbappDescIcon is the schema descriptor for icon field.
	dbappDescIcon := dbappFields[3].Descriptor()
	// dbapp.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	dbapp.IconValidator = dbappDescIcon.Validators[0].(func(string) error)
	// dbappDescURL is the schema descriptor for url field.
	dbappDescURL := dbappFields[4].Descriptor()
	// dbapp.URLValidator is a validator for the "url" field. It is called by the builders before save.
	dbapp.URLValidator = dbappDescURL.Validators[0].(func(string) error)
	dbfilterMixin := schema.DbFilter{}.Mixin()
	dbfilter.Policy = privacy.NewPolicies(dbfilterMixin[0], dbfilterMixin[1], schema.DbFilter{})
	dbfilter.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbfilter.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbfilterMixinHooks1 := dbfilterMixin[1].Hooks()

	dbfilter.Hooks[1] = dbfilterMixinHooks1[0]

	dbfilter.Hooks[2] = dbfilterMixinHooks1[1]
	dbfilterFields := schema.DbFilter{}.Fields()
	_ = dbfilterFields
	// dbfilterDescName is the schema descriptor for Name field.
	dbfilterDescName := dbfilterFields[0].Descriptor()
	// dbfilter.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	dbfilter.NameValidator = dbfilterDescName.Validators[0].(func(string) error)
	// dbfilterDescEnabled is the schema descriptor for Enabled field.
	dbfilterDescEnabled := dbfilterFields[3].Descriptor()
	// dbfilter.DefaultEnabled holds the default value on creation for the Enabled field.
	dbfilter.DefaultEnabled = dbfilterDescEnabled.Default.(bool)
	// dbfilterDescFilterImpl is the schema descriptor for FilterImpl field.
	dbfilterDescFilterImpl := dbfilterFields[4].Descriptor()
	// dbfilter.FilterImplValidator is a validator for the "FilterImpl" field. It is called by the builders before save.
	dbfilter.FilterImplValidator = dbfilterDescFilterImpl.Validators[0].(func(string) error)
	dbgroupMixin := schema.DbGroup{}.Mixin()
	dbgroup.Policy = privacy.NewPolicies(dbgroupMixin[0], dbgroupMixin[1], schema.DbGroup{})
	dbgroup.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbgroup.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbgroupMixinHooks1 := dbgroupMixin[1].Hooks()

	dbgroup.Hooks[1] = dbgroupMixinHooks1[0]

	dbgroup.Hooks[2] = dbgroupMixinHooks1[1]
	dbgroupFields := schema.DbGroup{}.Fields()
	_ = dbgroupFields
	// dbgroupDescName is the schema descriptor for Name field.
	dbgroupDescName := dbgroupFields[0].Descriptor()
	// dbgroup.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	dbgroup.NameValidator = dbgroupDescName.Validators[0].(func(string) error)
	dbmessageMixin := schema.DbMessage{}.Mixin()
	dbmessage.Policy = privacy.NewPolicies(dbmessageMixin[0], dbmessageMixin[1], schema.DbMessage{})
	dbmessage.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbmessage.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbmessageMixinHooks1 := dbmessageMixin[1].Hooks()

	dbmessage.Hooks[1] = dbmessageMixinHooks1[0]

	dbmessage.Hooks[2] = dbmessageMixinHooks1[1]
	dbmessageFields := schema.DbMessage{}.Fields()
	_ = dbmessageFields
	// dbmessageDescMessage is the schema descriptor for Message field.
	dbmessageDescMessage := dbmessageFields[1].Descriptor()
	// dbmessage.MessageValidator is a validator for the "Message" field. It is called by the builders before save.
	dbmessage.MessageValidator = func() func(string) error {
		validators := dbmessageDescMessage.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Message string) error {
			for _, fn := range fns {
				if err := fn(_Message); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// dbmessageDescTopic is the schema descriptor for Topic field.
	dbmessageDescTopic := dbmessageFields[3].Descriptor()
	// dbmessage.TopicValidator is a validator for the "Topic" field. It is called by the builders before save.
	dbmessage.TopicValidator = dbmessageDescTopic.Validators[0].(func(string) error)
	// dbmessageDescSeverity is the schema descriptor for Severity field.
	dbmessageDescSeverity := dbmessageFields[4].Descriptor()
	// dbmessage.DefaultSeverity holds the default value on creation for the Severity field.
	dbmessage.DefaultSeverity = dbmessageDescSeverity.Default.(int)
	// dbmessage.SeverityValidator is a validator for the "Severity" field. It is called by the builders before save.
	dbmessage.SeverityValidator = func() func(int) error {
		validators := dbmessageDescSeverity.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(_Severity int) error {
			for _, fn := range fns {
				if err := fn(_Severity); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// dbmessageDescTimestamp is the schema descriptor for Timestamp field.
	dbmessageDescTimestamp := dbmessageFields[5].Descriptor()
	// dbmessage.DefaultTimestamp holds the default value on creation for the Timestamp field.
	dbmessage.DefaultTimestamp = dbmessageDescTimestamp.Default.(time.Time)
	// dbmessageDescID is the schema descriptor for id field.
	dbmessageDescID := dbmessageFields[0].Descriptor()
	// dbmessage.DefaultID holds the default value on creation for the id field.
	dbmessage.DefaultID = dbmessageDescID.Default.(func() uuid.UUID)
	dbmessagefieldsMixin := schema.DbMessageFields{}.Mixin()
	dbmessagefields.Policy = privacy.NewPolicies(dbmessagefieldsMixin[0], dbmessagefieldsMixin[1], schema.DbMessageFields{})
	dbmessagefields.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbmessagefields.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbmessagefieldsMixinHooks1 := dbmessagefieldsMixin[1].Hooks()

	dbmessagefields.Hooks[1] = dbmessagefieldsMixinHooks1[0]

	dbmessagefields.Hooks[2] = dbmessagefieldsMixinHooks1[1]
	dbmessagefieldsFields := schema.DbMessageFields{}.Fields()
	_ = dbmessagefieldsFields
	// dbmessagefieldsDescName is the schema descriptor for Name field.
	dbmessagefieldsDescName := dbmessagefieldsFields[0].Descriptor()
	// dbmessagefields.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	dbmessagefields.NameValidator = dbmessagefieldsDescName.Validators[0].(func(string) error)
	// dbmessagefieldsDescValue is the schema descriptor for Value field.
	dbmessagefieldsDescValue := dbmessagefieldsFields[1].Descriptor()
	// dbmessagefields.ValueValidator is a validator for the "Value" field. It is called by the builders before save.
	dbmessagefields.ValueValidator = dbmessagefieldsDescValue.Validators[0].(func(string) error)
	dbtransportinstancesMixin := schema.DbTransportInstances{}.Mixin()
	dbtransportinstances.Policy = privacy.NewPolicies(dbtransportinstancesMixin[0], dbtransportinstancesMixin[1], schema.DbTransportInstances{})
	dbtransportinstances.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbtransportinstances.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbtransportinstancesMixinHooks1 := dbtransportinstancesMixin[1].Hooks()

	dbtransportinstances.Hooks[1] = dbtransportinstancesMixinHooks1[0]

	dbtransportinstances.Hooks[2] = dbtransportinstancesMixinHooks1[1]
	dbtransportinstancesFields := schema.DbTransportInstances{}.Fields()
	_ = dbtransportinstancesFields
	// dbtransportinstancesDescName is the schema descriptor for Name field.
	dbtransportinstancesDescName := dbtransportinstancesFields[0].Descriptor()
	// dbtransportinstances.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	dbtransportinstances.NameValidator = dbtransportinstancesDescName.Validators[0].(func(string) error)
	// dbtransportinstancesDescConfig is the schema descriptor for Config field.
	dbtransportinstancesDescConfig := dbtransportinstancesFields[2].Descriptor()
	// dbtransportinstances.ConfigValidator is a validator for the "Config" field. It is called by the builders before save.
	dbtransportinstances.ConfigValidator = dbtransportinstancesDescConfig.Validators[0].(func(string) error)
	// dbtransportinstancesDescTransportProvider is the schema descriptor for TransportProvider field.
	dbtransportinstancesDescTransportProvider := dbtransportinstancesFields[3].Descriptor()
	// dbtransportinstances.TransportProviderValidator is a validator for the "TransportProvider" field. It is called by the builders before save.
	dbtransportinstances.TransportProviderValidator = dbtransportinstancesDescTransportProvider.Validators[0].(func(string) error)
	dbtransportrecipientsMixin := schema.DbTransportRecipients{}.Mixin()
	dbtransportrecipients.Policy = privacy.NewPolicies(dbtransportrecipientsMixin[0], dbtransportrecipientsMixin[1], schema.DbTransportRecipients{})
	dbtransportrecipients.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbtransportrecipients.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbtransportrecipientsMixinHooks1 := dbtransportrecipientsMixin[1].Hooks()

	dbtransportrecipients.Hooks[1] = dbtransportrecipientsMixinHooks1[0]

	dbtransportrecipients.Hooks[2] = dbtransportrecipientsMixinHooks1[1]
	dbtransportrecipientsFields := schema.DbTransportRecipients{}.Fields()
	_ = dbtransportrecipientsFields
	// dbtransportrecipientsDescName is the schema descriptor for Name field.
	dbtransportrecipientsDescName := dbtransportrecipientsFields[0].Descriptor()
	// dbtransportrecipients.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	dbtransportrecipients.NameValidator = dbtransportrecipientsDescName.Validators[0].(func(string) error)
	dbuserMixin := schema.DbUser{}.Mixin()
	dbuser.Policy = privacy.NewPolicies(dbuserMixin[0], dbuserMixin[1], schema.DbUser{})
	dbuser.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbuser.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbuserMixinHooks1 := dbuserMixin[1].Hooks()

	dbuser.Hooks[1] = dbuserMixinHooks1[0]

	dbuser.Hooks[2] = dbuserMixinHooks1[1]
	dbuserFields := schema.DbUser{}.Fields()
	_ = dbuserFields
	// dbuserDescEmail is the schema descriptor for Email field.
	dbuserDescEmail := dbuserFields[0].Descriptor()
	// dbuser.EmailValidator is a validator for the "Email" field. It is called by the builders before save.
	dbuser.EmailValidator = func() func(string) error {
		validators := dbuserDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Email string) error {
			for _, fn := range fns {
				if err := fn(_Email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// dbuserDescName is the schema descriptor for Name field.
	dbuserDescName := dbuserFields[1].Descriptor()
	// dbuser.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	dbuser.NameValidator = dbuserDescName.Validators[0].(func(string) error)
	dbusermetadataMixin := schema.DbUserMetaData{}.Mixin()
	dbusermetadata.Policy = privacy.NewPolicies(dbusermetadataMixin[0], dbusermetadataMixin[1], schema.DbUserMetaData{})
	dbusermetadata.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := dbusermetadata.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	dbusermetadataMixinHooks1 := dbusermetadataMixin[1].Hooks()

	dbusermetadata.Hooks[1] = dbusermetadataMixinHooks1[0]

	dbusermetadata.Hooks[2] = dbusermetadataMixinHooks1[1]
	dbusermetadataFields := schema.DbUserMetaData{}.Fields()
	_ = dbusermetadataFields
	// dbusermetadataDescName is the schema descriptor for Name field.
	dbusermetadataDescName := dbusermetadataFields[0].Descriptor()
	// dbusermetadata.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	dbusermetadata.NameValidator = dbusermetadataDescName.Validators[0].(func(string) error)
	// dbusermetadataDescValue is the schema descriptor for Value field.
	dbusermetadataDescValue := dbusermetadataFields[1].Descriptor()
	// dbusermetadata.ValueValidator is a validator for the "Value" field. It is called by the builders before save.
	dbusermetadata.ValueValidator = dbusermetadataDescValue.Validators[0].(func(string) error)
	tenantMixin := schema.Tenant{}.Mixin()
	tenant.Policy = privacy.NewPolicies(tenantMixin[0], schema.Tenant{})
	tenant.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := tenant.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescName is the schema descriptor for name field.
	tenantDescName := tenantFields[0].Descriptor()
	// tenant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tenant.NameValidator = tenantDescName.Validators[0].(func(string) error)
}

const (
	Version = "v0.11.2"                                         // Version of ent codegen.
	Sum     = "h1:UM2/BUhF2FfsxPHRxLjQbhqJNaDdVlOwNIAMLs2jyto=" // Sum of ent codegen.
)