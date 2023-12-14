package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
)

// HookGroup runs on group mutations to set default values that are not provided
func HookGroup() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.GroupFunc(func(ctx context.Context, mutation *generated.GroupMutation) (generated.Value, error) {
			if name, ok := mutation.Name(); ok {
				displayName, _ := mutation.DisplayName()

				if displayName == "" {
					mutation.SetDisplayName(name)
				}
			}

			if mutation.Op().Is(ent.OpCreate) {
				// if this is empty generate a default group setting schema
				settingID, _ := mutation.SettingID()
				if settingID == "" {
					// sets up default group settings using schema defaults
					orgSettingID, err := defaultGroupSettings(ctx, mutation)
					if err != nil {
						return nil, err
					}

					// add the group setting ID to the input
					mutation.SetSettingID(orgSettingID)
				}
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}

// defaultGroupSettings creates the default group settings for a new group
func defaultGroupSettings(ctx context.Context, group *generated.GroupMutation) (string, error) {
	input := generated.CreateGroupSettingInput{}

	groupSetting, err := group.Client().GroupSetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return "", err
	}

	return groupSetting.ID, nil
}
