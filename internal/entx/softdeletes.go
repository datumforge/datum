package entx

import "context"

// SoftDeleteSkipKey is used to indicate to allow soft deleted records to be returned in records
// and to skip soft delete on mutations and proceed with a regular delete
type SoftDeleteSkipKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/hooks.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, SoftDeleteSkipKey{}, true)
}

// CheckSkipSoftDelete checks if the SoftDeleteSkipKey is set in the context
func CheckSkipSoftDelete(ctx context.Context) bool {
	return ctx.Value(SoftDeleteSkipKey{}) != nil
}

// SoftDeleteKey is used to indicate a soft delete mutation is in progress
type SoftDeleteKey struct{}

// IsSoftDelete returns a new context that informs the delete is a soft-delete for interceptor/hooks.
func IsSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, SoftDeleteKey{}, true)
}

// CheckIsSoftDelete checks if the softDeleteKey is set in the context
func CheckIsSoftDelete(ctx context.Context) bool {
	return ctx.Value(SoftDeleteKey{}) != nil
}
