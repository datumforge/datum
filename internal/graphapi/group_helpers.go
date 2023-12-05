package graphapi

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
)

// defaultGroupSettings creates the default group settings for a new group
func (r *mutationResolver) defaultGroupSettings(ctx context.Context) (string, error) {
	input := generated.CreateGroupSettingInput{}

	groupSetting, err := r.client.GroupSetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return "", err
	}

	return groupSetting.ID, nil
}
