package mixin

import (
	"entgo.io/ent/schema/mixin"
)

// BaseMixin for all schemas in the graph.
type BaseMixin struct {
	mixin.Schema
}

// // Policy defines the privacy policy of the BaseMixin.
// func (BaseMixin) Policy() ent.Policy {
// 	return privacy.Policy{
// 		Query: privacy.QueryPolicy{
// 			rule.AllowIfAdmin(),
// 			privacy.AlwaysDenyRule(),
// 		},
// 		Mutation: privacy.MutationPolicy{
// 			rule.DenyIfNoViewer(),
// 			rule.AllowIfAdmin(),
// 			privacy.AlwaysDenyRule(),
// 		},
// 	}
// }
