// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (dc *DeathCertificateQuery) CollectFields(ctx context.Context, satisfies ...string) *DeathCertificateQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		dc = dc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return dc
}

func (dc *DeathCertificateQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DeathCertificateQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "certified_by":
			dc = dc.WithCertifiedBy(func(query *PersonQuery) {
				query.collectField(ctx, field)
			})
		case "certifier":
			dc = dc.WithCertifier(func(query *DeathCertifierQuery) {
				query.collectField(ctx, field)
			})
		case "decedent":
			dc = dc.WithDecedent(func(query *PersonQuery) {
				query.collectField(ctx, field)
			})
		case "manner":
			dc = dc.WithManner(func(query *DeathMannerQuery) {
				query.collectField(ctx, field)
			})
		case "metadata":
			dc = dc.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		case "place":
			dc = dc.WithPlace(func(query *DeathPlaceQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return dc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (dc *DeathCertifierQuery) CollectFields(ctx context.Context, satisfies ...string) *DeathCertifierQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		dc = dc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return dc
}

func (dc *DeathCertifierQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DeathCertifierQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "death_certificates":
			dc = dc.WithDeathCertificates(func(query *DeathCertificateQuery) {
				query.collectField(ctx, field)
			})
		case "metadata":
			dc = dc.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return dc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (dm *DeathMannerQuery) CollectFields(ctx context.Context, satisfies ...string) *DeathMannerQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		dm = dm.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return dm
}

func (dm *DeathMannerQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DeathMannerQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "death_certificates":
			dm = dm.WithDeathCertificates(func(query *DeathCertificateQuery) {
				query.collectField(ctx, field)
			})
		case "metadata":
			dm = dm.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return dm
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (dp *DeathPlaceQuery) CollectFields(ctx context.Context, satisfies ...string) *DeathPlaceQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		dp = dp.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return dp
}

func (dp *DeathPlaceQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DeathPlaceQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "death_certificates":
			dp = dp.WithDeathCertificates(func(query *DeathCertificateQuery) {
				query.collectField(ctx, field)
			})
		case "metadata":
			dp = dp.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return dp
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (e *EthnicityQuery) CollectFields(ctx context.Context, satisfies ...string) *EthnicityQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		e = e.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return e
}

func (e *EthnicityQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *EthnicityQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			e = e.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return e
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gs *GameServerQuery) CollectFields(ctx context.Context, satisfies ...string) *GameServerQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		gs = gs.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return gs
}

func (gs *GameServerQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GameServerQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			gs = gs.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return gs
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MetadataQuery) CollectFields(ctx context.Context, satisfies ...string) *MetadataQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		m = m.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return m
}

func (m *MetadataQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *MetadataQuery {
	return m
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (oc *OAuthConnectionQuery) CollectFields(ctx context.Context, satisfies ...string) *OAuthConnectionQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		oc = oc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return oc
}

func (oc *OAuthConnectionQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *OAuthConnectionQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "user":
			oc = oc.WithUser(func(query *UserQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return oc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pe *PersonQuery) CollectFields(ctx context.Context, satisfies ...string) *PersonQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pe = pe.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pe
}

func (pe *PersonQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PersonQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			pe = pe.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return pe
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pl *PlayerQuery) CollectFields(ctx context.Context, satisfies ...string) *PlayerQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pl = pl.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pl
}

func (pl *PlayerQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PlayerQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			pl = pl.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return pl
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pi *PlayerIdentifierQuery) CollectFields(ctx context.Context, satisfies ...string) *PlayerIdentifierQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pi = pi.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pi
}

func (pi *PlayerIdentifierQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PlayerIdentifierQuery {
	return pi
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RaceQuery) CollectFields(ctx context.Context, satisfies ...string) *RaceQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		r = r.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return r
}

func (r *RaceQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *RaceQuery {
	return r
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SessionQuery) CollectFields(ctx context.Context, satisfies ...string) *SessionQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		s = s.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return s
}

func (s *SessionQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *SessionQuery {
	return s
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SettingQuery) CollectFields(ctx context.Context, satisfies ...string) *SettingQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		s = s.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return s
}

func (s *SettingQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *SettingQuery {
	return s
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SexQuery) CollectFields(ctx context.Context, satisfies ...string) *SexQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		s = s.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return s
}

func (s *SexQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *SexQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			s = s.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return s
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *StateQuery) CollectFields(ctx context.Context, satisfies ...string) *StateQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		s = s.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return s
}

func (s *StateQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *StateQuery {
	return s
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			u = u.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		case "oauth_connections":
			u = u.WithOauthConnections(func(query *OAuthConnectionQuery) {
				query.collectField(ctx, field)
			})
		case "people":
			u = u.WithPeople(func(query *PersonQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return u
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (v *VehicleQuery) CollectFields(ctx context.Context, satisfies ...string) *VehicleQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		v = v.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return v
}

func (v *VehicleQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *VehicleQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			v = v.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return v
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (vc *VehicleClassQuery) CollectFields(ctx context.Context, satisfies ...string) *VehicleClassQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		vc = vc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return vc
}

func (vc *VehicleClassQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *VehicleClassQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			vc = vc.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return vc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (vc *VehicleColorQuery) CollectFields(ctx context.Context, satisfies ...string) *VehicleColorQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		vc = vc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return vc
}

func (vc *VehicleColorQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *VehicleColorQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			vc = vc.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return vc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (vm *VehicleMakeQuery) CollectFields(ctx context.Context, satisfies ...string) *VehicleMakeQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		vm = vm.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return vm
}

func (vm *VehicleMakeQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *VehicleMakeQuery {
	return vm
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (vm *VehicleModelQuery) CollectFields(ctx context.Context, satisfies ...string) *VehicleModelQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		vm = vm.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return vm
}

func (vm *VehicleModelQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *VehicleModelQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			vm = vm.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return vm
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (vr *VehicleRegistrationQuery) CollectFields(ctx context.Context, satisfies ...string) *VehicleRegistrationQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		vr = vr.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return vr
}

func (vr *VehicleRegistrationQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *VehicleRegistrationQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "metadata":
			vr = vr.WithMetadata(func(query *MetadataQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return vr
}
