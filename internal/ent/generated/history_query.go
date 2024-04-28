// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/documentdatahistory"
	"github.com/datumforge/datum/internal/ent/generated/entitlementhistory"
	"github.com/datumforge/datum/internal/ent/generated/grouphistory"
	"github.com/datumforge/datum/internal/ent/generated/groupmembershiphistory"
	"github.com/datumforge/datum/internal/ent/generated/groupsettinghistory"
	"github.com/datumforge/datum/internal/ent/generated/hushhistory"
	"github.com/datumforge/datum/internal/ent/generated/integrationhistory"
	"github.com/datumforge/datum/internal/ent/generated/oauthproviderhistory"
	"github.com/datumforge/datum/internal/ent/generated/organizationhistory"
	"github.com/datumforge/datum/internal/ent/generated/organizationsettinghistory"
	"github.com/datumforge/datum/internal/ent/generated/orgmembershiphistory"
	"github.com/datumforge/datum/internal/ent/generated/templatehistory"
	"github.com/datumforge/datum/internal/ent/generated/userhistory"
	"github.com/datumforge/datum/internal/ent/generated/usersettinghistory"
)

func (dd *DocumentData) History() *DocumentDataHistoryQuery {
	historyClient := NewDocumentDataHistoryClient(dd.config)
	return historyClient.Query().Where(documentdatahistory.Ref(dd.ID))
}

func (ddh *DocumentDataHistory) Next(ctx context.Context) (*DocumentDataHistory, error) {
	client := NewDocumentDataHistoryClient(ddh.config)
	return client.Query().
		Where(
			documentdatahistory.Ref(ddh.Ref),
			documentdatahistory.HistoryTimeGT(ddh.HistoryTime),
		).
		Order(documentdatahistory.ByHistoryTime()).
		First(ctx)
}

func (ddh *DocumentDataHistory) Prev(ctx context.Context) (*DocumentDataHistory, error) {
	client := NewDocumentDataHistoryClient(ddh.config)
	return client.Query().
		Where(
			documentdatahistory.Ref(ddh.Ref),
			documentdatahistory.HistoryTimeLT(ddh.HistoryTime),
		).
		Order(documentdatahistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ddhq *DocumentDataHistoryQuery) Earliest(ctx context.Context) (*DocumentDataHistory, error) {
	return ddhq.
		Order(documentdatahistory.ByHistoryTime()).
		First(ctx)
}

func (ddhq *DocumentDataHistoryQuery) Latest(ctx context.Context) (*DocumentDataHistory, error) {
	return ddhq.
		Order(documentdatahistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ddhq *DocumentDataHistoryQuery) AsOf(ctx context.Context, time time.Time) (*DocumentDataHistory, error) {
	return ddhq.
		Where(documentdatahistory.HistoryTimeLTE(time)).
		Order(documentdatahistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
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

func (gm *GroupMembership) History() *GroupMembershipHistoryQuery {
	historyClient := NewGroupMembershipHistoryClient(gm.config)
	return historyClient.Query().Where(groupmembershiphistory.Ref(gm.ID))
}

func (gmh *GroupMembershipHistory) Next(ctx context.Context) (*GroupMembershipHistory, error) {
	client := NewGroupMembershipHistoryClient(gmh.config)
	return client.Query().
		Where(
			groupmembershiphistory.Ref(gmh.Ref),
			groupmembershiphistory.HistoryTimeGT(gmh.HistoryTime),
		).
		Order(groupmembershiphistory.ByHistoryTime()).
		First(ctx)
}

func (gmh *GroupMembershipHistory) Prev(ctx context.Context) (*GroupMembershipHistory, error) {
	client := NewGroupMembershipHistoryClient(gmh.config)
	return client.Query().
		Where(
			groupmembershiphistory.Ref(gmh.Ref),
			groupmembershiphistory.HistoryTimeLT(gmh.HistoryTime),
		).
		Order(groupmembershiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (gmhq *GroupMembershipHistoryQuery) Earliest(ctx context.Context) (*GroupMembershipHistory, error) {
	return gmhq.
		Order(groupmembershiphistory.ByHistoryTime()).
		First(ctx)
}

func (gmhq *GroupMembershipHistoryQuery) Latest(ctx context.Context) (*GroupMembershipHistory, error) {
	return gmhq.
		Order(groupmembershiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (gmhq *GroupMembershipHistoryQuery) AsOf(ctx context.Context, time time.Time) (*GroupMembershipHistory, error) {
	return gmhq.
		Where(groupmembershiphistory.HistoryTimeLTE(time)).
		Order(groupmembershiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
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

func (h *Hush) History() *HushHistoryQuery {
	historyClient := NewHushHistoryClient(h.config)
	return historyClient.Query().Where(hushhistory.Ref(h.ID))
}

func (hh *HushHistory) Next(ctx context.Context) (*HushHistory, error) {
	client := NewHushHistoryClient(hh.config)
	return client.Query().
		Where(
			hushhistory.Ref(hh.Ref),
			hushhistory.HistoryTimeGT(hh.HistoryTime),
		).
		Order(hushhistory.ByHistoryTime()).
		First(ctx)
}

func (hh *HushHistory) Prev(ctx context.Context) (*HushHistory, error) {
	client := NewHushHistoryClient(hh.config)
	return client.Query().
		Where(
			hushhistory.Ref(hh.Ref),
			hushhistory.HistoryTimeLT(hh.HistoryTime),
		).
		Order(hushhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (hhq *HushHistoryQuery) Earliest(ctx context.Context) (*HushHistory, error) {
	return hhq.
		Order(hushhistory.ByHistoryTime()).
		First(ctx)
}

func (hhq *HushHistoryQuery) Latest(ctx context.Context) (*HushHistory, error) {
	return hhq.
		Order(hushhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (hhq *HushHistoryQuery) AsOf(ctx context.Context, time time.Time) (*HushHistory, error) {
	return hhq.
		Where(hushhistory.HistoryTimeLTE(time)).
		Order(hushhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
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

func (om *OrgMembership) History() *OrgMembershipHistoryQuery {
	historyClient := NewOrgMembershipHistoryClient(om.config)
	return historyClient.Query().Where(orgmembershiphistory.Ref(om.ID))
}

func (omh *OrgMembershipHistory) Next(ctx context.Context) (*OrgMembershipHistory, error) {
	client := NewOrgMembershipHistoryClient(omh.config)
	return client.Query().
		Where(
			orgmembershiphistory.Ref(omh.Ref),
			orgmembershiphistory.HistoryTimeGT(omh.HistoryTime),
		).
		Order(orgmembershiphistory.ByHistoryTime()).
		First(ctx)
}

func (omh *OrgMembershipHistory) Prev(ctx context.Context) (*OrgMembershipHistory, error) {
	client := NewOrgMembershipHistoryClient(omh.config)
	return client.Query().
		Where(
			orgmembershiphistory.Ref(omh.Ref),
			orgmembershiphistory.HistoryTimeLT(omh.HistoryTime),
		).
		Order(orgmembershiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (omhq *OrgMembershipHistoryQuery) Earliest(ctx context.Context) (*OrgMembershipHistory, error) {
	return omhq.
		Order(orgmembershiphistory.ByHistoryTime()).
		First(ctx)
}

func (omhq *OrgMembershipHistoryQuery) Latest(ctx context.Context) (*OrgMembershipHistory, error) {
	return omhq.
		Order(orgmembershiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (omhq *OrgMembershipHistoryQuery) AsOf(ctx context.Context, time time.Time) (*OrgMembershipHistory, error) {
	return omhq.
		Where(orgmembershiphistory.HistoryTimeLTE(time)).
		Order(orgmembershiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
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

func (t *Template) History() *TemplateHistoryQuery {
	historyClient := NewTemplateHistoryClient(t.config)
	return historyClient.Query().Where(templatehistory.Ref(t.ID))
}

func (th *TemplateHistory) Next(ctx context.Context) (*TemplateHistory, error) {
	client := NewTemplateHistoryClient(th.config)
	return client.Query().
		Where(
			templatehistory.Ref(th.Ref),
			templatehistory.HistoryTimeGT(th.HistoryTime),
		).
		Order(templatehistory.ByHistoryTime()).
		First(ctx)
}

func (th *TemplateHistory) Prev(ctx context.Context) (*TemplateHistory, error) {
	client := NewTemplateHistoryClient(th.config)
	return client.Query().
		Where(
			templatehistory.Ref(th.Ref),
			templatehistory.HistoryTimeLT(th.HistoryTime),
		).
		Order(templatehistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (thq *TemplateHistoryQuery) Earliest(ctx context.Context) (*TemplateHistory, error) {
	return thq.
		Order(templatehistory.ByHistoryTime()).
		First(ctx)
}

func (thq *TemplateHistoryQuery) Latest(ctx context.Context) (*TemplateHistory, error) {
	return thq.
		Order(templatehistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (thq *TemplateHistoryQuery) AsOf(ctx context.Context, time time.Time) (*TemplateHistory, error) {
	return thq.
		Where(templatehistory.HistoryTimeLTE(time)).
		Order(templatehistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
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
