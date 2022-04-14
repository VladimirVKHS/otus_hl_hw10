package counter

import (
	"context"
	"database/sql"
	"otus_sn_counters/internal/otusdb"
)

func GetCounter(ctx context.Context, userId int, counter *Counter) error {
	err := otusdb.Db.QueryRowContext(
		ctx,
		"SELECT user_id, unread_messages_count FROM counters WHERE user_id = ?",
		userId,
	).Scan(&counter.UserId, &counter.UnreadMessagesCount)
	return err
}

func CreateCounter(ctx context.Context, userId int) error {
	_, err := otusdb.Db.ExecContext(
		ctx,
		"INSERT INTO counters SET user_id = ?",
		userId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *CounterUpdateRequest) Apply(ctx context.Context, userId int) error {
	tx, err := otusdb.Db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	if r.UnreadMessagesCountDelta != 0 {
		_, err := tx.ExecContext(
			ctx,
			"UPDATE counters SET unread_messages_count =  unread_messages_count + ? where user_id = ?",
			r.UnreadMessagesCountDelta,
			userId,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *CounterUpdateRequest) CheckAndApply(ctx context.Context, userId int) error {
	counter := Counter{}
	if err := GetCounter(ctx, userId, &counter); err != nil {
		if err2 := CreateCounter(ctx, userId); err2 != nil {
			return err2
		}
	}
	if err := r.Apply(ctx, userId); err != nil {
		return err
	}
	return nil
}
