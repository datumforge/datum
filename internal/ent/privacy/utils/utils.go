package utils

import (
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

func NewMutationPolicyWithoutNil(source privacy.MutationPolicy) privacy.MutationPolicy {
	newSlice := make(privacy.MutationPolicy, 0, len(source))

	for _, item := range source {
		if item != nil {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}
