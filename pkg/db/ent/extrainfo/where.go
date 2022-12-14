// Code generated by ent, DO NOT EDIT.

package extrainfo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// VoteCount applies equality check predicate on the "vote_count" field. It's identical to VoteCountEQ.
func VoteCount(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVoteCount), v))
	})
}

// Rating applies equality check predicate on the "rating" field. It's identical to RatingEQ.
func Rating(v float32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRating), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// PostersIsNil applies the IsNil predicate on the "posters" field.
func PostersIsNil() predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPosters)))
	})
}

// PostersNotNil applies the NotNil predicate on the "posters" field.
func PostersNotNil() predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPosters)))
	})
}

// LabelsIsNil applies the IsNil predicate on the "labels" field.
func LabelsIsNil() predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLabels)))
	})
}

// LabelsNotNil applies the NotNil predicate on the "labels" field.
func LabelsNotNil() predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLabels)))
	})
}

// VoteCountEQ applies the EQ predicate on the "vote_count" field.
func VoteCountEQ(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVoteCount), v))
	})
}

// VoteCountNEQ applies the NEQ predicate on the "vote_count" field.
func VoteCountNEQ(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVoteCount), v))
	})
}

// VoteCountIn applies the In predicate on the "vote_count" field.
func VoteCountIn(vs ...uint32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVoteCount), v...))
	})
}

// VoteCountNotIn applies the NotIn predicate on the "vote_count" field.
func VoteCountNotIn(vs ...uint32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVoteCount), v...))
	})
}

// VoteCountGT applies the GT predicate on the "vote_count" field.
func VoteCountGT(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVoteCount), v))
	})
}

// VoteCountGTE applies the GTE predicate on the "vote_count" field.
func VoteCountGTE(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVoteCount), v))
	})
}

// VoteCountLT applies the LT predicate on the "vote_count" field.
func VoteCountLT(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVoteCount), v))
	})
}

// VoteCountLTE applies the LTE predicate on the "vote_count" field.
func VoteCountLTE(v uint32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVoteCount), v))
	})
}

// VoteCountIsNil applies the IsNil predicate on the "vote_count" field.
func VoteCountIsNil() predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldVoteCount)))
	})
}

// VoteCountNotNil applies the NotNil predicate on the "vote_count" field.
func VoteCountNotNil() predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldVoteCount)))
	})
}

// RatingEQ applies the EQ predicate on the "rating" field.
func RatingEQ(v float32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRating), v))
	})
}

// RatingNEQ applies the NEQ predicate on the "rating" field.
func RatingNEQ(v float32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRating), v))
	})
}

// RatingIn applies the In predicate on the "rating" field.
func RatingIn(vs ...float32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRating), v...))
	})
}

// RatingNotIn applies the NotIn predicate on the "rating" field.
func RatingNotIn(vs ...float32) predicate.ExtraInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRating), v...))
	})
}

// RatingGT applies the GT predicate on the "rating" field.
func RatingGT(v float32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRating), v))
	})
}

// RatingGTE applies the GTE predicate on the "rating" field.
func RatingGTE(v float32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRating), v))
	})
}

// RatingLT applies the LT predicate on the "rating" field.
func RatingLT(v float32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRating), v))
	})
}

// RatingLTE applies the LTE predicate on the "rating" field.
func RatingLTE(v float32) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRating), v))
	})
}

// RatingIsNil applies the IsNil predicate on the "rating" field.
func RatingIsNil() predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRating)))
	})
}

// RatingNotNil applies the NotNil predicate on the "rating" field.
func RatingNotNil() predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRating)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExtraInfo) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExtraInfo) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
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
func Not(p predicate.ExtraInfo) predicate.ExtraInfo {
	return predicate.ExtraInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}
