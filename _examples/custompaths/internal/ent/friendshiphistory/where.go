// Code generated by ent, DO NOT EDIT.

package friendshiphistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/flume/enthistory"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLTE(FieldID, id))
}

// HistoryTime applies equality check predicate on the "history_time" field. It's identical to HistoryTimeEQ.
func HistoryTime(v time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// Ref applies equality check predicate on the "ref" field. It's identical to RefEQ.
func Ref(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldRef, v))
}

// CharacterID applies equality check predicate on the "character_id" field. It's identical to CharacterIDEQ.
func CharacterID(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldCharacterID, v))
}

// FriendID applies equality check predicate on the "friend_id" field. It's identical to FriendIDEQ.
func FriendID(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldFriendID, v))
}

// HistoryTimeEQ applies the EQ predicate on the "history_time" field.
func HistoryTimeEQ(v time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// HistoryTimeNEQ applies the NEQ predicate on the "history_time" field.
func HistoryTimeNEQ(v time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNEQ(FieldHistoryTime, v))
}

// HistoryTimeIn applies the In predicate on the "history_time" field.
func HistoryTimeIn(vs ...time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldIn(FieldHistoryTime, vs...))
}

// HistoryTimeNotIn applies the NotIn predicate on the "history_time" field.
func HistoryTimeNotIn(vs ...time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNotIn(FieldHistoryTime, vs...))
}

// HistoryTimeGT applies the GT predicate on the "history_time" field.
func HistoryTimeGT(v time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGT(FieldHistoryTime, v))
}

// HistoryTimeGTE applies the GTE predicate on the "history_time" field.
func HistoryTimeGTE(v time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGTE(FieldHistoryTime, v))
}

// HistoryTimeLT applies the LT predicate on the "history_time" field.
func HistoryTimeLT(v time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLT(FieldHistoryTime, v))
}

// HistoryTimeLTE applies the LTE predicate on the "history_time" field.
func HistoryTimeLTE(v time.Time) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLTE(FieldHistoryTime, v))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v enthistory.OpType) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v enthistory.OpType) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...enthistory.OpType) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...enthistory.OpType) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNotIn(FieldOperation, vs...))
}

// RefEQ applies the EQ predicate on the "ref" field.
func RefEQ(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldRef, v))
}

// RefNEQ applies the NEQ predicate on the "ref" field.
func RefNEQ(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNEQ(FieldRef, v))
}

// RefIn applies the In predicate on the "ref" field.
func RefIn(vs ...int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldIn(FieldRef, vs...))
}

// RefNotIn applies the NotIn predicate on the "ref" field.
func RefNotIn(vs ...int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNotIn(FieldRef, vs...))
}

// RefGT applies the GT predicate on the "ref" field.
func RefGT(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGT(FieldRef, v))
}

// RefGTE applies the GTE predicate on the "ref" field.
func RefGTE(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGTE(FieldRef, v))
}

// RefLT applies the LT predicate on the "ref" field.
func RefLT(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLT(FieldRef, v))
}

// RefLTE applies the LTE predicate on the "ref" field.
func RefLTE(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLTE(FieldRef, v))
}

// RefIsNil applies the IsNil predicate on the "ref" field.
func RefIsNil() predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldIsNull(FieldRef))
}

// RefNotNil applies the NotNil predicate on the "ref" field.
func RefNotNil() predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNotNull(FieldRef))
}

// CharacterIDEQ applies the EQ predicate on the "character_id" field.
func CharacterIDEQ(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldCharacterID, v))
}

// CharacterIDNEQ applies the NEQ predicate on the "character_id" field.
func CharacterIDNEQ(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNEQ(FieldCharacterID, v))
}

// CharacterIDIn applies the In predicate on the "character_id" field.
func CharacterIDIn(vs ...int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldIn(FieldCharacterID, vs...))
}

// CharacterIDNotIn applies the NotIn predicate on the "character_id" field.
func CharacterIDNotIn(vs ...int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNotIn(FieldCharacterID, vs...))
}

// CharacterIDGT applies the GT predicate on the "character_id" field.
func CharacterIDGT(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGT(FieldCharacterID, v))
}

// CharacterIDGTE applies the GTE predicate on the "character_id" field.
func CharacterIDGTE(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGTE(FieldCharacterID, v))
}

// CharacterIDLT applies the LT predicate on the "character_id" field.
func CharacterIDLT(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLT(FieldCharacterID, v))
}

// CharacterIDLTE applies the LTE predicate on the "character_id" field.
func CharacterIDLTE(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLTE(FieldCharacterID, v))
}

// FriendIDEQ applies the EQ predicate on the "friend_id" field.
func FriendIDEQ(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldEQ(FieldFriendID, v))
}

// FriendIDNEQ applies the NEQ predicate on the "friend_id" field.
func FriendIDNEQ(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNEQ(FieldFriendID, v))
}

// FriendIDIn applies the In predicate on the "friend_id" field.
func FriendIDIn(vs ...int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldIn(FieldFriendID, vs...))
}

// FriendIDNotIn applies the NotIn predicate on the "friend_id" field.
func FriendIDNotIn(vs ...int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldNotIn(FieldFriendID, vs...))
}

// FriendIDGT applies the GT predicate on the "friend_id" field.
func FriendIDGT(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGT(FieldFriendID, v))
}

// FriendIDGTE applies the GTE predicate on the "friend_id" field.
func FriendIDGTE(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldGTE(FieldFriendID, v))
}

// FriendIDLT applies the LT predicate on the "friend_id" field.
func FriendIDLT(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLT(FieldFriendID, v))
}

// FriendIDLTE applies the LTE predicate on the "friend_id" field.
func FriendIDLTE(v int) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.FieldLTE(FieldFriendID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FriendshipHistory) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FriendshipHistory) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FriendshipHistory) predicate.FriendshipHistory {
	return predicate.FriendshipHistory(sql.NotPredicates(p))
}
