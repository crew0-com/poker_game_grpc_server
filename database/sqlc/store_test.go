package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSQLStore_GetOrCreatePlayer(t *testing.T) {
	store := NewStore(testDB)

	t.Run("Should get a existing player", func(t *testing.T) {
		existingPlayer := CreateRandomPlayerTest(t)

		param := GetOrCreatePlayerParams{
			Playername: existingPlayer.Name,
			PlayerId:   existingPlayer.PlayerID,
		}
		player, err := store.GetOrCreatePlayer(context.Background(), param)
		require.NoError(t, err)
		require.NotEmpty(t, player)
		require.Equal(t, existingPlayer.PlayerID, player.PlayerID)
		require.Equal(t, existingPlayer.Name, player.Name)
	})

	t.Run("Should create a new player", func(t *testing.T) {
		param := GetOrCreatePlayerParams{
			Playername: "New Player",
			PlayerId:   uuid.New(),
		}
		_, err := store.GetPlayer(context.Background(), param.PlayerId)
		require.Error(t, err)

		player, err := store.GetOrCreatePlayer(context.Background(), param)
		require.NoError(t, err)
		require.NotEmpty(t, player)
		require.Equal(t, param.PlayerId, player.PlayerID)
		require.Equal(t, param.Playername, player.Name)
	})
}

func TestSQLStore_GetGameRoomWithPlayers(t *testing.T) {
	store := NewStore(testDB)

	t.Run("Should get a game room with players", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		players := CreateRandomPlayersTest(t, 3)

		for _, player := range players {
			_, err := store.AddPlayerToGameRoom(context.Background(), AddPlayerToGameRoomParams{
				GameRoomID: gameRoom.GameRoomID,
				PlayerID:   player.PlayerID,
			})
			require.NoError(t, err)
		}

		gameRoomWithPlayers, err := store.GetGameRoomWithPlayers(context.Background(), gameRoom.GameRoomID)
		require.NoError(t, err)
		require.NotEmpty(t, gameRoomWithPlayers)
		require.Equal(t, gameRoom.GameRoomID, gameRoomWithPlayers.GameRoomID)
		require.Equal(t, len(players), len(gameRoomWithPlayers.Players))

		for i, row := range gameRoomWithPlayers.Players {
			require.Equal(t, players[i].PlayerID, row.PlayerID)
		}
	})

	t.Run("Should return error if gameroom is not found ", func(t *testing.T) {
		_, err := store.GetGameRoomWithPlayers(context.Background(), uuid.New())
		require.Error(t, err)
	})

	t.Run("Should return error if gameroom has no players ", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		_, err := store.GetGameRoomWithPlayers(context.Background(), gameRoom.GameRoomID)
		require.Error(t, err)
	})
}
