// Package example does stuff
package example

// func EntOpToHistoryOp(op ent.Op) entfga.OpType {
// 	switch op {
// 	case ent.OpDelete, ent.OpDeleteOne:
// 		return entfga.OpTypeDelete
// 	case ent.OpUpdate, ent.OpUpdateOne:
// 		return entfga.OpTypeUpdate
// 	default:
// 		return entfga.OpTypeInsert
// 	}
// }

// // CreateTuplesFromCreate with the user, object and role
// func (m *GroupMembershipMutation) CreateTuplesFromCreate(ctx context.Context) error {
// 	userID, _ := m.UserID()
// 	objectID, _ := m.GroupID()
// 	role, _ := m.Role()

// 	// get tuple key
// 	tuple, err := fga.GetTupleKey(userID, "user", objectID, "object", role.String())
// 	if err != nil {
// 		return err
// 	}

// 	if _, err := m.Authz.WriteTupleKeys(ctx, []fga.TupleKey{tuple}, nil); err != nil {
// 		m.Logger.Errorw("failed to create relationship tuple", "error", err)

// 		return err
// 	}

// 	m.Logger.Debugw("created relationship tuples", "relation", role, "object", tuple.Object)

// 	return nil
// }

// // CreateTuplesFromUpdate with the user, object and role and remove old relations
// func (m *GroupMembershipMutation) CreateTuplesFromUpdate(ctx context.Context) error {
// 	// get ids that will be updated
// 	ids, err := m.IDs(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	var (
// 		writes  []fga.TupleKey
// 		deletes []fga.TupleKey
// 	)

// 	oldRole, err := m.OldRole(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	newRole, exists := m.Role()
// 	if !exists {
// 		return entfga.ErrMissingRole
// 	}

// 	if oldRole == newRole {
// 		m.Logger.Debugw("nothing to update, roles are the same", "old_role", oldRole, "new_role", newRole)

// 		return nil
// 	}

// 	// User the IDs of the group memberships and delete all related tuples
// 	for _, id := range ids {
// 		member, err := m.Client().GroupMembership.Get(ctx, id)
// 		if err != nil {
// 			return err
// 		}

// 		d, err := fga.GetTupleKey(member.UserID, "user", member.GroupID, "group", oldRole.String())
// 		if err != nil {
// 			return err
// 		}

// 		deletes = append(deletes, d)

// 		w, err := fga.GetTupleKey(member.UserID, "user", member.GroupID, "group", newRole.String())
// 		if err != nil {
// 			return err
// 		}

// 		writes = append(writes, w)
// 	}

// 	if len(writes) == 0 && len(deletes) == 0 {
// 		m.Logger.Debugw("no relationships to create or delete")

// 		return nil
// 	}

// 	if _, err := m.Authz.WriteTupleKeys(ctx, writes, deletes); err != nil {
// 		m.Logger.Errorw("failed to update relationship tuple", "error", err)

// 		return err
// 	}

// 	return nil
// }

// // CreateTuplesFromDelete with the user, object and role to remove the relation
// func (m *GroupMembershipMutation) CreateTuplesFromDelete(ctx context.Context) error {
// 	// get ids that will be deleted
// 	ids, err := m.IDs(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	tuples := []fga.TupleKey{}

// 	// User the IDs of the group memberships and delete all related tuples
// 	for _, id := range ids {
// 		// this wont work with soft deletes
// 		members, err := m.Client().GroupMembership.Get(ctx, id)
// 		if err != nil {
// 			return err
// 		}

// 		t, err := fga.GetTupleKey(members.UserID, "user", members.GroupID, "group", members.Role.String())
// 		if err != nil {
// 			return err
// 		}

// 		tuples = append(tuples, t)
// 	}

// 	if len(tuples) > 0 {
// 		if _, err := m.Authz.WriteTupleKeys(ctx, nil, tuples); err != nil {
// 			m.Logger.Errorw("failed to delete relationship tuple", "error", err)

// 			return err
// 		}

// 		m.Logger.Debugw("deleted group relationship tuples")
// 	}

// 	return nil
// }
