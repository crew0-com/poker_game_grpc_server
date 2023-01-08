// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addGameRoomPlayerStmt, err = db.PrepareContext(ctx, addGameRoomPlayer); err != nil {
		return nil, fmt.Errorf("error preparing query AddGameRoomPlayer: %w", err)
	}
	if q.createGameStmt, err = db.PrepareContext(ctx, createGame); err != nil {
		return nil, fmt.Errorf("error preparing query CreateGame: %w", err)
	}
	if q.createGameRoomStmt, err = db.PrepareContext(ctx, createGameRoom); err != nil {
		return nil, fmt.Errorf("error preparing query CreateGameRoom: %w", err)
	}
	if q.createPlayerStmt, err = db.PrepareContext(ctx, createPlayer); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePlayer: %w", err)
	}
	if q.finishStmt, err = db.PrepareContext(ctx, finish); err != nil {
		return nil, fmt.Errorf("error preparing query Finish: %w", err)
	}
	if q.getActiveGameByRoomIdStmt, err = db.PrepareContext(ctx, getActiveGameByRoomId); err != nil {
		return nil, fmt.Errorf("error preparing query GetActiveGameByRoomId: %w", err)
	}
	if q.getGameStmt, err = db.PrepareContext(ctx, getGame); err != nil {
		return nil, fmt.Errorf("error preparing query GetGame: %w", err)
	}
	if q.getGameByRoomIdStmt, err = db.PrepareContext(ctx, getGameByRoomId); err != nil {
		return nil, fmt.Errorf("error preparing query GetGameByRoomId: %w", err)
	}
	if q.getGameRoomWithPlayersStmt, err = db.PrepareContext(ctx, getGameRoomWithPlayers); err != nil {
		return nil, fmt.Errorf("error preparing query GetGameRoomWithPlayers: %w", err)
	}
	if q.getPlayerStmt, err = db.PrepareContext(ctx, getPlayer); err != nil {
		return nil, fmt.Errorf("error preparing query GetPlayer: %w", err)
	}
	if q.startGameStmt, err = db.PrepareContext(ctx, startGame); err != nil {
		return nil, fmt.Errorf("error preparing query StartGame: %w", err)
	}
	if q.updateGameStmt, err = db.PrepareContext(ctx, updateGame); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateGame: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addGameRoomPlayerStmt != nil {
		if cerr := q.addGameRoomPlayerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addGameRoomPlayerStmt: %w", cerr)
		}
	}
	if q.createGameStmt != nil {
		if cerr := q.createGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createGameStmt: %w", cerr)
		}
	}
	if q.createGameRoomStmt != nil {
		if cerr := q.createGameRoomStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createGameRoomStmt: %w", cerr)
		}
	}
	if q.createPlayerStmt != nil {
		if cerr := q.createPlayerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPlayerStmt: %w", cerr)
		}
	}
	if q.finishStmt != nil {
		if cerr := q.finishStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing finishStmt: %w", cerr)
		}
	}
	if q.getActiveGameByRoomIdStmt != nil {
		if cerr := q.getActiveGameByRoomIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getActiveGameByRoomIdStmt: %w", cerr)
		}
	}
	if q.getGameStmt != nil {
		if cerr := q.getGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGameStmt: %w", cerr)
		}
	}
	if q.getGameByRoomIdStmt != nil {
		if cerr := q.getGameByRoomIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGameByRoomIdStmt: %w", cerr)
		}
	}
	if q.getGameRoomWithPlayersStmt != nil {
		if cerr := q.getGameRoomWithPlayersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGameRoomWithPlayersStmt: %w", cerr)
		}
	}
	if q.getPlayerStmt != nil {
		if cerr := q.getPlayerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPlayerStmt: %w", cerr)
		}
	}
	if q.startGameStmt != nil {
		if cerr := q.startGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing startGameStmt: %w", cerr)
		}
	}
	if q.updateGameStmt != nil {
		if cerr := q.updateGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateGameStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                         DBTX
	tx                         *sql.Tx
	addGameRoomPlayerStmt      *sql.Stmt
	createGameStmt             *sql.Stmt
	createGameRoomStmt         *sql.Stmt
	createPlayerStmt           *sql.Stmt
	finishStmt                 *sql.Stmt
	getActiveGameByRoomIdStmt  *sql.Stmt
	getGameStmt                *sql.Stmt
	getGameByRoomIdStmt        *sql.Stmt
	getGameRoomWithPlayersStmt *sql.Stmt
	getPlayerStmt              *sql.Stmt
	startGameStmt              *sql.Stmt
	updateGameStmt             *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                         tx,
		tx:                         tx,
		addGameRoomPlayerStmt:      q.addGameRoomPlayerStmt,
		createGameStmt:             q.createGameStmt,
		createGameRoomStmt:         q.createGameRoomStmt,
		createPlayerStmt:           q.createPlayerStmt,
		finishStmt:                 q.finishStmt,
		getActiveGameByRoomIdStmt:  q.getActiveGameByRoomIdStmt,
		getGameStmt:                q.getGameStmt,
		getGameByRoomIdStmt:        q.getGameByRoomIdStmt,
		getGameRoomWithPlayersStmt: q.getGameRoomWithPlayersStmt,
		getPlayerStmt:              q.getPlayerStmt,
		startGameStmt:              q.startGameStmt,
		updateGameStmt:             q.updateGameStmt,
	}
}
