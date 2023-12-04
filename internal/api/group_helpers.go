package api

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
)

// defaultUserSettings creates the default user settings for a new user
func (r *mutationResolver) defaultGroupSettings(ctx context.Context) (string, error) {
	input := generated.CreateGroupSettingInput{}

	groupSetting, err := r.client.GroupSetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return "", err
	}

	return groupSetting.ID, nil
}
