// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"_examples/testdata/debug/internal/ent/characterhistory"
	"_examples/testdata/debug/internal/ent/friendshiphistory"

	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

type Change struct {
	FieldName string
	Old       any
	New       any
}

func NewChange(fieldName string, old, new any) Change {
	return Change{
		FieldName: fieldName,
		Old:       old,
		New:       new,
	}
}

type HistoryDiff[T any] struct {
	Old     *T
	New     *T
	Changes []Change
}

var (
	MismatchedRefError    = errors.New("cannot take diff of histories with different Refs")
	IdenticalHistoryError = errors.New("cannot take diff of identical history")
)

func (ch *CharacterHistory) changes(new *CharacterHistory) []Change {
	var changes []Change
	if !reflect.DeepEqual(ch.Age, new.Age) {
		changes = append(changes, NewChange(characterhistory.FieldAge, ch.Age, new.Age))
	}
	if !reflect.DeepEqual(ch.Name, new.Name) {
		changes = append(changes, NewChange(characterhistory.FieldName, ch.Name, new.Name))
	}
	if !reflect.DeepEqual(ch.Nicknames, new.Nicknames) {
		changes = append(changes, NewChange(characterhistory.FieldNicknames, ch.Nicknames, new.Nicknames))
	}
	if !reflect.DeepEqual(ch.Info, new.Info) {
		changes = append(changes, NewChange(characterhistory.FieldInfo, ch.Info, new.Info))
	}
	return changes
}

func (ch *CharacterHistory) Diff(history *CharacterHistory) (*HistoryDiff[CharacterHistory], error) {
	if ch.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	chUnix, historyUnix := ch.HistoryTime.UnixMilli(), history.HistoryTime.UnixMilli()
	chOlder := chUnix < historyUnix
	historyOlder := chUnix > historyUnix

	if chOlder {
		return &HistoryDiff[CharacterHistory]{
			Old:     ch,
			New:     history,
			Changes: ch.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[CharacterHistory]{
			Old:     history,
			New:     ch,
			Changes: history.changes(ch),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (fh *FriendshipHistory) changes(new *FriendshipHistory) []Change {
	var changes []Change
	if !reflect.DeepEqual(fh.CharacterID, new.CharacterID) {
		changes = append(changes, NewChange(friendshiphistory.FieldCharacterID, fh.CharacterID, new.CharacterID))
	}
	if !reflect.DeepEqual(fh.FriendID, new.FriendID) {
		changes = append(changes, NewChange(friendshiphistory.FieldFriendID, fh.FriendID, new.FriendID))
	}
	return changes
}

func (fh *FriendshipHistory) Diff(history *FriendshipHistory) (*HistoryDiff[FriendshipHistory], error) {
	if fh.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	fhUnix, historyUnix := fh.HistoryTime.UnixMilli(), history.HistoryTime.UnixMilli()
	fhOlder := fhUnix < historyUnix
	historyOlder := fhUnix > historyUnix

	if fhOlder {
		return &HistoryDiff[FriendshipHistory]{
			Old:     fh,
			New:     history,
			Changes: fh.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[FriendshipHistory]{
			Old:     history,
			New:     fh,
			Changes: history.changes(fh),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (c Change) String(op enthistory.OpType) string {
	var newstr, oldstr string
	if c.New != nil {
		val, err := json.Marshal(c.New)
		if err != nil {
			newstr = fmt.Sprintf("%#v", c.New)
		} else {
			newstr = string(val)
		}
	}
	if c.Old != nil {
		val, err := json.Marshal(c.Old)
		if err != nil {
			oldstr = fmt.Sprintf("%#v", c.Old)
		} else {
			oldstr = string(val)
		}
	}
	switch op {
	case enthistory.OpTypeInsert:
		return fmt.Sprintf("%s: %s", c.FieldName, newstr)
	case enthistory.OpTypeDelete:
		return fmt.Sprintf("%s: %s", c.FieldName, oldstr)
	default:
		return fmt.Sprintf("%s: %s -> %s", c.FieldName, oldstr, newstr)
	}
}

func (c *Client) Audit(ctx context.Context) ([][]string, error) {
	records := [][]string{
		{"Table", "Ref Id", "History Time", "Operation", "Changes", "Updated By"},
	}
	var rec [][]string
	var err error
	rec, err = auditCharacterHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, rec...)

	rec, err = auditFriendshipHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, rec...)

	return records, nil
}

type record struct {
	Table       string
	RefId       any
	HistoryTime time.Time
	Operation   enthistory.OpType
	Changes     []Change
	UpdatedBy   *uuid.UUID
}

func (r *record) toRow() []string {
	row := make([]string, 6)

	row[0] = r.Table
	row[1] = fmt.Sprintf("%v", r.RefId)
	row[2] = r.HistoryTime.Format(time.ANSIC)
	row[3] = r.Operation.String()
	for i, change := range r.Changes {
		if i == 0 {
			row[4] = change.String(r.Operation)
			continue
		}
		row[4] = fmt.Sprintf("%s\n%s", row[4], change.String(r.Operation))
	}
	if r.UpdatedBy != nil {
		row[5] = fmt.Sprintf("%v", *r.UpdatedBy)
	}
	return row
}

type characterhistoryref struct {
	Ref uuid.UUID
}

func auditCharacterHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []characterhistoryref
	client := NewCharacterHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(characterhistory.ByHistoryTime()).
		Select(characterhistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(characterhistory.Ref(currRef.Ref)).
			Order(characterhistory.ByHistoryTime()).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			r := record{
				Table:       "CharacterHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				r.Changes = (&CharacterHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				r.Changes = curr.changes(&CharacterHistory{})
			default:
				if i == 0 {
					r.Changes = (&CharacterHistory{}).changes(curr)
				} else {
					r.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, r.toRow())
		}
	}
	return records, nil
}

type friendshiphistoryref struct {
	Ref uuid.UUID
}

func auditFriendshipHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []friendshiphistoryref
	client := NewFriendshipHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(friendshiphistory.ByHistoryTime()).
		Select(friendshiphistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(friendshiphistory.Ref(currRef.Ref)).
			Order(friendshiphistory.ByHistoryTime()).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			r := record{
				Table:       "FriendshipHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				r.Changes = (&FriendshipHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				r.Changes = curr.changes(&FriendshipHistory{})
			default:
				if i == 0 {
					r.Changes = (&FriendshipHistory{}).changes(curr)
				} else {
					r.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, r.toRow())
		}
	}
	return records, nil
}
