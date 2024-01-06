// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtokenhistory"
	"github.com/datumforge/datum/internal/ent/generated/entitlementhistory"
	"github.com/datumforge/datum/internal/ent/generated/grouphistory"
	"github.com/datumforge/datum/internal/ent/generated/groupsettinghistory"
	"github.com/datumforge/datum/internal/ent/generated/integrationhistory"
	"github.com/datumforge/datum/internal/ent/generated/oauthproviderhistory"
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootokenhistory"
	"github.com/datumforge/datum/internal/ent/generated/organizationhistory"
	"github.com/datumforge/datum/internal/ent/generated/organizationsettinghistory"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettokenhistory"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstokenhistory"
	"github.com/datumforge/datum/internal/ent/generated/sessionhistory"
	"github.com/datumforge/datum/internal/ent/generated/userhistory"
	"github.com/datumforge/datum/internal/ent/generated/usersettinghistory"
)

func (evt *EmailVerificationToken) History() *EmailVerificationTokenHistoryQuery {
	historyClient := NewEmailVerificationTokenHistoryClient(evt.config)
	return historyClient.Query().Where(emailverificationtokenhistory.Ref(evt.ID))
}

func (evth *EmailVerificationTokenHistory) Next(ctx context.Context) (*EmailVerificationTokenHistory, error) {
	client := NewEmailVerificationTokenHistoryClient(evth.config)
	return client.Query().
		Where(
			emailverificationtokenhistory.Ref(evth.Ref),
			emailverificationtokenhistory.HistoryTimeGT(evth.HistoryTime),
		).
		Order(emailverificationtokenhistory.ByHistoryTime()).
		First(ctx)
}

func (evth *EmailVerificationTokenHistory) Prev(ctx context.Context) (*EmailVerificationTokenHistory, error) {
	client := NewEmailVerificationTokenHistoryClient(evth.config)
	return client.Query().
		Where(
			emailverificationtokenhistory.Ref(evth.Ref),
			emailverificationtokenhistory.HistoryTimeLT(evth.HistoryTime),
		).
		Order(emailverificationtokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (evthq *EmailVerificationTokenHistoryQuery) Earliest(ctx context.Context) (*EmailVerificationTokenHistory, error) {
	return evthq.
		Order(emailverificationtokenhistory.ByHistoryTime()).
		First(ctx)
}

func (evthq *EmailVerificationTokenHistoryQuery) Latest(ctx context.Context) (*EmailVerificationTokenHistory, error) {
	return evthq.
		Order(emailverificationtokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (evthq *EmailVerificationTokenHistoryQuery) AsOf(ctx context.Context, time time.Time) (*EmailVerificationTokenHistory, error) {
	return evthq.
		Where(emailverificationtokenhistory.HistoryTimeLTE(time)).
		Order(emailverificationtokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (evth *EmailVerificationTokenHistory) Restore(ctx context.Context) (*EmailVerificationToken, error) {
	client := NewEmailVerificationTokenClient(evth.config)
	return client.
		UpdateOneID(evth.Ref).
		SetUpdatedAt(evth.UpdatedAt).
		SetUpdatedBy(evth.UpdatedBy).
		SetDeletedAt(evth.DeletedAt).
		SetDeletedBy(evth.DeletedBy).
		SetToken(evth.Token).
		SetNillableTTL(evth.TTL).
		SetEmail(evth.Email).
		SetNillableSecret(evth.Secret).
		Save(ctx)
}

func (e *Entitlement) History() *EntitlementHistoryQuery {
	historyClient := NewEntitlementHistoryClient(e.config)
	return historyClient.Query().Where(entitlementhistory.Ref(e.ID))
}

func (eh *EntitlementHistory) Next(ctx context.Context) (*EntitlementHistory, error) {
	client := NewEntitlementHistoryClient(eh.config)
	return client.Query().
		Where(
			entitlementhistory.Ref(eh.Ref),
			entitlementhistory.HistoryTimeGT(eh.HistoryTime),
		).
		Order(entitlementhistory.ByHistoryTime()).
		First(ctx)
}

func (eh *EntitlementHistory) Prev(ctx context.Context) (*EntitlementHistory, error) {
	client := NewEntitlementHistoryClient(eh.config)
	return client.Query().
		Where(
			entitlementhistory.Ref(eh.Ref),
			entitlementhistory.HistoryTimeLT(eh.HistoryTime),
		).
		Order(entitlementhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ehq *EntitlementHistoryQuery) Earliest(ctx context.Context) (*EntitlementHistory, error) {
	return ehq.
		Order(entitlementhistory.ByHistoryTime()).
		First(ctx)
}

func (ehq *EntitlementHistoryQuery) Latest(ctx context.Context) (*EntitlementHistory, error) {
	return ehq.
		Order(entitlementhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ehq *EntitlementHistoryQuery) AsOf(ctx context.Context, time time.Time) (*EntitlementHistory, error) {
	return ehq.
		Where(entitlementhistory.HistoryTimeLTE(time)).
		Order(entitlementhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (eh *EntitlementHistory) Restore(ctx context.Context) (*Entitlement, error) {
	client := NewEntitlementClient(eh.config)
	return client.
		UpdateOneID(eh.Ref).
		SetUpdatedAt(eh.UpdatedAt).
		SetUpdatedBy(eh.UpdatedBy).
		SetDeletedAt(eh.DeletedAt).
		SetDeletedBy(eh.DeletedBy).
		SetTier(eh.Tier).
		SetExternalCustomerID(eh.ExternalCustomerID).
		SetExternalSubscriptionID(eh.ExternalSubscriptionID).
		SetExpires(eh.Expires).
		SetNillableExpiresAt(eh.ExpiresAt).
		SetCancelled(eh.Cancelled).
		Save(ctx)
}

func (gr *Group) History() *GroupHistoryQuery {
	historyClient := NewGroupHistoryClient(gr.config)
	return historyClient.Query().Where(grouphistory.Ref(gr.ID))
}

func (gh *GroupHistory) Next(ctx context.Context) (*GroupHistory, error) {
	client := NewGroupHistoryClient(gh.config)
	return client.Query().
		Where(
			grouphistory.Ref(gh.Ref),
			grouphistory.HistoryTimeGT(gh.HistoryTime),
		).
		Order(grouphistory.ByHistoryTime()).
		First(ctx)
}

func (gh *GroupHistory) Prev(ctx context.Context) (*GroupHistory, error) {
	client := NewGroupHistoryClient(gh.config)
	return client.Query().
		Where(
			grouphistory.Ref(gh.Ref),
			grouphistory.HistoryTimeLT(gh.HistoryTime),
		).
		Order(grouphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ghq *GroupHistoryQuery) Earliest(ctx context.Context) (*GroupHistory, error) {
	return ghq.
		Order(grouphistory.ByHistoryTime()).
		First(ctx)
}

func (ghq *GroupHistoryQuery) Latest(ctx context.Context) (*GroupHistory, error) {
	return ghq.
		Order(grouphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ghq *GroupHistoryQuery) AsOf(ctx context.Context, time time.Time) (*GroupHistory, error) {
	return ghq.
		Where(grouphistory.HistoryTimeLTE(time)).
		Order(grouphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (gh *GroupHistory) Restore(ctx context.Context) (*Group, error) {
	client := NewGroupClient(gh.config)
	return client.
		UpdateOneID(gh.Ref).
		SetUpdatedAt(gh.UpdatedAt).
		SetUpdatedBy(gh.UpdatedBy).
		SetDeletedAt(gh.DeletedAt).
		SetDeletedBy(gh.DeletedBy).
		SetName(gh.Name).
		SetDescription(gh.Description).
		SetGravatarLogoURL(gh.GravatarLogoURL).
		SetLogoURL(gh.LogoURL).
		SetDisplayName(gh.DisplayName).
		Save(ctx)
}

func (gs *GroupSetting) History() *GroupSettingHistoryQuery {
	historyClient := NewGroupSettingHistoryClient(gs.config)
	return historyClient.Query().Where(groupsettinghistory.Ref(gs.ID))
}

func (gsh *GroupSettingHistory) Next(ctx context.Context) (*GroupSettingHistory, error) {
	client := NewGroupSettingHistoryClient(gsh.config)
	return client.Query().
		Where(
			groupsettinghistory.Ref(gsh.Ref),
			groupsettinghistory.HistoryTimeGT(gsh.HistoryTime),
		).
		Order(groupsettinghistory.ByHistoryTime()).
		First(ctx)
}

func (gsh *GroupSettingHistory) Prev(ctx context.Context) (*GroupSettingHistory, error) {
	client := NewGroupSettingHistoryClient(gsh.config)
	return client.Query().
		Where(
			groupsettinghistory.Ref(gsh.Ref),
			groupsettinghistory.HistoryTimeLT(gsh.HistoryTime),
		).
		Order(groupsettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (gshq *GroupSettingHistoryQuery) Earliest(ctx context.Context) (*GroupSettingHistory, error) {
	return gshq.
		Order(groupsettinghistory.ByHistoryTime()).
		First(ctx)
}

func (gshq *GroupSettingHistoryQuery) Latest(ctx context.Context) (*GroupSettingHistory, error) {
	return gshq.
		Order(groupsettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (gshq *GroupSettingHistoryQuery) AsOf(ctx context.Context, time time.Time) (*GroupSettingHistory, error) {
	return gshq.
		Where(groupsettinghistory.HistoryTimeLTE(time)).
		Order(groupsettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (gsh *GroupSettingHistory) Restore(ctx context.Context) (*GroupSetting, error) {
	client := NewGroupSettingClient(gsh.config)
	return client.
		UpdateOneID(gsh.Ref).
		SetUpdatedAt(gsh.UpdatedAt).
		SetUpdatedBy(gsh.UpdatedBy).
		SetDeletedAt(gsh.DeletedAt).
		SetDeletedBy(gsh.DeletedBy).
		SetVisibility(gsh.Visibility).
		SetJoinPolicy(gsh.JoinPolicy).
		SetTags(gsh.Tags).
		SetSyncToSlack(gsh.SyncToSlack).
		SetSyncToGithub(gsh.SyncToGithub).
		Save(ctx)
}

func (i *Integration) History() *IntegrationHistoryQuery {
	historyClient := NewIntegrationHistoryClient(i.config)
	return historyClient.Query().Where(integrationhistory.Ref(i.ID))
}

func (ih *IntegrationHistory) Next(ctx context.Context) (*IntegrationHistory, error) {
	client := NewIntegrationHistoryClient(ih.config)
	return client.Query().
		Where(
			integrationhistory.Ref(ih.Ref),
			integrationhistory.HistoryTimeGT(ih.HistoryTime),
		).
		Order(integrationhistory.ByHistoryTime()).
		First(ctx)
}

func (ih *IntegrationHistory) Prev(ctx context.Context) (*IntegrationHistory, error) {
	client := NewIntegrationHistoryClient(ih.config)
	return client.Query().
		Where(
			integrationhistory.Ref(ih.Ref),
			integrationhistory.HistoryTimeLT(ih.HistoryTime),
		).
		Order(integrationhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ihq *IntegrationHistoryQuery) Earliest(ctx context.Context) (*IntegrationHistory, error) {
	return ihq.
		Order(integrationhistory.ByHistoryTime()).
		First(ctx)
}

func (ihq *IntegrationHistoryQuery) Latest(ctx context.Context) (*IntegrationHistory, error) {
	return ihq.
		Order(integrationhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ihq *IntegrationHistoryQuery) AsOf(ctx context.Context, time time.Time) (*IntegrationHistory, error) {
	return ihq.
		Where(integrationhistory.HistoryTimeLTE(time)).
		Order(integrationhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ih *IntegrationHistory) Restore(ctx context.Context) (*Integration, error) {
	client := NewIntegrationClient(ih.config)
	return client.
		UpdateOneID(ih.Ref).
		SetUpdatedAt(ih.UpdatedAt).
		SetUpdatedBy(ih.UpdatedBy).
		SetDeletedAt(ih.DeletedAt).
		SetDeletedBy(ih.DeletedBy).
		SetName(ih.Name).
		SetDescription(ih.Description).
		SetKind(ih.Kind).
		Save(ctx)
}

func (op *OauthProvider) History() *OauthProviderHistoryQuery {
	historyClient := NewOauthProviderHistoryClient(op.config)
	return historyClient.Query().Where(oauthproviderhistory.Ref(op.ID))
}

func (oph *OauthProviderHistory) Next(ctx context.Context) (*OauthProviderHistory, error) {
	client := NewOauthProviderHistoryClient(oph.config)
	return client.Query().
		Where(
			oauthproviderhistory.Ref(oph.Ref),
			oauthproviderhistory.HistoryTimeGT(oph.HistoryTime),
		).
		Order(oauthproviderhistory.ByHistoryTime()).
		First(ctx)
}

func (oph *OauthProviderHistory) Prev(ctx context.Context) (*OauthProviderHistory, error) {
	client := NewOauthProviderHistoryClient(oph.config)
	return client.Query().
		Where(
			oauthproviderhistory.Ref(oph.Ref),
			oauthproviderhistory.HistoryTimeLT(oph.HistoryTime),
		).
		Order(oauthproviderhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ophq *OauthProviderHistoryQuery) Earliest(ctx context.Context) (*OauthProviderHistory, error) {
	return ophq.
		Order(oauthproviderhistory.ByHistoryTime()).
		First(ctx)
}

func (ophq *OauthProviderHistoryQuery) Latest(ctx context.Context) (*OauthProviderHistory, error) {
	return ophq.
		Order(oauthproviderhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ophq *OauthProviderHistoryQuery) AsOf(ctx context.Context, time time.Time) (*OauthProviderHistory, error) {
	return ophq.
		Where(oauthproviderhistory.HistoryTimeLTE(time)).
		Order(oauthproviderhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (oph *OauthProviderHistory) Restore(ctx context.Context) (*OauthProvider, error) {
	client := NewOauthProviderClient(oph.config)
	return client.
		UpdateOneID(oph.Ref).
		SetUpdatedAt(oph.UpdatedAt).
		SetUpdatedBy(oph.UpdatedBy).
		SetDeletedAt(oph.DeletedAt).
		SetDeletedBy(oph.DeletedBy).
		SetName(oph.Name).
		SetClientID(oph.ClientID).
		SetClientSecret(oph.ClientSecret).
		SetRedirectURL(oph.RedirectURL).
		SetScopes(oph.Scopes).
		SetAuthURL(oph.AuthURL).
		SetTokenURL(oph.TokenURL).
		SetAuthStyle(oph.AuthStyle).
		SetInfoURL(oph.InfoURL).
		Save(ctx)
}

func (oatt *OhAuthTooToken) History() *OhAuthTooTokenHistoryQuery {
	historyClient := NewOhAuthTooTokenHistoryClient(oatt.config)
	return historyClient.Query().Where(ohauthtootokenhistory.Ref(oatt.ID))
}

func (oatth *OhAuthTooTokenHistory) Next(ctx context.Context) (*OhAuthTooTokenHistory, error) {
	client := NewOhAuthTooTokenHistoryClient(oatth.config)
	return client.Query().
		Where(
			ohauthtootokenhistory.Ref(oatth.Ref),
			ohauthtootokenhistory.HistoryTimeGT(oatth.HistoryTime),
		).
		Order(ohauthtootokenhistory.ByHistoryTime()).
		First(ctx)
}

func (oatth *OhAuthTooTokenHistory) Prev(ctx context.Context) (*OhAuthTooTokenHistory, error) {
	client := NewOhAuthTooTokenHistoryClient(oatth.config)
	return client.Query().
		Where(
			ohauthtootokenhistory.Ref(oatth.Ref),
			ohauthtootokenhistory.HistoryTimeLT(oatth.HistoryTime),
		).
		Order(ohauthtootokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (oatthq *OhAuthTooTokenHistoryQuery) Earliest(ctx context.Context) (*OhAuthTooTokenHistory, error) {
	return oatthq.
		Order(ohauthtootokenhistory.ByHistoryTime()).
		First(ctx)
}

func (oatthq *OhAuthTooTokenHistoryQuery) Latest(ctx context.Context) (*OhAuthTooTokenHistory, error) {
	return oatthq.
		Order(ohauthtootokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (oatthq *OhAuthTooTokenHistoryQuery) AsOf(ctx context.Context, time time.Time) (*OhAuthTooTokenHistory, error) {
	return oatthq.
		Where(ohauthtootokenhistory.HistoryTimeLTE(time)).
		Order(ohauthtootokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (oatth *OhAuthTooTokenHistory) Restore(ctx context.Context) (*OhAuthTooToken, error) {
	client := NewOhAuthTooTokenClient(oatth.config)
	return client.
		UpdateOneID(oatth.Ref).
		SetClientID(oatth.ClientID).
		SetScopes(oatth.Scopes).
		SetNonce(oatth.Nonce).
		SetClaimsUserID(oatth.ClaimsUserID).
		SetClaimsUsername(oatth.ClaimsUsername).
		SetClaimsEmail(oatth.ClaimsEmail).
		SetClaimsEmailVerified(oatth.ClaimsEmailVerified).
		SetClaimsGroups(oatth.ClaimsGroups).
		SetClaimsPreferredUsername(oatth.ClaimsPreferredUsername).
		SetConnectorID(oatth.ConnectorID).
		SetConnectorData(oatth.ConnectorData).
		SetLastUsed(oatth.LastUsed).
		Save(ctx)
}

func (o *Organization) History() *OrganizationHistoryQuery {
	historyClient := NewOrganizationHistoryClient(o.config)
	return historyClient.Query().Where(organizationhistory.Ref(o.ID))
}

func (oh *OrganizationHistory) Next(ctx context.Context) (*OrganizationHistory, error) {
	client := NewOrganizationHistoryClient(oh.config)
	return client.Query().
		Where(
			organizationhistory.Ref(oh.Ref),
			organizationhistory.HistoryTimeGT(oh.HistoryTime),
		).
		Order(organizationhistory.ByHistoryTime()).
		First(ctx)
}

func (oh *OrganizationHistory) Prev(ctx context.Context) (*OrganizationHistory, error) {
	client := NewOrganizationHistoryClient(oh.config)
	return client.Query().
		Where(
			organizationhistory.Ref(oh.Ref),
			organizationhistory.HistoryTimeLT(oh.HistoryTime),
		).
		Order(organizationhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ohq *OrganizationHistoryQuery) Earliest(ctx context.Context) (*OrganizationHistory, error) {
	return ohq.
		Order(organizationhistory.ByHistoryTime()).
		First(ctx)
}

func (ohq *OrganizationHistoryQuery) Latest(ctx context.Context) (*OrganizationHistory, error) {
	return ohq.
		Order(organizationhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ohq *OrganizationHistoryQuery) AsOf(ctx context.Context, time time.Time) (*OrganizationHistory, error) {
	return ohq.
		Where(organizationhistory.HistoryTimeLTE(time)).
		Order(organizationhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (oh *OrganizationHistory) Restore(ctx context.Context) (*Organization, error) {
	client := NewOrganizationClient(oh.config)
	return client.
		UpdateOneID(oh.Ref).
		SetUpdatedAt(oh.UpdatedAt).
		SetUpdatedBy(oh.UpdatedBy).
		SetDeletedAt(oh.DeletedAt).
		SetDeletedBy(oh.DeletedBy).
		SetName(oh.Name).
		SetDisplayName(oh.DisplayName).
		SetDescription(oh.Description).
		Save(ctx)
}

func (os *OrganizationSetting) History() *OrganizationSettingHistoryQuery {
	historyClient := NewOrganizationSettingHistoryClient(os.config)
	return historyClient.Query().Where(organizationsettinghistory.Ref(os.ID))
}

func (osh *OrganizationSettingHistory) Next(ctx context.Context) (*OrganizationSettingHistory, error) {
	client := NewOrganizationSettingHistoryClient(osh.config)
	return client.Query().
		Where(
			organizationsettinghistory.Ref(osh.Ref),
			organizationsettinghistory.HistoryTimeGT(osh.HistoryTime),
		).
		Order(organizationsettinghistory.ByHistoryTime()).
		First(ctx)
}

func (osh *OrganizationSettingHistory) Prev(ctx context.Context) (*OrganizationSettingHistory, error) {
	client := NewOrganizationSettingHistoryClient(osh.config)
	return client.Query().
		Where(
			organizationsettinghistory.Ref(osh.Ref),
			organizationsettinghistory.HistoryTimeLT(osh.HistoryTime),
		).
		Order(organizationsettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (oshq *OrganizationSettingHistoryQuery) Earliest(ctx context.Context) (*OrganizationSettingHistory, error) {
	return oshq.
		Order(organizationsettinghistory.ByHistoryTime()).
		First(ctx)
}

func (oshq *OrganizationSettingHistoryQuery) Latest(ctx context.Context) (*OrganizationSettingHistory, error) {
	return oshq.
		Order(organizationsettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (oshq *OrganizationSettingHistoryQuery) AsOf(ctx context.Context, time time.Time) (*OrganizationSettingHistory, error) {
	return oshq.
		Where(organizationsettinghistory.HistoryTimeLTE(time)).
		Order(organizationsettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (osh *OrganizationSettingHistory) Restore(ctx context.Context) (*OrganizationSetting, error) {
	client := NewOrganizationSettingClient(osh.config)
	return client.
		UpdateOneID(osh.Ref).
		SetUpdatedAt(osh.UpdatedAt).
		SetUpdatedBy(osh.UpdatedBy).
		SetDeletedAt(osh.DeletedAt).
		SetDeletedBy(osh.DeletedBy).
		SetDomains(osh.Domains).
		SetSSOCert(osh.SSOCert).
		SetSSOEntrypoint(osh.SSOEntrypoint).
		SetSSOIssuer(osh.SSOIssuer).
		SetBillingContact(osh.BillingContact).
		SetBillingEmail(osh.BillingEmail).
		SetBillingPhone(osh.BillingPhone).
		SetBillingAddress(osh.BillingAddress).
		SetTaxIdentifier(osh.TaxIdentifier).
		SetTags(osh.Tags).
		Save(ctx)
}

func (prt *PasswordResetToken) History() *PasswordResetTokenHistoryQuery {
	historyClient := NewPasswordResetTokenHistoryClient(prt.config)
	return historyClient.Query().Where(passwordresettokenhistory.Ref(prt.ID))
}

func (prth *PasswordResetTokenHistory) Next(ctx context.Context) (*PasswordResetTokenHistory, error) {
	client := NewPasswordResetTokenHistoryClient(prth.config)
	return client.Query().
		Where(
			passwordresettokenhistory.Ref(prth.Ref),
			passwordresettokenhistory.HistoryTimeGT(prth.HistoryTime),
		).
		Order(passwordresettokenhistory.ByHistoryTime()).
		First(ctx)
}

func (prth *PasswordResetTokenHistory) Prev(ctx context.Context) (*PasswordResetTokenHistory, error) {
	client := NewPasswordResetTokenHistoryClient(prth.config)
	return client.Query().
		Where(
			passwordresettokenhistory.Ref(prth.Ref),
			passwordresettokenhistory.HistoryTimeLT(prth.HistoryTime),
		).
		Order(passwordresettokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (prthq *PasswordResetTokenHistoryQuery) Earliest(ctx context.Context) (*PasswordResetTokenHistory, error) {
	return prthq.
		Order(passwordresettokenhistory.ByHistoryTime()).
		First(ctx)
}

func (prthq *PasswordResetTokenHistoryQuery) Latest(ctx context.Context) (*PasswordResetTokenHistory, error) {
	return prthq.
		Order(passwordresettokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (prthq *PasswordResetTokenHistoryQuery) AsOf(ctx context.Context, time time.Time) (*PasswordResetTokenHistory, error) {
	return prthq.
		Where(passwordresettokenhistory.HistoryTimeLTE(time)).
		Order(passwordresettokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (prth *PasswordResetTokenHistory) Restore(ctx context.Context) (*PasswordResetToken, error) {
	client := NewPasswordResetTokenClient(prth.config)
	return client.
		UpdateOneID(prth.Ref).
		SetUpdatedAt(prth.UpdatedAt).
		SetUpdatedBy(prth.UpdatedBy).
		SetDeletedAt(prth.DeletedAt).
		SetDeletedBy(prth.DeletedBy).
		SetToken(prth.Token).
		SetNillableTTL(prth.TTL).
		SetEmail(prth.Email).
		SetNillableSecret(prth.Secret).
		Save(ctx)
}

func (pat *PersonalAccessToken) History() *PersonalAccessTokenHistoryQuery {
	historyClient := NewPersonalAccessTokenHistoryClient(pat.config)
	return historyClient.Query().Where(personalaccesstokenhistory.Ref(pat.ID))
}

func (path *PersonalAccessTokenHistory) Next(ctx context.Context) (*PersonalAccessTokenHistory, error) {
	client := NewPersonalAccessTokenHistoryClient(path.config)
	return client.Query().
		Where(
			personalaccesstokenhistory.Ref(path.Ref),
			personalaccesstokenhistory.HistoryTimeGT(path.HistoryTime),
		).
		Order(personalaccesstokenhistory.ByHistoryTime()).
		First(ctx)
}

func (path *PersonalAccessTokenHistory) Prev(ctx context.Context) (*PersonalAccessTokenHistory, error) {
	client := NewPersonalAccessTokenHistoryClient(path.config)
	return client.Query().
		Where(
			personalaccesstokenhistory.Ref(path.Ref),
			personalaccesstokenhistory.HistoryTimeLT(path.HistoryTime),
		).
		Order(personalaccesstokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (pathq *PersonalAccessTokenHistoryQuery) Earliest(ctx context.Context) (*PersonalAccessTokenHistory, error) {
	return pathq.
		Order(personalaccesstokenhistory.ByHistoryTime()).
		First(ctx)
}

func (pathq *PersonalAccessTokenHistoryQuery) Latest(ctx context.Context) (*PersonalAccessTokenHistory, error) {
	return pathq.
		Order(personalaccesstokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (pathq *PersonalAccessTokenHistoryQuery) AsOf(ctx context.Context, time time.Time) (*PersonalAccessTokenHistory, error) {
	return pathq.
		Where(personalaccesstokenhistory.HistoryTimeLTE(time)).
		Order(personalaccesstokenhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (path *PersonalAccessTokenHistory) Restore(ctx context.Context) (*PersonalAccessToken, error) {
	client := NewPersonalAccessTokenClient(path.config)
	return client.
		UpdateOneID(path.Ref).
		SetUpdatedAt(path.UpdatedAt).
		SetUpdatedBy(path.UpdatedBy).
		SetDeletedAt(path.DeletedAt).
		SetDeletedBy(path.DeletedBy).
		SetName(path.Name).
		SetAbilities(path.Abilities).
		SetNillableExpiresAt(path.ExpiresAt).
		SetDescription(path.Description).
		SetNillableLastUsedAt(path.LastUsedAt).
		Save(ctx)
}

func (s *Session) History() *SessionHistoryQuery {
	historyClient := NewSessionHistoryClient(s.config)
	return historyClient.Query().Where(sessionhistory.Ref(s.ID))
}

func (sh *SessionHistory) Next(ctx context.Context) (*SessionHistory, error) {
	client := NewSessionHistoryClient(sh.config)
	return client.Query().
		Where(
			sessionhistory.Ref(sh.Ref),
			sessionhistory.HistoryTimeGT(sh.HistoryTime),
		).
		Order(sessionhistory.ByHistoryTime()).
		First(ctx)
}

func (sh *SessionHistory) Prev(ctx context.Context) (*SessionHistory, error) {
	client := NewSessionHistoryClient(sh.config)
	return client.Query().
		Where(
			sessionhistory.Ref(sh.Ref),
			sessionhistory.HistoryTimeLT(sh.HistoryTime),
		).
		Order(sessionhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (shq *SessionHistoryQuery) Earliest(ctx context.Context) (*SessionHistory, error) {
	return shq.
		Order(sessionhistory.ByHistoryTime()).
		First(ctx)
}

func (shq *SessionHistoryQuery) Latest(ctx context.Context) (*SessionHistory, error) {
	return shq.
		Order(sessionhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (shq *SessionHistoryQuery) AsOf(ctx context.Context, time time.Time) (*SessionHistory, error) {
	return shq.
		Where(sessionhistory.HistoryTimeLTE(time)).
		Order(sessionhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (sh *SessionHistory) Restore(ctx context.Context) (*Session, error) {
	client := NewSessionClient(sh.config)
	return client.
		UpdateOneID(sh.Ref).
		SetUpdatedAt(sh.UpdatedAt).
		SetUpdatedBy(sh.UpdatedBy).
		SetIssuedAt(sh.IssuedAt).
		SetExpiresAt(sh.ExpiresAt).
		SetOrganizationID(sh.OrganizationID).
		SetUserID(sh.UserID).
		Save(ctx)
}

func (u *User) History() *UserHistoryQuery {
	historyClient := NewUserHistoryClient(u.config)
	return historyClient.Query().Where(userhistory.Ref(u.ID))
}

func (uh *UserHistory) Next(ctx context.Context) (*UserHistory, error) {
	client := NewUserHistoryClient(uh.config)
	return client.Query().
		Where(
			userhistory.Ref(uh.Ref),
			userhistory.HistoryTimeGT(uh.HistoryTime),
		).
		Order(userhistory.ByHistoryTime()).
		First(ctx)
}

func (uh *UserHistory) Prev(ctx context.Context) (*UserHistory, error) {
	client := NewUserHistoryClient(uh.config)
	return client.Query().
		Where(
			userhistory.Ref(uh.Ref),
			userhistory.HistoryTimeLT(uh.HistoryTime),
		).
		Order(userhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (uhq *UserHistoryQuery) Earliest(ctx context.Context) (*UserHistory, error) {
	return uhq.
		Order(userhistory.ByHistoryTime()).
		First(ctx)
}

func (uhq *UserHistoryQuery) Latest(ctx context.Context) (*UserHistory, error) {
	return uhq.
		Order(userhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (uhq *UserHistoryQuery) AsOf(ctx context.Context, time time.Time) (*UserHistory, error) {
	return uhq.
		Where(userhistory.HistoryTimeLTE(time)).
		Order(userhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (uh *UserHistory) Restore(ctx context.Context) (*User, error) {
	client := NewUserClient(uh.config)
	return client.
		UpdateOneID(uh.Ref).
		SetUpdatedAt(uh.UpdatedAt).
		SetUpdatedBy(uh.UpdatedBy).
		SetDeletedAt(uh.DeletedAt).
		SetDeletedBy(uh.DeletedBy).
		SetEmail(uh.Email).
		SetFirstName(uh.FirstName).
		SetLastName(uh.LastName).
		SetDisplayName(uh.DisplayName).
		SetNillableAvatarRemoteURL(uh.AvatarRemoteURL).
		SetNillableAvatarLocalFile(uh.AvatarLocalFile).
		SetNillableAvatarUpdatedAt(uh.AvatarUpdatedAt).
		SetNillableLastSeen(uh.LastSeen).
		SetNillablePassword(uh.Password).
		SetSub(uh.Sub).
		SetOauth(uh.Oauth).
		Save(ctx)
}

func (us *UserSetting) History() *UserSettingHistoryQuery {
	historyClient := NewUserSettingHistoryClient(us.config)
	return historyClient.Query().Where(usersettinghistory.Ref(us.ID))
}

func (ush *UserSettingHistory) Next(ctx context.Context) (*UserSettingHistory, error) {
	client := NewUserSettingHistoryClient(ush.config)
	return client.Query().
		Where(
			usersettinghistory.Ref(ush.Ref),
			usersettinghistory.HistoryTimeGT(ush.HistoryTime),
		).
		Order(usersettinghistory.ByHistoryTime()).
		First(ctx)
}

func (ush *UserSettingHistory) Prev(ctx context.Context) (*UserSettingHistory, error) {
	client := NewUserSettingHistoryClient(ush.config)
	return client.Query().
		Where(
			usersettinghistory.Ref(ush.Ref),
			usersettinghistory.HistoryTimeLT(ush.HistoryTime),
		).
		Order(usersettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ushq *UserSettingHistoryQuery) Earliest(ctx context.Context) (*UserSettingHistory, error) {
	return ushq.
		Order(usersettinghistory.ByHistoryTime()).
		First(ctx)
}

func (ushq *UserSettingHistoryQuery) Latest(ctx context.Context) (*UserSettingHistory, error) {
	return ushq.
		Order(usersettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ushq *UserSettingHistoryQuery) AsOf(ctx context.Context, time time.Time) (*UserSettingHistory, error) {
	return ushq.
		Where(usersettinghistory.HistoryTimeLTE(time)).
		Order(usersettinghistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ush *UserSettingHistory) Restore(ctx context.Context) (*UserSetting, error) {
	client := NewUserSettingClient(ush.config)
	return client.
		UpdateOneID(ush.Ref).
		SetUpdatedAt(ush.UpdatedAt).
		SetUpdatedBy(ush.UpdatedBy).
		SetDeletedAt(ush.DeletedAt).
		SetDeletedBy(ush.DeletedBy).
		SetLocked(ush.Locked).
		SetNillableSilencedAt(ush.SilencedAt).
		SetNillableSuspendedAt(ush.SuspendedAt).
		SetNillableRecoveryCode(ush.RecoveryCode).
		SetStatus(ush.Status).
		SetRole(ush.Role).
		SetPermissions(ush.Permissions).
		SetEmailConfirmed(ush.EmailConfirmed).
		SetTags(ush.Tags).
		Save(ctx)
}
