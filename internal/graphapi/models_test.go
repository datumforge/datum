package graphapi_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

type OrganizationBuilder struct {
	Name        string
	DisplayName string
	Description *string
	OrgID       string
	ParentOrgID string
}

type OrganizationCleanup struct {
	OrgID string
}

type UserBuilder struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// MustNew organization builder is used to create, without authz checks, orgs in the database
func (o *OrganizationBuilder) MustNew(ctx context.Context) *generated.Organization {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if o.Name == "" {
		o.Name = gofakeit.AppName()
	}

	if o.DisplayName == "" {
		o.Name = gofakeit.LetterN(40)
	}

	if o.Description == nil {
		desc := gofakeit.HipsterSentence(10)
		o.Description = &desc
	}

	m := EntClient.Organization.Create().SetName(o.Name).SetDescription(*o.Description)

	if o.ParentOrgID != "" {
		m.SetParentID(o.ParentOrgID).SaveX(ctx)
	}

	return EntClient.Organization.Create().SetName(o.Name).SetDescription(*o.Description).SaveX(ctx)
}

// MustDelete is used to cleanup, without authz checks, orgs in the database
func (o *OrganizationCleanup) MustDelete(ctx context.Context) {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	EntClient.Organization.DeleteOneID(o.OrgID).ExecX(ctx)
}

// MustNew user builder is used to create, without authz checks, users in the database
func (u *UserBuilder) MustNew(ctx context.Context) *generated.User {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if u.FirstName == "" {
		u.FirstName = gofakeit.FirstName()
	}

	if u.LastName == "" {
		u.LastName = gofakeit.LastName()
	}

	if u.Email == "" {
		u.Email = gofakeit.Email()
	}

	if u.Password == "" {
		u.Password = gofakeit.Password(true, true, true, true, false, 20)
	}

	// create user setting
	userSetting := EntClient.UserSetting.Create().SaveX(ctx)

	return EntClient.User.Create().
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetSetting(userSetting).
		SaveX(ctx)
}
