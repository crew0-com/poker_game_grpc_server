package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomGameRoomTest(t *testing.T) GameRoom {
	player := CreateRandomPlayerTest(t)

	gameRoom, err := testQueries.CreateGameRoom(context.Background(), player.PlayerID)
	require.NoError(t, err)
	require.NotEmpty(t, gameRoom)

	require.Equal(t, player.PlayerID, gameRoom.CreatedBy)

	return gameRoom
}

func TestQueries_CreateGameRoom(t *testing.T) {
	t.Run("should create a new game room", func(t *testing.T) {
		CreateRandomGameRoomTest(t)
	})

	t.Run("should return error when player not found", func(t *testing.T) {
		_, err := testQueries.CreateGameRoom(context.Background(), uuid.New())

		if assert.Error(t, err) {
			if pqErr, ok := err.(*pq.Error); ok {
				assert.Equal(t, pq.ErrorCode("23503"), pqErr.Code)
			} else {
				t.Errorf("unexpected error type: %T", err)
			}
		}
	})
}

func TestQueries_AddGameRoomPlayer(t *testing.T) {
	t.Run("should add many players", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		players := CreateRandomPlayersTest(t, 5)

		for _, player := range players {
			_, err := testQueries.AddPlayerToGameRoom(context.Background(), AddPlayerToGameRoomParams{
				GameRoomID: gameRoom.GameRoomID,
				PlayerID:   player.PlayerID,
			})

			require.NoError(t, err)
		}

		rows, err := testQueries.GetGameRoomAndPlayerRows(context.Background(), gameRoom.GameRoomID)
		require.NoError(t, err)
		require.NotEmpty(t, rows)
		require.Len(t, rows, len(players))

		for i, row := range rows {
			require.Equal(t, gameRoom.GameRoomID, row.GameRoomID)
			require.Equal(t, players[i].PlayerID, row.PlayerID)
		}
	})

	t.Run("should return error when player to be added is not found", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		_, err := testQueries.AddPlayerToGameRoom(context.Background(), AddPlayerToGameRoomParams{
			GameRoomID: gameRoom.GameRoomID,
			PlayerID:   uuid.New(),
		})

		if assert.Error(t, err) {
			if pqErr, ok := err.(*pq.Error); ok {
				assert.Equal(t, pq.ErrorCode("23503"), pqErr.Code)
			} else {
				t.Errorf("unexpected error type: %T", err)
			}
		}
	})
}
