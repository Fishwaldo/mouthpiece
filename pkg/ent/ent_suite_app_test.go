package ent_test

import (
	//	"context"
	//	"fmt"

	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/app"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
)

var _ = Describe("Applications: ", func() {
	When("Creating", func() {
		DescribeTable("Some Apps",
			func(name string, description string, status app.Status, Icon string, URL string, 
				expectedName string, expectedDescription string, expectedStatus app.Status, expectedIcon string, expectedURL string) {
				newapp, err := eClient.App.Create().
					SetName(name).
					SetStatus(status).
					SetDescription(description).
					SetIcon(Icon).
					SetURL(URL).
					Save(ctx)
				Expect(err).To(BeNil())
				Expect(newapp).ToNot(BeNil())
				Expect(newapp.Name).To(Equal(expectedName))
				Expect(newapp.Status).To(Equal(expectedStatus))
				Expect(newapp.Description).To(Equal(expectedDescription))
				Expect(newapp.Icon).To(Equal(expectedIcon))
				Expect(newapp.URL).To(Equal(expectedURL))
			},
		Entry("a Normal App", "TestApp", "Test Description", app.StatusEnabled, "https://example.com/icon.png", "https://example.com", "TestApp", "Test Description", app.StatusEnabled, "https://example.com/icon.png", "https://example.com"),
		Entry("a Disabled App", "TestApp2", "Test Description", app.StatusDisabled, "https://example.com/icon.png", "https://example.com", "TestApp2", "Test Description", app.StatusDisabled, "https://example.com/icon.png", "https://example.com"),
		)
	})
	When("Creating a New App From Invalid Params Params", func() { 		
		It("should fail", func() {
			_, err := eClient.App.Create().
				SetName("").
				SetStatus(app.StatusEnabled).
				SetDescription("Test Description").
				SetIcon("https://example.com/icon.png").
				SetURL("https://example.com").
				Save(ctx)
			Expect(err).ToNot(BeNil())
		})
	})
	When("Getting a Disabled App", func() {
		It("should be disabled", func() {
			_, _ = eClient.App.Query().All(ctx)
			newapp, err := eClient.App.Query().Where(app.Name("TestApp2")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			Expect(newapp.Status).To(Equal(app.StatusDisabled))
		})
	})
	When("Getting a Enabled App", func() {
		It("should be enabled", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			Expect(newapp.Status).To(Equal(app.StatusEnabled))
		})
	})
	When("Creating a New App From a Struct", func() {
		It("Can be saved to the Database", func() {
			newapp := &ent.App{
				Name:        "TestApp3",
				Description: "Test Description",
				Status:      app.StatusEnabled,
				Icon:        "https://example.com/icon.png",
				URL:         "https://example.com",
			}
			app, err := eClient.App.Create().SetAppFromStruct(newapp).Save(ctx)
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			Expect(app.Name).To(Equal(newapp.Name))
			Expect(app.Status).To(Equal(newapp.Status))
			Expect(app.Description).To(Equal(newapp.Description))
			Expect(app.Icon).To(Equal(newapp.Icon))
			Expect(app.URL).To(Equal(newapp.URL))
		})
		It("with a invlaid field it can not be saved to the Database", func() {
			newapp := &ent.App{
				Name:        "TestApp3",
				Description: "Test Description",
				Status:      app.StatusEnabled,
				Icon:        "https//example.com/icon.png",
				URL:         "https://example.com",
			}
			app, err := eClient.App.Create().SetAppFromStruct(newapp).Save(ctx)
			Expect(err).To(Not(BeNil()))
			Expect(app).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("failed on the 'url' tag"))
		})
	})
	When("Getting a existing app", func() {
		It("can modify name", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			newapp, err = eClient.App.UpdateOne(newapp).SetName("TestApp4").Save(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			Expect(newapp.Name).To(Equal("TestApp4"))
		})
		It("can modify description", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp4")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			newapp, err = eClient.App.UpdateOne(newapp).SetDescription("Test Description2").Save(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			Expect(newapp.Description).To(Equal("Test Description2"))
		})
		It("can modify status", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp4")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			newapp, err = eClient.App.UpdateOne(newapp).SetStatus(app.StatusDisabled).Save(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			Expect(newapp.Status).To(Equal(app.StatusDisabled))
		})
		It("can modify icon", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp4")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			newapp, err = eClient.App.UpdateOne(newapp).SetIcon("https://example.com/icon2.png").Save(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			Expect(newapp.Icon).To(Equal("https://example.com/icon2.png"))
		})
		It("can modify url", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp4")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			newapp, err = eClient.App.UpdateOne(newapp).SetURL("https://example.com").Save(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			Expect(newapp.URL).To(Equal("https://example.com"))
		})
		It("Can be deleted", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp4")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			err = eClient.App.DeleteOne(newapp).Exec(ctx)
			Expect(err).To(BeNil())
			newapp, err = eClient.App.Query().Where(app.Name("TestApp4")).Only(ctx)
			Expect(ent.IsNotFound(err)).To(BeTrue())
			Expect(newapp).To(BeNil())
		})
		It("has no Messages", func() {
			newapp, err := eClient.App.Query().WithMessages().Where(app.Name("TestApp2")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).To(Not(BeNil()))
			msg, err := newapp.Edges.MessagesOrErr()
			Expect(err).To(BeNil())
			Expect(len(msg)).To(Equal(0))
		})
		It("Has no Filters", func() {
			newapp, err := eClient.App.Query().WithFilters().Where(app.Name("TestApp2")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).To(Not(BeNil()))
			flt, err := newapp.Edges.FiltersOrErr()
			Expect(err).To(BeNil())
			Expect(len(flt)).To(Equal(0))
		})
		It("Has no Groups", func() {
			newapp, err := eClient.App.Query().WithGroups().Where(app.Name("TestApp2")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).To(Not(BeNil()))
			grp, err := newapp.Edges.GroupsOrErr()
			Expect(err).To(BeNil())
			Expect(len(grp)).To(Equal(0))
		})
		It("Has No Transports",func() {
			newapp, err := eClient.App.Query().WithTransportRecipients().Where(app.Name("TestApp2")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).To(Not(BeNil()))
			trp, err := newapp.Edges.TransportRecipientsOrErr()
			Expect(err).To(BeNil())
			Expect(len(trp)).To(Equal(0))
		})
		It("can not modify tenant id", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp2")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			newtnt := *rules.FromContextGetTenant(ctx)
			newtnt.ID = newtnt.ID + 1
			newapp, err = eClient.App.UpdateOne(newapp).SetTenant(&newtnt).Save(ctx)
			Expect(err).ToNot(BeNil())
			Expect(newapp).To(BeNil())
		})
		It("Can Modify the Tenant ID if Global Admin", func() {
			newapp, err := eClient.App.Query().Where(app.Name("TestApp2")).Only(ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			viewer := rules.FromContext(ctx)
			uv := viewer.(*rules.UserViewer)
			newctx := rules.NewContext(context.Background(), rules.UserViewer{T: uv.T, Role: rules.GlobalAdmin})
			tnt, err := eClient.Tenant.Create().SetName("Default2").Save(newctx)
			Expect(err).To(BeNil())
			Expect(tnt).ToNot(BeNil())
			newapp1, err1 := eClient.App.UpdateOne(newapp).SetTenant(tnt).Save(newctx)
			Expect(err1).To(BeNil())
			Expect(newapp1).ToNot(BeNil())
			Expect(newapp1.TenantID).To(Equal(tnt.ID))
		})
	})

})	
