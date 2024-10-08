// Code generated by ent, DO NOT EDIT.

package ohauthtootoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldID, id))
}

// MappingID applies equality check predicate on the "mapping_id" field. It's identical to MappingIDEQ.
func MappingID(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldMappingID, v))
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClientID, v))
}

// Nonce applies equality check predicate on the "nonce" field. It's identical to NonceEQ.
func Nonce(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldNonce, v))
}

// ClaimsUserID applies equality check predicate on the "claims_user_id" field. It's identical to ClaimsUserIDEQ.
func ClaimsUserID(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsUserID, v))
}

// ClaimsUsername applies equality check predicate on the "claims_username" field. It's identical to ClaimsUsernameEQ.
func ClaimsUsername(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsUsername, v))
}

// ClaimsEmail applies equality check predicate on the "claims_email" field. It's identical to ClaimsEmailEQ.
func ClaimsEmail(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsEmail, v))
}

// ClaimsEmailVerified applies equality check predicate on the "claims_email_verified" field. It's identical to ClaimsEmailVerifiedEQ.
func ClaimsEmailVerified(v bool) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsEmailVerified, v))
}

// ClaimsPreferredUsername applies equality check predicate on the "claims_preferred_username" field. It's identical to ClaimsPreferredUsernameEQ.
func ClaimsPreferredUsername(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsPreferredUsername, v))
}

// ConnectorID applies equality check predicate on the "connector_id" field. It's identical to ConnectorIDEQ.
func ConnectorID(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldConnectorID, v))
}

// LastUsed applies equality check predicate on the "last_used" field. It's identical to LastUsedEQ.
func LastUsed(v time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldLastUsed, v))
}

// MappingIDEQ applies the EQ predicate on the "mapping_id" field.
func MappingIDEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldMappingID, v))
}

// MappingIDNEQ applies the NEQ predicate on the "mapping_id" field.
func MappingIDNEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldMappingID, v))
}

// MappingIDIn applies the In predicate on the "mapping_id" field.
func MappingIDIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldMappingID, vs...))
}

// MappingIDNotIn applies the NotIn predicate on the "mapping_id" field.
func MappingIDNotIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldMappingID, vs...))
}

// MappingIDGT applies the GT predicate on the "mapping_id" field.
func MappingIDGT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldMappingID, v))
}

// MappingIDGTE applies the GTE predicate on the "mapping_id" field.
func MappingIDGTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldMappingID, v))
}

// MappingIDLT applies the LT predicate on the "mapping_id" field.
func MappingIDLT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldMappingID, v))
}

// MappingIDLTE applies the LTE predicate on the "mapping_id" field.
func MappingIDLTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldMappingID, v))
}

// MappingIDContains applies the Contains predicate on the "mapping_id" field.
func MappingIDContains(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContains(FieldMappingID, v))
}

// MappingIDHasPrefix applies the HasPrefix predicate on the "mapping_id" field.
func MappingIDHasPrefix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasPrefix(FieldMappingID, v))
}

// MappingIDHasSuffix applies the HasSuffix predicate on the "mapping_id" field.
func MappingIDHasSuffix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasSuffix(FieldMappingID, v))
}

// MappingIDEqualFold applies the EqualFold predicate on the "mapping_id" field.
func MappingIDEqualFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldMappingID, v))
}

// MappingIDContainsFold applies the ContainsFold predicate on the "mapping_id" field.
func MappingIDContainsFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldMappingID, v))
}

// TagsIsNil applies the IsNil predicate on the "tags" field.
func TagsIsNil() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIsNull(FieldTags))
}

// TagsNotNil applies the NotNil predicate on the "tags" field.
func TagsNotNil() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotNull(FieldTags))
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClientID, v))
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldClientID, v))
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldClientID, vs...))
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldClientID, vs...))
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldClientID, v))
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldClientID, v))
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldClientID, v))
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldClientID, v))
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContains(FieldClientID, v))
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasPrefix(FieldClientID, v))
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasSuffix(FieldClientID, v))
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldClientID, v))
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldClientID, v))
}

// ScopesIsNil applies the IsNil predicate on the "scopes" field.
func ScopesIsNil() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIsNull(FieldScopes))
}

// ScopesNotNil applies the NotNil predicate on the "scopes" field.
func ScopesNotNil() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotNull(FieldScopes))
}

// NonceEQ applies the EQ predicate on the "nonce" field.
func NonceEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldNonce, v))
}

// NonceNEQ applies the NEQ predicate on the "nonce" field.
func NonceNEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldNonce, v))
}

// NonceIn applies the In predicate on the "nonce" field.
func NonceIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldNonce, vs...))
}

// NonceNotIn applies the NotIn predicate on the "nonce" field.
func NonceNotIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldNonce, vs...))
}

// NonceGT applies the GT predicate on the "nonce" field.
func NonceGT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldNonce, v))
}

// NonceGTE applies the GTE predicate on the "nonce" field.
func NonceGTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldNonce, v))
}

// NonceLT applies the LT predicate on the "nonce" field.
func NonceLT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldNonce, v))
}

// NonceLTE applies the LTE predicate on the "nonce" field.
func NonceLTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldNonce, v))
}

// NonceContains applies the Contains predicate on the "nonce" field.
func NonceContains(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContains(FieldNonce, v))
}

// NonceHasPrefix applies the HasPrefix predicate on the "nonce" field.
func NonceHasPrefix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasPrefix(FieldNonce, v))
}

// NonceHasSuffix applies the HasSuffix predicate on the "nonce" field.
func NonceHasSuffix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasSuffix(FieldNonce, v))
}

// NonceEqualFold applies the EqualFold predicate on the "nonce" field.
func NonceEqualFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldNonce, v))
}

// NonceContainsFold applies the ContainsFold predicate on the "nonce" field.
func NonceContainsFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldNonce, v))
}

// ClaimsUserIDEQ applies the EQ predicate on the "claims_user_id" field.
func ClaimsUserIDEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsUserID, v))
}

// ClaimsUserIDNEQ applies the NEQ predicate on the "claims_user_id" field.
func ClaimsUserIDNEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldClaimsUserID, v))
}

// ClaimsUserIDIn applies the In predicate on the "claims_user_id" field.
func ClaimsUserIDIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldClaimsUserID, vs...))
}

// ClaimsUserIDNotIn applies the NotIn predicate on the "claims_user_id" field.
func ClaimsUserIDNotIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldClaimsUserID, vs...))
}

// ClaimsUserIDGT applies the GT predicate on the "claims_user_id" field.
func ClaimsUserIDGT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldClaimsUserID, v))
}

// ClaimsUserIDGTE applies the GTE predicate on the "claims_user_id" field.
func ClaimsUserIDGTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldClaimsUserID, v))
}

// ClaimsUserIDLT applies the LT predicate on the "claims_user_id" field.
func ClaimsUserIDLT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldClaimsUserID, v))
}

// ClaimsUserIDLTE applies the LTE predicate on the "claims_user_id" field.
func ClaimsUserIDLTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldClaimsUserID, v))
}

// ClaimsUserIDContains applies the Contains predicate on the "claims_user_id" field.
func ClaimsUserIDContains(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContains(FieldClaimsUserID, v))
}

// ClaimsUserIDHasPrefix applies the HasPrefix predicate on the "claims_user_id" field.
func ClaimsUserIDHasPrefix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasPrefix(FieldClaimsUserID, v))
}

// ClaimsUserIDHasSuffix applies the HasSuffix predicate on the "claims_user_id" field.
func ClaimsUserIDHasSuffix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasSuffix(FieldClaimsUserID, v))
}

// ClaimsUserIDEqualFold applies the EqualFold predicate on the "claims_user_id" field.
func ClaimsUserIDEqualFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldClaimsUserID, v))
}

// ClaimsUserIDContainsFold applies the ContainsFold predicate on the "claims_user_id" field.
func ClaimsUserIDContainsFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldClaimsUserID, v))
}

// ClaimsUsernameEQ applies the EQ predicate on the "claims_username" field.
func ClaimsUsernameEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsUsername, v))
}

// ClaimsUsernameNEQ applies the NEQ predicate on the "claims_username" field.
func ClaimsUsernameNEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldClaimsUsername, v))
}

// ClaimsUsernameIn applies the In predicate on the "claims_username" field.
func ClaimsUsernameIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldClaimsUsername, vs...))
}

// ClaimsUsernameNotIn applies the NotIn predicate on the "claims_username" field.
func ClaimsUsernameNotIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldClaimsUsername, vs...))
}

// ClaimsUsernameGT applies the GT predicate on the "claims_username" field.
func ClaimsUsernameGT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldClaimsUsername, v))
}

// ClaimsUsernameGTE applies the GTE predicate on the "claims_username" field.
func ClaimsUsernameGTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldClaimsUsername, v))
}

// ClaimsUsernameLT applies the LT predicate on the "claims_username" field.
func ClaimsUsernameLT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldClaimsUsername, v))
}

// ClaimsUsernameLTE applies the LTE predicate on the "claims_username" field.
func ClaimsUsernameLTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldClaimsUsername, v))
}

// ClaimsUsernameContains applies the Contains predicate on the "claims_username" field.
func ClaimsUsernameContains(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContains(FieldClaimsUsername, v))
}

// ClaimsUsernameHasPrefix applies the HasPrefix predicate on the "claims_username" field.
func ClaimsUsernameHasPrefix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasPrefix(FieldClaimsUsername, v))
}

// ClaimsUsernameHasSuffix applies the HasSuffix predicate on the "claims_username" field.
func ClaimsUsernameHasSuffix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasSuffix(FieldClaimsUsername, v))
}

// ClaimsUsernameEqualFold applies the EqualFold predicate on the "claims_username" field.
func ClaimsUsernameEqualFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldClaimsUsername, v))
}

// ClaimsUsernameContainsFold applies the ContainsFold predicate on the "claims_username" field.
func ClaimsUsernameContainsFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldClaimsUsername, v))
}

// ClaimsEmailEQ applies the EQ predicate on the "claims_email" field.
func ClaimsEmailEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsEmail, v))
}

// ClaimsEmailNEQ applies the NEQ predicate on the "claims_email" field.
func ClaimsEmailNEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldClaimsEmail, v))
}

// ClaimsEmailIn applies the In predicate on the "claims_email" field.
func ClaimsEmailIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldClaimsEmail, vs...))
}

// ClaimsEmailNotIn applies the NotIn predicate on the "claims_email" field.
func ClaimsEmailNotIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldClaimsEmail, vs...))
}

// ClaimsEmailGT applies the GT predicate on the "claims_email" field.
func ClaimsEmailGT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldClaimsEmail, v))
}

// ClaimsEmailGTE applies the GTE predicate on the "claims_email" field.
func ClaimsEmailGTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldClaimsEmail, v))
}

// ClaimsEmailLT applies the LT predicate on the "claims_email" field.
func ClaimsEmailLT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldClaimsEmail, v))
}

// ClaimsEmailLTE applies the LTE predicate on the "claims_email" field.
func ClaimsEmailLTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldClaimsEmail, v))
}

// ClaimsEmailContains applies the Contains predicate on the "claims_email" field.
func ClaimsEmailContains(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContains(FieldClaimsEmail, v))
}

// ClaimsEmailHasPrefix applies the HasPrefix predicate on the "claims_email" field.
func ClaimsEmailHasPrefix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasPrefix(FieldClaimsEmail, v))
}

// ClaimsEmailHasSuffix applies the HasSuffix predicate on the "claims_email" field.
func ClaimsEmailHasSuffix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasSuffix(FieldClaimsEmail, v))
}

// ClaimsEmailEqualFold applies the EqualFold predicate on the "claims_email" field.
func ClaimsEmailEqualFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldClaimsEmail, v))
}

// ClaimsEmailContainsFold applies the ContainsFold predicate on the "claims_email" field.
func ClaimsEmailContainsFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldClaimsEmail, v))
}

// ClaimsEmailVerifiedEQ applies the EQ predicate on the "claims_email_verified" field.
func ClaimsEmailVerifiedEQ(v bool) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsEmailVerified, v))
}

// ClaimsEmailVerifiedNEQ applies the NEQ predicate on the "claims_email_verified" field.
func ClaimsEmailVerifiedNEQ(v bool) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldClaimsEmailVerified, v))
}

// ClaimsGroupsIsNil applies the IsNil predicate on the "claims_groups" field.
func ClaimsGroupsIsNil() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIsNull(FieldClaimsGroups))
}

// ClaimsGroupsNotNil applies the NotNil predicate on the "claims_groups" field.
func ClaimsGroupsNotNil() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotNull(FieldClaimsGroups))
}

// ClaimsPreferredUsernameEQ applies the EQ predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameNEQ applies the NEQ predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameNEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameIn applies the In predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldClaimsPreferredUsername, vs...))
}

// ClaimsPreferredUsernameNotIn applies the NotIn predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameNotIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldClaimsPreferredUsername, vs...))
}

// ClaimsPreferredUsernameGT applies the GT predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameGT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameGTE applies the GTE predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameGTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameLT applies the LT predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameLT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameLTE applies the LTE predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameLTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameContains applies the Contains predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameContains(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContains(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameHasPrefix applies the HasPrefix predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameHasPrefix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasPrefix(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameHasSuffix applies the HasSuffix predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameHasSuffix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasSuffix(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameEqualFold applies the EqualFold predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameEqualFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameContainsFold applies the ContainsFold predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameContainsFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldClaimsPreferredUsername, v))
}

// ConnectorIDEQ applies the EQ predicate on the "connector_id" field.
func ConnectorIDEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldConnectorID, v))
}

// ConnectorIDNEQ applies the NEQ predicate on the "connector_id" field.
func ConnectorIDNEQ(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldConnectorID, v))
}

// ConnectorIDIn applies the In predicate on the "connector_id" field.
func ConnectorIDIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldConnectorID, vs...))
}

// ConnectorIDNotIn applies the NotIn predicate on the "connector_id" field.
func ConnectorIDNotIn(vs ...string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldConnectorID, vs...))
}

// ConnectorIDGT applies the GT predicate on the "connector_id" field.
func ConnectorIDGT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldConnectorID, v))
}

// ConnectorIDGTE applies the GTE predicate on the "connector_id" field.
func ConnectorIDGTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldConnectorID, v))
}

// ConnectorIDLT applies the LT predicate on the "connector_id" field.
func ConnectorIDLT(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldConnectorID, v))
}

// ConnectorIDLTE applies the LTE predicate on the "connector_id" field.
func ConnectorIDLTE(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldConnectorID, v))
}

// ConnectorIDContains applies the Contains predicate on the "connector_id" field.
func ConnectorIDContains(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContains(FieldConnectorID, v))
}

// ConnectorIDHasPrefix applies the HasPrefix predicate on the "connector_id" field.
func ConnectorIDHasPrefix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasPrefix(FieldConnectorID, v))
}

// ConnectorIDHasSuffix applies the HasSuffix predicate on the "connector_id" field.
func ConnectorIDHasSuffix(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldHasSuffix(FieldConnectorID, v))
}

// ConnectorIDEqualFold applies the EqualFold predicate on the "connector_id" field.
func ConnectorIDEqualFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEqualFold(FieldConnectorID, v))
}

// ConnectorIDContainsFold applies the ContainsFold predicate on the "connector_id" field.
func ConnectorIDContainsFold(v string) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldContainsFold(FieldConnectorID, v))
}

// ConnectorDataIsNil applies the IsNil predicate on the "connector_data" field.
func ConnectorDataIsNil() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIsNull(FieldConnectorData))
}

// ConnectorDataNotNil applies the NotNil predicate on the "connector_data" field.
func ConnectorDataNotNil() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotNull(FieldConnectorData))
}

// LastUsedEQ applies the EQ predicate on the "last_used" field.
func LastUsedEQ(v time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldEQ(FieldLastUsed, v))
}

// LastUsedNEQ applies the NEQ predicate on the "last_used" field.
func LastUsedNEQ(v time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNEQ(FieldLastUsed, v))
}

// LastUsedIn applies the In predicate on the "last_used" field.
func LastUsedIn(vs ...time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldIn(FieldLastUsed, vs...))
}

// LastUsedNotIn applies the NotIn predicate on the "last_used" field.
func LastUsedNotIn(vs ...time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldNotIn(FieldLastUsed, vs...))
}

// LastUsedGT applies the GT predicate on the "last_used" field.
func LastUsedGT(v time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGT(FieldLastUsed, v))
}

// LastUsedGTE applies the GTE predicate on the "last_used" field.
func LastUsedGTE(v time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldGTE(FieldLastUsed, v))
}

// LastUsedLT applies the LT predicate on the "last_used" field.
func LastUsedLT(v time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLT(FieldLastUsed, v))
}

// LastUsedLTE applies the LTE predicate on the "last_used" field.
func LastUsedLTE(v time.Time) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.FieldLTE(FieldLastUsed, v))
}

// HasIntegration applies the HasEdge predicate on the "integration" edge.
func HasIntegration() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, IntegrationTable, IntegrationPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Integration
		step.Edge.Schema = schemaConfig.IntegrationOauth2tokens
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIntegrationWith applies the HasEdge predicate on the "integration" edge with a given conditions (other predicates).
func HasIntegrationWith(preds ...predicate.Integration) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(func(s *sql.Selector) {
		step := newIntegrationStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Integration
		step.Edge.Schema = schemaConfig.IntegrationOauth2tokens
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, EventsTable, EventsPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Event
		step.Edge.Schema = schemaConfig.OhAuthTooTokenEvents
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(func(s *sql.Selector) {
		step := newEventsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Event
		step.Edge.Schema = schemaConfig.OhAuthTooTokenEvents
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OhAuthTooToken) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OhAuthTooToken) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OhAuthTooToken) predicate.OhAuthTooToken {
	return predicate.OhAuthTooToken(sql.NotPredicates(p))
}
