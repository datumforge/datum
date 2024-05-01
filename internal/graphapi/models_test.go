package graphapi_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	mock_fga "github.com/datumforge/fgax/mockery"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

type OrganizationBuilder struct {
	client *client

	// Fields
	Name        string
	DisplayName string
	Description *string
	OrgID       string
	ParentOrgID string
	PersonalOrg bool
}

type OrganizationCleanup struct {
	client *client

	// Fields
	OrgID string
}

type GroupBuilder struct {
	client *client

	// Fields
	Name  string
	Owner string
}

type GroupCleanup struct {
	client *client

	// Fields
	GroupID string
}

type UserBuilder struct {
	client *client

	// Fields
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserCleanup struct {
	client *client

	// Fields
	UserID string
}

type TFASettingBuilder struct {
	client *client
}

type OrgMemberBuilder struct {
	client *client

	// Fields
	UserID string
	OrgID  string
	Role   string
}

type OrgMemberCleanup struct {
	client *client

	// Fields
	ID string
}

type GroupMemberBuilder struct {
	client *client

	// Fields
	UserID  string
	GroupID string
	Role    string
}

type GroupMemberCleanup struct {
	client *client

	// Fields
	ID string
}

type InviteBuilder struct {
	client *client

	// Fields
	Recipient string
	OrgID     string
	Role      string
}

type InviteCleanup struct {
	client *client

	// Fields
	ID string
}

type PersonalAccessTokenBuilder struct {
	client *client

	// Fields
	Name           string
	Token          string
	Abilities      []string
	Description    string
	ExpiresAt      time.Time
	OwnerID        string
	OrganizationID string
}

// MustNew organization builder is used to create, without authz checks, orgs in the database
func (o *OrganizationBuilder) MustNew(ctx context.Context, t *testing.T) *ent.Organization {
	if !o.PersonalOrg {
		// mock writes
		mock_fga.WriteOnce(t, o.client.fga)
	}

	// no auth, so allow policy
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

	m := o.client.db.Organization.Create().SetName(o.Name).SetDescription(*o.Description).SetDisplayName(o.DisplayName).SetPersonalOrg(o.PersonalOrg)

	if o.ParentOrgID != "" {
		mock_fga.ListAny(t, o.client.fga, []string{fmt.Sprintf("organization:%s", o.ParentOrgID)})

		m.SetParentID(o.ParentOrgID)
	}

	org := m.SaveX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(o.client.fga)

	return org
}

// MustDelete is used to cleanup, without authz checks, orgs in the database
func (o *OrganizationCleanup) MustDelete(ctx context.Context, t *testing.T) {
	// mock checks
	mock_fga.ListAny(t, o.client.fga, []string{fmt.Sprintf("organization:%s", o.OrgID)})

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	o.client.db.Organization.DeleteOneID(o.OrgID).ExecX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(o.client.fga)
}

// MustNew user builder is used to create, without authz checks, users in the database
func (u *UserBuilder) MustNew(ctx context.Context, t *testing.T) *ent.User {
	// mock writes
	mock_fga.WriteOnce(t, u.client.fga)

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
	userSetting := u.client.db.UserSetting.Create().SaveX(ctx)

	user := u.client.db.User.Create().
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetSetting(userSetting).
		SaveX(ctx)

	user.Edges.Setting.DefaultOrg(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(u.client.fga)

	return user
}

// MustDelete is used to cleanup, without authz checks, users in the database
func (u *UserCleanup) MustDelete(ctx context.Context, t *testing.T) {
	// mock checks
	mock_fga.ListAny(t, u.client.fga, []string{})

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	u.client.db.User.DeleteOneID(u.UserID).ExecX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(u.client.fga)
}

// MustNew user builder is used to create, without authz checks, org members in the database
func (tf *TFASettingBuilder) MustNew(ctx context.Context, t *testing.T, userID string) *ent.TFASetting {
	return tf.client.db.TFASetting.Create().
		SetTotpAllowed(true).
		SetOwnerID(userID).
		SaveX(ctx)
}

// MustNew user builder is used to create, without authz checks, org members in the database
func (om *OrgMemberBuilder) MustNew(ctx context.Context, t *testing.T) *ent.OrgMembership {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if om.OrgID == "" {
		org := (&OrganizationBuilder{client: om.client}).MustNew(ctx, t)
		om.OrgID = org.ID
	}

	if om.UserID == "" {
		user := (&UserBuilder{client: om.client}).MustNew(ctx, t)
		om.UserID = user.ID
	}

	role := enums.ToRole(om.Role)
	if role == &enums.RoleInvalid {
		role = &enums.RoleMember
	}

	// mock writes
	mock_fga.ListAny(t, om.client.fga, []string{fmt.Sprintf("organization:%s", om.OrgID)})
	mock_fga.WriteOnce(t, om.client.fga)

	orgMembers := om.client.db.OrgMembership.Create().
		SetUserID(om.UserID).
		SetOrganizationID(om.OrgID).
		SetRole(*role).
		SaveX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(om.client.fga)

	return orgMembers
}

// MustDelete is used to cleanup, without authz checks, org members in the database
func (om *OrgMemberCleanup) MustDelete(ctx context.Context, t *testing.T) {
	// mock writes
	mock_fga.WriteOnce(t, om.client.fga)

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	om.client.db.OrgMembership.DeleteOneID(om.ID).ExecX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(om.client.fga)
}

// MustNew group builder is used to create, without authz checks, groups in the database
func (g *GroupBuilder) MustNew(ctx context.Context, t *testing.T) *ent.Group {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if g.Name == "" {
		g.Name = gofakeit.AppName()
	}

	// create owner if not provided
	owner := g.Owner

	if g.Owner == "" {
		owner = testPersonalOrgID
	}

	// mock writes
	mock_fga.WriteAny(t, g.client.fga)

	mock_fga.ListAny(t, g.client.fga, []string{fmt.Sprintf("group:%s", owner)})

	group := g.client.db.Group.Create().SetName(g.Name).SetOwnerID(owner).SaveX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(g.client.fga)

	return group
}

// MustDelete is used to cleanup, without authz checks, groups in the database
func (g *GroupCleanup) MustDelete(ctx context.Context, t *testing.T) {
	mock_fga.ClearMocks(g.client.fga)

	// mock writes
	mock_fga.ReadAny(t, g.client.fga)
	mock_fga.ListAny(t, g.client.fga, []string{fmt.Sprintf("group:%s", g.GroupID)})

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	g.client.db.Group.DeleteOneID(g.GroupID).ExecX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(g.client.fga)
}

// MustNew group builder is used to create, without authz checks, groups in the database
func (i *InviteBuilder) MustNew(ctx context.Context, t *testing.T) *ent.Invite {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	// create owner if not provided
	orgID := i.OrgID

	if orgID == "" {
		org := (&OrganizationBuilder{client: i.client}).MustNew(ctx, t)
		orgID = org.ID
	}

	// create user if not provided
	rec := i.Recipient

	if rec == "" {
		rec = gofakeit.Email()
	}

	// mock check
	mock_fga.ListAny(t, i.client.fga, []string{fmt.Sprintf("organization:%s", orgID)})

	inviteQuery := i.client.db.Invite.Create().
		SetOwnerID(orgID).
		SetRecipient(rec)

	if i.Role != "" {
		inviteQuery.SetRole(*enums.ToRole(i.Role))
	}

	invite := inviteQuery.SaveX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(i.client.fga)

	return invite
}

// MustDelete is used to cleanup, without authz checks, invites in the database
func (i *InviteCleanup) MustDelete(ctx context.Context, t *testing.T) {
	// mock writes
	mock_fga.ReadAny(t, i.client.fga)

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	i.client.db.Invite.DeleteOneID(i.ID).ExecX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(i.client.fga)
}

// MustNew group builder is used to create, without authz checks, personal access tokens in the database
func (pat *PersonalAccessTokenBuilder) MustNew(ctx context.Context, t *testing.T) *ent.PersonalAccessToken {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if pat.Name == "" {
		pat.Name = gofakeit.AppName()
	}

	if pat.Description == "" {
		pat.Description = gofakeit.HipsterSentence(5)
	}

	if pat.OwnerID == "" {
		owner := (&UserBuilder{client: pat.client}).MustNew(ctx, t)
		pat.OwnerID = owner.ID
	}

	if pat.OrganizationID == "" {
		org := (&OrganizationBuilder{client: pat.client}).MustNew(ctx, t)
		pat.OrganizationID = org.ID
	}

	token := pat.client.db.PersonalAccessToken.Create().
		SetName(pat.Name).
		SetOwnerID(pat.OwnerID).
		SetDescription(pat.Description).
		SetExpiresAt(pat.ExpiresAt).
		AddOrganizationIDs(pat.OrganizationID).
		SaveX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(pat.client.fga)

	return token
}

// MustNew user builder is used to create, without authz checks, group members in the database
func (gm *GroupMemberBuilder) MustNew(ctx context.Context, t *testing.T) *ent.GroupMembership {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	if gm.GroupID == "" {
		group := (&GroupBuilder{client: gm.client}).MustNew(ctx, t)
		gm.GroupID = group.ID
	}

	if gm.UserID == "" {
		user := (&UserBuilder{client: gm.client}).MustNew(ctx, t)
		gm.UserID = user.ID
	}

	// mock writes
	mock_fga.ListAny(t, gm.client.fga, []string{fmt.Sprintf("organization:%s", testPersonalOrgID)})
	mock_fga.WriteOnce(t, gm.client.fga)

	groupMember := gm.client.db.GroupMembership.Create().
		SetUserID(gm.UserID).
		SetGroupID(gm.GroupID).
		SaveX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(gm.client.fga)

	return groupMember
}

// MustDelete is used to cleanup, without authz checks, group members in the database
func (gm *GroupMemberCleanup) MustDelete(ctx context.Context, t *testing.T) {
	// mock writes
	mock_fga.WriteOnce(t, gm.client.fga)

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	gm.client.db.GroupMembership.DeleteOneID(gm.ID).ExecX(ctx)

	// clear mocks before going to tests
	mock_fga.ClearMocks(gm.client.fga)
}
