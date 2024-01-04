package graphapi_test

import (
	"context"
	"time"

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
	PersonalOrg bool
}

type OrganizationCleanup struct {
	OrgID string
}

type GroupBuilder struct {
	Name  string
	Owner string
}

type GroupCleanup struct {
	GroupID string
}

type UserBuilder struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserCleanup struct {
	UserID string
}

type PersonalAccessTokenBuilder struct {
	Name        string
	Token       string
	Abilities   []string
	Description string
	ExpiresAt   time.Time
	OwnerID     string
}

// MustNew organization builder is used to create, without authz checks, orgs in the database
func (o *OrganizationBuilder) MustNew(ctx context.Context) *generated.Organization {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if o.Name == "" {
		o.Name = gofakeit.AppName()
	}

	if o.DisplayName == "" {
		o.DisplayName = gofakeit.LetterN(40)
	}

	if o.Description == nil {
		desc := gofakeit.HipsterSentence(10)
		o.Description = &desc
	}

	m := EntClient.Organization.Create().SetName(o.Name).SetDescription(*o.Description).SetDisplayName(o.DisplayName).SetPersonalOrg(o.PersonalOrg)

	if o.ParentOrgID != "" {
		m.SetParentID(o.ParentOrgID)
	}

	return m.SaveX(ctx)
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

// MustDelete is used to cleanup, without authz checks, orgs in the database
func (u *UserCleanup) MustDelete(ctx context.Context) {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	EntClient.User.DeleteOneID(u.UserID).ExecX(ctx)
}

// MustNew group builder is used to create, without authz checks, groups in the database
func (g *GroupBuilder) MustNew(ctx context.Context) *generated.Group {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if g.Name == "" {
		g.Name = gofakeit.AppName()
	}

	// create owner if not provided
	owner := g.Owner

	if g.Owner == "" {
		org := (&OrganizationBuilder{}).MustNew(ctx)
		owner = org.ID
	}

	return EntClient.Group.Create().SetName(g.Name).SetOwnerID(owner).SaveX(ctx)
}

// MustDelete is used to cleanup, without authz checks, groups in the database
func (g *GroupCleanup) MustDelete(ctx context.Context) {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	EntClient.Group.DeleteOneID(g.GroupID).ExecX(ctx)
}

func (t *PersonalAccessTokenBuilder) MustNew(ctx context.Context) *generated.PersonalAccessToken {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if t.Name == "" {
		t.Name = gofakeit.AppName()
	}

	if t.Description == "" {
		t.Description = gofakeit.HipsterSentence(5)
	}

	if t.OwnerID == "" {
		owner := (&UserBuilder{}).MustNew(ctx)
		t.OwnerID = owner.ID
	}

	return EntClient.PersonalAccessToken.Create().
		SetName(t.Name).
		SetOwnerID(t.OwnerID).
		SetToken(t.Token).
		SetDescription(t.Description).
		SetExpiresAt(t.ExpiresAt).
		SaveX(ctx)
}
