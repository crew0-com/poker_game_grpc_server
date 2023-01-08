package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/crew_0/poker/internal/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomGamesTest(t *testing.T, gameRoomId uuid.UUID, n int) []Game {
	var games []Game
	for i := 0; i < n; i++ {
		games = append(games, CreateGameInRoomTest(t, gameRoomId))
	}
	return games
}

func CreateGameInRoomTest(t *testing.T, gameRoomId uuid.UUID) Game {
	gameStateJson := map[string]string{"deck": utils.RandomString(8), "players": utils.RandomString(8), "turn": utils.RandomString(8)}
	gameMessageJson := map[string]string{"message": utils.RandomString(8)}
	gameStateJsonBytes, _ := json.Marshal(gameStateJson)
	gameMessageJsonBytes, _ := json.Marshal(gameMessageJson)

	rawGameMessageJSON := json.RawMessage(gameMessageJsonBytes)
	rawGameStateJSON := json.RawMessage(gameStateJsonBytes)
	arg := CreateGameParams{
		GameState:  rawGameStateJSON,
		Messages:   rawGameMessageJSON,
		GameRoomID: gameRoomId,
	}
	game, err := testQueries.CreateGame(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, game)

	require.False(t, game.HasStarted)
	require.False(t, game.HasFinished)
	require.Equal(t, game.GameRoomID, gameRoomId)

	var gotGameState map[string]string
	err = json.Unmarshal(game.GameState, &gotGameState)
	require.NoError(t, err)
	require.Equal(t, gotGameState, gameStateJson)

	var gotGameMessage map[string]string
	err = json.Unmarshal(game.Messages, &gotGameMessage)
	require.NoError(t, err)
	require.Equal(t, gotGameMessage, gameMessageJson)

	return game
}

func CreateRandomGameTest(t *testing.T) Game {
	gameRoom := CreateRandomGameRoomTest(t)

	return CreateGameInRoomTest(t, gameRoom.GameRoomID)
}

func TestCreateGame(t *testing.T) {
	CreateRandomGameTest(t)
}

func TestQueries_FinishGame(t *testing.T) {
	game := CreateRandomGameTest(t)

	gotGame, err := testQueries.FinishGame(context.Background(), game.GameID)
	require.NoError(t, err)
	require.NotEmpty(t, gotGame)
	require.True(t, gotGame.HasFinished)
	require.True(t, gotGame.FinishedAt.Valid)
	require.WithinDuration(t, time.Now(), gotGame.FinishedAt.Time, 100*time.Millisecond)
}

func TestQueries_GetGame(t *testing.T) {
	game := CreateRandomGameTest(t)

	gotGame, err := testQueries.GetGame(context.Background(), game.GameID)
	require.NoError(t, err)
	require.NotEmpty(t, gotGame)
	require.Equal(t, gotGame.GameID, game.GameID)
	require.Equal(t, gotGame.GameRoomID, game.GameRoomID)
	require.Equal(t, gotGame.HasFinished, game.HasFinished)
	require.Equal(t, gotGame.HasStarted, game.HasStarted)
	require.Equal(t, gotGame.StartedAt, game.StartedAt)
	require.Equal(t, gotGame.FinishedAt, game.FinishedAt)
}

func TestQueries_GetGamesByGameRoomID(t *testing.T) {
	gameRoom := CreateRandomGameRoomTest(t)
	n := 10

	games := CreateRandomGamesTest(t, gameRoom.GameRoomID, n)

	gotGames, err := testQueries.GetGamesByRoomId(context.Background(), gameRoom.GameRoomID)
	require.NoError(t, err)
	require.NotEmpty(t, gotGames)
	require.Len(t, gotGames, n)

	for i, gotGame := range gotGames {
		require.Equal(t, games[i], gotGame)
	}
}

func TestQueries_StartGame(t *testing.T) {
	game := CreateRandomGameTest(t)

	gotGame, err := testQueries.StartGame(context.Background(), game.GameID)
	require.NoError(t, err)
	require.NotEmpty(t, gotGame)
	require.True(t, gotGame.HasStarted)
	require.True(t, gotGame.StartedAt.Valid)
	require.WithinDuration(t, time.Now(), gotGame.StartedAt.Time, 100*time.Millisecond)
}

func TestQueries_UpdateGame(t *testing.T) {
	game := CreateRandomGameTest(t)

	gameStateJson := map[string]string{"deck": utils.RandomString(8), "players": utils.RandomString(8), "turn": utils.RandomString(8)}
	gameMessageJson := map[string]string{"message": utils.RandomString(8)}
	gameStateJsonBytes, _ := json.Marshal(gameStateJson)
	gameMessageJsonBytes, _ := json.Marshal(gameMessageJson)

	rawGameMessageJSON := json.RawMessage(gameMessageJsonBytes)
	rawGameStateJSON := json.RawMessage(gameStateJsonBytes)

	arg := UpdateGameParams{
		GameID:    game.GameID,
		GameState: rawGameStateJSON,
		Messages:  rawGameMessageJSON,
	}

	gotGame, err := testQueries.UpdateGame(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, gotGame)
	require.Equal(t, game.GameID, gotGame.GameID)
	require.Equal(t, game.GameRoomID, gotGame.GameRoomID)

	var gotGameState map[string]string
	err = json.Unmarshal(gotGame.GameState, &gotGameState)
	require.NoError(t, err)
	require.Equal(t, gotGameState, gameStateJson)

	var gotGameMessage map[string]string
	err = json.Unmarshal(gotGame.Messages, &gotGameMessage)
	require.NoError(t, err)
	require.Equal(t, gotGameMessage, gameMessageJson)
}

func TestQueries_SetActiveGame(t *testing.T) {
	t.Run("set active game", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		game := CreateGameInRoomTest(t, gameRoom.GameRoomID)

		retVal, err := testQueries.SetActiveGame(context.Background(), game.GameID)

		require.NoError(t, err)
		require.NotEmpty(t, retVal)
		require.Equal(t, game.GameID, retVal.GameID)
		require.Equal(t, game.GameRoomID, retVal.GameRoomID)
	})

	t.Run("return error if game is set as active for a room that has an active game", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		game := CreateGameInRoomTest(t, gameRoom.GameRoomID)
		_, err := testQueries.SetActiveGame(context.Background(), game.GameID)
		require.NoError(t, err)

		anotherGame := CreateGameInRoomTest(t, gameRoom.GameRoomID)
		_, err = testQueries.SetActiveGame(context.Background(), anotherGame.GameID)

		require.Error(t, err)
	})
}

func TestQueries_GetActiveGame(t *testing.T) {
	t.Run("get active game", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		game := CreateGameInRoomTest(t, gameRoom.GameRoomID)
		_, err := testQueries.SetActiveGame(context.Background(), game.GameID)
		require.NoError(t, err)

		retVal, err := testQueries.GetActiveGame(context.Background(), gameRoom.GameRoomID)

		require.NoError(t, err)
		require.NotEmpty(t, retVal)
		require.Equal(t, game, retVal)
	})

	t.Run("return error if no active game is found", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		_, err := testQueries.GetActiveGame(context.Background(), gameRoom.GameRoomID)

		require.Error(t, err)
		require.EqualError(t, sql.ErrNoRows, err.Error())
	})
}

func TestQueries_UnsetActiveGame(t *testing.T) {
	t.Run("unset active game", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		game := CreateGameInRoomTest(t, gameRoom.GameRoomID)
		_, err := testQueries.SetActiveGame(context.Background(), game.GameID)
		require.NoError(t, err)

		_, err = testQueries.UnsetActiveGame(context.Background(), game.GameID)
		require.NoError(t, err)

		_, err = testQueries.GetActiveGame(context.Background(), gameRoom.GameRoomID)
		require.Error(t, err)
		require.EqualError(t, sql.ErrNoRows, err.Error())
	})

	t.Run("return error if no active game is found", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)
		_, err := testQueries.UnsetActiveGame(context.Background(), gameRoom.GameRoomID)

		require.Error(t, err)
		require.EqualError(t, sql.ErrNoRows, err.Error())
	})
}
