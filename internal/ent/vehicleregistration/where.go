// Code generated by entc, DO NOT EDIT.

package vehicleregistration

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/responserms/response/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedWith applies equality check predicate on the "created_with" field. It's identical to CreatedWithEQ.
func CreatedWith(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedWith), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedWith applies equality check predicate on the "updated_with" field. It's identical to UpdatedWithEQ.
func UpdatedWith(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedWith), v))
	})
}

// Plate applies equality check predicate on the "plate" field. It's identical to PlateEQ.
func Plate(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlate), v))
	})
}

// ExpiredAt applies equality check predicate on the "expired_at" field. It's identical to ExpiredAtEQ.
func ExpiredAt(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiredAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedBy)))
	})
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedBy)))
	})
}

// CreatedWithEQ applies the EQ predicate on the "created_with" field.
func CreatedWithEQ(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithNEQ applies the NEQ predicate on the "created_with" field.
func CreatedWithNEQ(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithIn applies the In predicate on the "created_with" field.
func CreatedWithIn(vs ...string) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedWith), v...))
	})
}

// CreatedWithNotIn applies the NotIn predicate on the "created_with" field.
func CreatedWithNotIn(vs ...string) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedWith), v...))
	})
}

// CreatedWithGT applies the GT predicate on the "created_with" field.
func CreatedWithGT(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithGTE applies the GTE predicate on the "created_with" field.
func CreatedWithGTE(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithLT applies the LT predicate on the "created_with" field.
func CreatedWithLT(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithLTE applies the LTE predicate on the "created_with" field.
func CreatedWithLTE(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithContains applies the Contains predicate on the "created_with" field.
func CreatedWithContains(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithHasPrefix applies the HasPrefix predicate on the "created_with" field.
func CreatedWithHasPrefix(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithHasSuffix applies the HasSuffix predicate on the "created_with" field.
func CreatedWithHasSuffix(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithIsNil applies the IsNil predicate on the "created_with" field.
func CreatedWithIsNil() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedWith)))
	})
}

// CreatedWithNotNil applies the NotNil predicate on the "created_with" field.
func CreatedWithNotNil() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedWith)))
	})
}

// CreatedWithEqualFold applies the EqualFold predicate on the "created_with" field.
func CreatedWithEqualFold(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreatedWith), v))
	})
}

// CreatedWithContainsFold applies the ContainsFold predicate on the "created_with" field.
func CreatedWithContainsFold(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreatedWith), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedBy)))
	})
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedBy)))
	})
}

// UpdatedWithEQ applies the EQ predicate on the "updated_with" field.
func UpdatedWithEQ(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithNEQ applies the NEQ predicate on the "updated_with" field.
func UpdatedWithNEQ(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithIn applies the In predicate on the "updated_with" field.
func UpdatedWithIn(vs ...string) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedWith), v...))
	})
}

// UpdatedWithNotIn applies the NotIn predicate on the "updated_with" field.
func UpdatedWithNotIn(vs ...string) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedWith), v...))
	})
}

// UpdatedWithGT applies the GT predicate on the "updated_with" field.
func UpdatedWithGT(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithGTE applies the GTE predicate on the "updated_with" field.
func UpdatedWithGTE(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithLT applies the LT predicate on the "updated_with" field.
func UpdatedWithLT(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithLTE applies the LTE predicate on the "updated_with" field.
func UpdatedWithLTE(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithContains applies the Contains predicate on the "updated_with" field.
func UpdatedWithContains(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithHasPrefix applies the HasPrefix predicate on the "updated_with" field.
func UpdatedWithHasPrefix(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithHasSuffix applies the HasSuffix predicate on the "updated_with" field.
func UpdatedWithHasSuffix(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithIsNil applies the IsNil predicate on the "updated_with" field.
func UpdatedWithIsNil() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedWith)))
	})
}

// UpdatedWithNotNil applies the NotNil predicate on the "updated_with" field.
func UpdatedWithNotNil() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedWith)))
	})
}

// UpdatedWithEqualFold applies the EqualFold predicate on the "updated_with" field.
func UpdatedWithEqualFold(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUpdatedWith), v))
	})
}

// UpdatedWithContainsFold applies the ContainsFold predicate on the "updated_with" field.
func UpdatedWithContainsFold(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUpdatedWith), v))
	})
}

// PlateEQ applies the EQ predicate on the "plate" field.
func PlateEQ(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlate), v))
	})
}

// PlateNEQ applies the NEQ predicate on the "plate" field.
func PlateNEQ(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlate), v))
	})
}

// PlateIn applies the In predicate on the "plate" field.
func PlateIn(vs ...string) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlate), v...))
	})
}

// PlateNotIn applies the NotIn predicate on the "plate" field.
func PlateNotIn(vs ...string) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlate), v...))
	})
}

// PlateGT applies the GT predicate on the "plate" field.
func PlateGT(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlate), v))
	})
}

// PlateGTE applies the GTE predicate on the "plate" field.
func PlateGTE(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlate), v))
	})
}

// PlateLT applies the LT predicate on the "plate" field.
func PlateLT(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlate), v))
	})
}

// PlateLTE applies the LTE predicate on the "plate" field.
func PlateLTE(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlate), v))
	})
}

// PlateContains applies the Contains predicate on the "plate" field.
func PlateContains(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPlate), v))
	})
}

// PlateHasPrefix applies the HasPrefix predicate on the "plate" field.
func PlateHasPrefix(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPlate), v))
	})
}

// PlateHasSuffix applies the HasSuffix predicate on the "plate" field.
func PlateHasSuffix(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPlate), v))
	})
}

// PlateEqualFold applies the EqualFold predicate on the "plate" field.
func PlateEqualFold(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPlate), v))
	})
}

// PlateContainsFold applies the ContainsFold predicate on the "plate" field.
func PlateContainsFold(v string) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPlate), v))
	})
}

// ExpiredAtEQ applies the EQ predicate on the "expired_at" field.
func ExpiredAtEQ(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtNEQ applies the NEQ predicate on the "expired_at" field.
func ExpiredAtNEQ(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtIn applies the In predicate on the "expired_at" field.
func ExpiredAtIn(vs ...time.Time) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiredAt), v...))
	})
}

// ExpiredAtNotIn applies the NotIn predicate on the "expired_at" field.
func ExpiredAtNotIn(vs ...time.Time) predicate.VehicleRegistration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiredAt), v...))
	})
}

// ExpiredAtGT applies the GT predicate on the "expired_at" field.
func ExpiredAtGT(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtGTE applies the GTE predicate on the "expired_at" field.
func ExpiredAtGTE(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtLT applies the LT predicate on the "expired_at" field.
func ExpiredAtLT(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtLTE applies the LTE predicate on the "expired_at" field.
func ExpiredAtLTE(v time.Time) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiredAt), v))
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasState applies the HasEdge predicate on the "state" edge.
func HasState() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StateTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StateTable, StateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStateWith applies the HasEdge predicate on the "state" edge with a given conditions (other predicates).
func HasStateWith(preds ...predicate.State) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StateInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StateTable, StateColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPerson applies the HasEdge predicate on the "person" edge.
func HasPerson() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonTable, PersonColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonWith applies the HasEdge predicate on the "person" edge with a given conditions (other predicates).
func HasPersonWith(preds ...predicate.Person) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonTable, PersonColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVehicle applies the HasEdge predicate on the "vehicle" edge.
func HasVehicle() predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VehicleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VehicleTable, VehicleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVehicleWith applies the HasEdge predicate on the "vehicle" edge with a given conditions (other predicates).
func HasVehicleWith(preds ...predicate.Vehicle) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VehicleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VehicleTable, VehicleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VehicleRegistration) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VehicleRegistration) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VehicleRegistration) predicate.VehicleRegistration {
	return predicate.VehicleRegistration(func(s *sql.Selector) {
		p(s.Not())
	})
}
