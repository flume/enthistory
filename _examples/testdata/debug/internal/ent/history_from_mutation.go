// Code generated by enthistory, DO NOT EDIT.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent"

	"github.com/flume/enthistory"

	"github.com/google/uuid"
)

var (
	idNotFoundError = errors.New("could not get id from mutation")
)

func EntOpToHistoryOp(op ent.Op) enthistory.OpType {
	switch op {
	case ent.OpDelete, ent.OpDeleteOne:
		return enthistory.OpTypeDelete
	case ent.OpUpdate, ent.OpUpdateOne:
		return enthistory.OpTypeUpdate
	default:
		return enthistory.OpTypeInsert
	}
}

func rollback(tx *Tx, err error) error {
	if tx != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}
	return err
}

func (m *CharacterMutation) CreateHistoryFromCreate(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userid").(uuid.UUID)

	id, ok := m.ID()
	if !ok {
		return rollback(tx, idNotFoundError)
	}

	create := client.CharacterHistory.Create()
	if tx != nil {
		create = tx.CharacterHistory.Create()
	}

	create = create.
		SetOperation(EntOpToHistoryOp(m.Op())).
		SetHistoryTime(time.Now()).
		SetRef(id)
	if updatedBy != uuid.Nil {
		create = create.SetUpdatedBy(updatedBy)
	}

	if age, exists := m.Age(); exists {
		create = create.SetAge(age)
	}

	if name, exists := m.Name(); exists {
		create = create.SetName(name)
	}

	if nicknames, exists := m.Nicknames(); exists {
		create = create.SetNicknames(nicknames)
	}

	if info, exists := m.Info(); exists {
		create = create.SetInfo(info)
	}

	_, err = create.Save(ctx)
	if err != nil {
		rollback(tx, err)
	}
	return nil
}

func (m *CharacterMutation) CreateHistoryFromUpdate(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userid").(uuid.UUID)

	ids, err := m.IDs(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("getting ids: %w", err))
	}

	for _, id := range ids {
		character, err := client.Character.Get(ctx, id)
		if err != nil {
			return rollback(tx, err)
		}

		create := client.CharacterHistory.Create()
		if tx != nil {
			create = tx.CharacterHistory.Create()
		}

		create = create.
			SetOperation(EntOpToHistoryOp(m.Op())).
			SetHistoryTime(time.Now()).
			SetRef(id)
		if updatedBy != uuid.Nil {
			create = create.SetUpdatedBy(updatedBy)
		}

		if age, exists := m.Age(); exists {
			create = create.SetAge(age)
		} else {
			create = create.SetAge(character.Age)
		}

		if name, exists := m.Name(); exists {
			create = create.SetName(name)
		} else {
			create = create.SetName(character.Name)
		}

		if nicknames, exists := m.Nicknames(); exists {
			create = create.SetNicknames(nicknames)
		} else {
			create = create.SetNicknames(character.Nicknames)
		}

		if info, exists := m.Info(); exists {
			create = create.SetInfo(info)
		} else {
			create = create.SetInfo(character.Info)
		}

		_, err = create.Save(ctx)
		if err != nil {
			rollback(tx, err)
		}
	}

	return nil
}

func (m *CharacterMutation) CreateHistoryFromDelete(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userid").(uuid.UUID)

	ids, err := m.IDs(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("getting ids: %w", err))
	}

	for _, id := range ids {
		character, err := client.Character.Get(ctx, id)
		if err != nil {
			return rollback(tx, err)
		}

		create := client.CharacterHistory.Create()
		if tx != nil {
			create = tx.CharacterHistory.Create()
		}
		if updatedBy != uuid.Nil {
			create = create.SetUpdatedBy(updatedBy)
		}

		_, err = create.
			SetOperation(EntOpToHistoryOp(m.Op())).
			SetHistoryTime(time.Now()).
			SetRef(id).
			SetAge(character.Age).
			SetName(character.Name).
			SetNicknames(character.Nicknames).
			SetInfo(character.Info).
			Save(ctx)
		if err != nil {
			rollback(tx, err)
		}
	}

	return nil
}

func (m *FriendshipMutation) CreateHistoryFromCreate(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userid").(uuid.UUID)

	id, ok := m.ID()
	if !ok {
		return rollback(tx, idNotFoundError)
	}

	create := client.FriendshipHistory.Create()
	if tx != nil {
		create = tx.FriendshipHistory.Create()
	}

	create = create.
		SetOperation(EntOpToHistoryOp(m.Op())).
		SetHistoryTime(time.Now()).
		SetRef(id)
	if updatedBy != uuid.Nil {
		create = create.SetUpdatedBy(updatedBy)
	}

	if characterID, exists := m.CharacterID(); exists {
		create = create.SetCharacterID(characterID)
	}

	if friendID, exists := m.FriendID(); exists {
		create = create.SetFriendID(friendID)
	}

	_, err = create.Save(ctx)
	if err != nil {
		rollback(tx, err)
	}
	return nil
}

func (m *FriendshipMutation) CreateHistoryFromUpdate(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userid").(uuid.UUID)

	ids, err := m.IDs(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("getting ids: %w", err))
	}

	for _, id := range ids {
		friendship, err := client.Friendship.Get(ctx, id)
		if err != nil {
			return rollback(tx, err)
		}

		create := client.FriendshipHistory.Create()
		if tx != nil {
			create = tx.FriendshipHistory.Create()
		}

		create = create.
			SetOperation(EntOpToHistoryOp(m.Op())).
			SetHistoryTime(time.Now()).
			SetRef(id)
		if updatedBy != uuid.Nil {
			create = create.SetUpdatedBy(updatedBy)
		}

		if characterID, exists := m.CharacterID(); exists {
			create = create.SetCharacterID(characterID)
		} else {
			create = create.SetCharacterID(friendship.CharacterID)
		}

		if friendID, exists := m.FriendID(); exists {
			create = create.SetFriendID(friendID)
		} else {
			create = create.SetFriendID(friendship.FriendID)
		}

		_, err = create.Save(ctx)
		if err != nil {
			rollback(tx, err)
		}
	}

	return nil
}

func (m *FriendshipMutation) CreateHistoryFromDelete(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userid").(uuid.UUID)

	ids, err := m.IDs(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("getting ids: %w", err))
	}

	for _, id := range ids {
		friendship, err := client.Friendship.Get(ctx, id)
		if err != nil {
			return rollback(tx, err)
		}

		create := client.FriendshipHistory.Create()
		if tx != nil {
			create = tx.FriendshipHistory.Create()
		}
		if updatedBy != uuid.Nil {
			create = create.SetUpdatedBy(updatedBy)
		}

		_, err = create.
			SetOperation(EntOpToHistoryOp(m.Op())).
			SetHistoryTime(time.Now()).
			SetRef(id).
			SetCharacterID(friendship.CharacterID).
			SetFriendID(friendship.FriendID).
			Save(ctx)
		if err != nil {
			rollback(tx, err)
		}
	}

	return nil
}
