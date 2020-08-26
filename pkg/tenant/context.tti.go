// Copyright © 2019 The Things Industries B.V.

package tenant

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttipb"
)

type tenantIDKeyType struct{}

var tenantIDKey = tenantIDKeyType{}

// FromContext returns the current TenantIdentifier based on the given context.
// Returns empty identifier if not found.
func FromContext(ctx context.Context) ttipb.TenantIdentifiers {
	if id, ok := ctx.Value(tenantIDKey).(ttipb.TenantIdentifiers); ok { // set by NewContext
		return id
	}
	return ttipb.TenantIdentifiers{}
}

// NewContext returns a context containing the tenant identifier.
func NewContext(parent context.Context, id ttipb.TenantIdentifiers) context.Context {
	ctx := context.WithValue(parent, tenantIDKey, id)
	return log.NewContextWithField(ctx, "tenant_id", id.TenantID)
}