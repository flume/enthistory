// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"time"

	"_examples/uuidmixinid/ent/menuitemhistory"

	"entgo.io/ent/dialect/sql"
)

func (mi *MenuItem) History() *MenuItemHistoryQuery {
	historyClient := NewMenuItemHistoryClient(mi.config)
	return historyClient.Query().Where(menuitemhistory.Ref(mi.ID))
}

func (mih *MenuItemHistory) Next(ctx context.Context) (*MenuItemHistory, error) {
	client := NewMenuItemHistoryClient(mih.config)
	return client.Query().
		Where(
			menuitemhistory.Ref(mih.Ref),
			menuitemhistory.HistoryTimeGT(mih.HistoryTime),
		).
		Order(menuitemhistory.ByHistoryTime()).
		First(ctx)
}

func (mih *MenuItemHistory) Prev(ctx context.Context) (*MenuItemHistory, error) {
	client := NewMenuItemHistoryClient(mih.config)
	return client.Query().
		Where(
			menuitemhistory.Ref(mih.Ref),
			menuitemhistory.HistoryTimeLT(mih.HistoryTime),
		).
		Order(menuitemhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (mihq *MenuItemHistoryQuery) Earliest(ctx context.Context) (*MenuItemHistory, error) {
	return mihq.
		Order(menuitemhistory.ByHistoryTime()).
		First(ctx)
}

func (mihq *MenuItemHistoryQuery) Latest(ctx context.Context) (*MenuItemHistory, error) {
	return mihq.
		Order(menuitemhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (mihq *MenuItemHistoryQuery) AsOf(ctx context.Context, time time.Time) (*MenuItemHistory, error) {
	return mihq.
		Where(menuitemhistory.HistoryTimeLTE(time)).
		Order(menuitemhistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (mih *MenuItemHistory) Restore(ctx context.Context) (*MenuItem, error) {
	client := NewMenuItemClient(mih.config)
	return client.
		UpdateOneID(mih.Ref).
		SetUpdatedAt(mih.UpdatedAt).
		SetName(mih.Name).
		SetPrice(mih.Price).
		SetDescription(mih.Description).
		Save(ctx)
}
