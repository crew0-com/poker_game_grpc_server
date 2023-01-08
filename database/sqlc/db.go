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
	if q.addGameMessageStmt, err = db.PrepareContext(ctx, addGameMessage); err != nil {
		return nil, fmt.Errorf("error preparing query AddGameMessage: %w", err)
	}
	if q.addPlayerToGameRoomStmt, err = db.PrepareContext(ctx, addPlayerToGameRoom); err != nil {
		return nil, fmt.Errorf("error preparing query AddPlayerToGameRoom: %w", err)
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
	if q.finishGameStmt, err = db.PrepareContext(ctx, finishGame); err != nil {
		return nil, fmt.Errorf("error preparing query FinishGame: %w", err)
	}
	if q.getActiveGameStmt, err = db.PrepareContext(ctx, getActiveGame); err != nil {
		return nil, fmt.Errorf("error preparing query GetActiveGame: %w", err)
	}
	if q.getGameStmt, err = db.PrepareContext(ctx, getGame); err != nil {
		return nil, fmt.Errorf("error preparing query GetGame: %w", err)
	}
	if q.getGameMessagesStmt, err = db.PrepareContext(ctx, getGameMessages); err != nil {
		return nil, fmt.Errorf("error preparing query GetGameMessages: %w", err)
	}
	if q.getGameRoomAndPlayerRowsStmt, err = db.PrepareContext(ctx, getGameRoomAndPlayerRows); err != nil {
		return nil, fmt.Errorf("error preparing query GetGameRoomAndPlayerRows: %w", err)
	}
	if q.getGamesByRoomIdStmt, err = db.PrepareContext(ctx, getGamesByRoomId); err != nil {
		return nil, fmt.Errorf("error preparing query GetGamesByRoomId: %w", err)
	}
	if q.getPlayerStmt, err = db.PrepareContext(ctx, getPlayer); err != nil {
		return nil, fmt.Errorf("error preparing query GetPlayer: %w", err)
	}
	if q.paginatedGameMessageByGameRoomStmt, err = db.PrepareContext(ctx, paginatedGameMessageByGameRoom); err != nil {
		return nil, fmt.Errorf("error preparing query PaginatedGameMessageByGameRoom: %w", err)
	}
	if q.paginatedGameMessagesStmt, err = db.PrepareContext(ctx, paginatedGameMessages); err != nil {
		return nil, fmt.Errorf("error preparing query PaginatedGameMessages: %w", err)
	}
	if q.setActiveGameStmt, err = db.PrepareContext(ctx, setActiveGame); err != nil {
		return nil, fmt.Errorf("error preparing query SetActiveGame: %w", err)
	}
	if q.startGameStmt, err = db.PrepareContext(ctx, startGame); err != nil {
		return nil, fmt.Errorf("error preparing query StartGame: %w", err)
	}
	if q.unsetActiveGameStmt, err = db.PrepareContext(ctx, unsetActiveGame); err != nil {
		return nil, fmt.Errorf("error preparing query UnsetActiveGame: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addGameMessageStmt != nil {
		if cerr := q.addGameMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addGameMessageStmt: %w", cerr)
		}
	}
	if q.addPlayerToGameRoomStmt != nil {
		if cerr := q.addPlayerToGameRoomStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addPlayerToGameRoomStmt: %w", cerr)
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
	if q.finishGameStmt != nil {
		if cerr := q.finishGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing finishGameStmt: %w", cerr)
		}
	}
	if q.getActiveGameStmt != nil {
		if cerr := q.getActiveGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getActiveGameStmt: %w", cerr)
		}
	}
	if q.getGameStmt != nil {
		if cerr := q.getGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGameStmt: %w", cerr)
		}
	}
	if q.getGameMessagesStmt != nil {
		if cerr := q.getGameMessagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGameMessagesStmt: %w", cerr)
		}
	}
	if q.getGameRoomAndPlayerRowsStmt != nil {
		if cerr := q.getGameRoomAndPlayerRowsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGameRoomAndPlayerRowsStmt: %w", cerr)
		}
	}
	if q.getGamesByRoomIdStmt != nil {
		if cerr := q.getGamesByRoomIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGamesByRoomIdStmt: %w", cerr)
		}
	}
	if q.getPlayerStmt != nil {
		if cerr := q.getPlayerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPlayerStmt: %w", cerr)
		}
	}
	if q.paginatedGameMessageByGameRoomStmt != nil {
		if cerr := q.paginatedGameMessageByGameRoomStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing paginatedGameMessageByGameRoomStmt: %w", cerr)
		}
	}
	if q.paginatedGameMessagesStmt != nil {
		if cerr := q.paginatedGameMessagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing paginatedGameMessagesStmt: %w", cerr)
		}
	}
	if q.setActiveGameStmt != nil {
		if cerr := q.setActiveGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setActiveGameStmt: %w", cerr)
		}
	}
	if q.startGameStmt != nil {
		if cerr := q.startGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing startGameStmt: %w", cerr)
		}
	}
	if q.unsetActiveGameStmt != nil {
		if cerr := q.unsetActiveGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unsetActiveGameStmt: %w", cerr)
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
	db                                 DBTX
	tx                                 *sql.Tx
	addGameMessageStmt                 *sql.Stmt
	addPlayerToGameRoomStmt            *sql.Stmt
	createGameStmt                     *sql.Stmt
	createGameRoomStmt                 *sql.Stmt
	createPlayerStmt                   *sql.Stmt
	finishGameStmt                     *sql.Stmt
	getActiveGameStmt                  *sql.Stmt
	getGameStmt                        *sql.Stmt
	getGameMessagesStmt                *sql.Stmt
	getGameRoomAndPlayerRowsStmt       *sql.Stmt
	getGamesByRoomIdStmt               *sql.Stmt
	getPlayerStmt                      *sql.Stmt
	paginatedGameMessageByGameRoomStmt *sql.Stmt
	paginatedGameMessagesStmt          *sql.Stmt
	setActiveGameStmt                  *sql.Stmt
	startGameStmt                      *sql.Stmt
	unsetActiveGameStmt                *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                 tx,
		tx:                                 tx,
		addGameMessageStmt:                 q.addGameMessageStmt,
		addPlayerToGameRoomStmt:            q.addPlayerToGameRoomStmt,
		createGameStmt:                     q.createGameStmt,
		createGameRoomStmt:                 q.createGameRoomStmt,
		createPlayerStmt:                   q.createPlayerStmt,
		finishGameStmt:                     q.finishGameStmt,
		getActiveGameStmt:                  q.getActiveGameStmt,
		getGameStmt:                        q.getGameStmt,
		getGameMessagesStmt:                q.getGameMessagesStmt,
		getGameRoomAndPlayerRowsStmt:       q.getGameRoomAndPlayerRowsStmt,
		getGamesByRoomIdStmt:               q.getGamesByRoomIdStmt,
		getPlayerStmt:                      q.getPlayerStmt,
		paginatedGameMessageByGameRoomStmt: q.paginatedGameMessageByGameRoomStmt,
		paginatedGameMessagesStmt:          q.paginatedGameMessagesStmt,
		setActiveGameStmt:                  q.setActiveGameStmt,
		startGameStmt:                      q.startGameStmt,
		unsetActiveGameStmt:                q.unsetActiveGameStmt,
	}
}
