// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/flume/enthistory/_examples/basic/ent/characterhistory"
	"github.com/flume/enthistory/_examples/basic/ent/friendshiphistory"
)

func (c *Character) History() *CharacterHistoryQuery {
	historyClient := NewCharacterHistoryClient(c.config)
	return historyClient.Query().Where(characterhistory.Ref(c.ID))
}

func (ch *CharacterHistory) Next(ctx context.Context) (*CharacterHistory, error) {
	client := NewCharacterHistoryClient(ch.config)
	return client.Query().
		Where(
			characterhistory.Ref(ch.Ref),
			characterhistory.HistoryTimeGT(ch.HistoryTime),
		).
		Order(characterhistory.ByHistoryTime()).
		First(ctx)
}

func (ch *CharacterHistory) Prev(ctx context.Context) (*CharacterHistory, error) {
	client := NewCharacterHistoryClient(ch.config)
	return client.Query().
		Where(
			characterhistory.Ref(ch.Ref),
			characterhistory.HistoryTimeLT(ch.HistoryTime),
		).
		Order(characterhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (chq *CharacterHistoryQuery) Earliest(ctx context.Context) (*CharacterHistory, error) {
	return chq.
		Order(characterhistory.ByHistoryTime()).
		First(ctx)
}

func (chq *CharacterHistoryQuery) Latest(ctx context.Context) (*CharacterHistory, error) {
	return chq.
		Order(characterhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (chq *CharacterHistoryQuery) AsOf(ctx context.Context, time time.Time) (*CharacterHistory, error) {
	return chq.
		Where(characterhistory.HistoryTimeLTE(time)).
		Order(characterhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (ch *CharacterHistory) Restore(ctx context.Context) (*Character, error) {
	client := NewCharacterClient(ch.config)
	return client.
		UpdateOneID(ch.Ref).
		SetUpdatedAt(ch.UpdatedAt).
		SetAge(ch.Age).
		SetName(ch.Name).
		Save(ctx)
}

func (f *Friendship) History() *FriendshipHistoryQuery {
	historyClient := NewFriendshipHistoryClient(f.config)
	return historyClient.Query().Where(friendshiphistory.Ref(f.ID))
}

func (fh *FriendshipHistory) Next(ctx context.Context) (*FriendshipHistory, error) {
	client := NewFriendshipHistoryClient(fh.config)
	return client.Query().
		Where(
			friendshiphistory.Ref(fh.Ref),
			friendshiphistory.HistoryTimeGT(fh.HistoryTime),
		).
		Order(friendshiphistory.ByHistoryTime()).
		First(ctx)
}

func (fh *FriendshipHistory) Prev(ctx context.Context) (*FriendshipHistory, error) {
	client := NewFriendshipHistoryClient(fh.config)
	return client.Query().
		Where(
			friendshiphistory.Ref(fh.Ref),
			friendshiphistory.HistoryTimeLT(fh.HistoryTime),
		).
		Order(friendshiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (fhq *FriendshipHistoryQuery) Earliest(ctx context.Context) (*FriendshipHistory, error) {
	return fhq.
		Order(friendshiphistory.ByHistoryTime()).
		First(ctx)
}

func (fhq *FriendshipHistoryQuery) Latest(ctx context.Context) (*FriendshipHistory, error) {
	return fhq.
		Order(friendshiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (fhq *FriendshipHistoryQuery) AsOf(ctx context.Context, time time.Time) (*FriendshipHistory, error) {
	return fhq.
		Where(friendshiphistory.HistoryTimeLTE(time)).
		Order(friendshiphistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (fh *FriendshipHistory) Restore(ctx context.Context) (*Friendship, error) {
	client := NewFriendshipClient(fh.config)
	return client.
		UpdateOneID(fh.Ref).
		SetUpdatedAt(fh.UpdatedAt).
		SetCharacterID(fh.CharacterID).
		SetFriendID(fh.FriendID).
		Save(ctx)
}
