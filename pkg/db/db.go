package db

import (
	"context"
	"log"

	"github.com/art-Hasan/gqlgen-todos/ent"
)

type Tx struct {
	*ent.Tx

	rollbackCtx context.Context
	commit      bool
}

func (t *Tx) Commit() error {
	if err := t.Tx.Commit(); err != nil {
		return err
	}
	t.commit = true
	return nil
}

func (t *Tx) Rollback() {
	if t == nil || t.commit {
		return
	}
	if err := t.Tx.Rollback(); err != nil {
		log.Println("db: failed to rollback")
	}
}

func Start(ctx context.Context, ent *ent.Client) (*Tx, error) {
	etx, err := ent.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{Tx: etx, rollbackCtx: ctx}, nil
}
