package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/crew_0/poker/internal/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomMessageInGame(t *testing.T, game Game) GameMessage {
	arg := AddGameMessageParams{
		GameID:     game.GameID,
		GameRoomID: game.GameRoomID,
		Message:    json.RawMessage(fmt.Sprintf("{\"content\": \"%s\"}", utils.RandomString(10))),
	}

	message, err := testQueries.AddGameMessage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, message)

	require.Equal(t, arg.Message, message.Message)

	require.NotZero(t, message.MessageID)
	require.NotZero(t, message.CreatedAt)
	require.WithinDuration(t, message.CreatedAt, time.Now(), time.Second)

	return message
}

func CreateRandomMessagesInGame(t *testing.T, n int) ([]GameMessage, Game) {
	game := CreateRandomGameTest(t)
	var messages []GameMessage
	for i := 0; i < n; i++ {
		messages = append(messages, CreateRandomMessageInGame(t, game))
	}
	return messages, game
}

func CreateRandomMessage(t *testing.T) (GameMessage, Game) {
	game := CreateRandomGameTest(t)

	return CreateRandomMessageInGame(t, game), game
}

func TestQueries_AddGameMessage(t *testing.T) {
	CreateRandomMessage(t)
}

func TestQueries_GetGameMessages(t *testing.T) {
	t.Run("get empty set if no messages in game room", func(t *testing.T) {
		game := CreateRandomGameTest(t)
		messages, err := testQueries.GetGameMessages(context.Background(), game.GameID)
		require.NoError(t, err)
		require.Empty(t, messages)
	})

	t.Run("get game messages in desc order", func(t *testing.T) {
		count := 5
		messages, game := CreateRandomMessagesInGame(t, count)

		gotMessages, err := testQueries.GetGameMessages(context.Background(), game.GameID)
		require.NoError(t, err)
		require.NotEmpty(t, gotMessages)
		require.Equal(t, len(gotMessages), len(messages))

		for i := 0; i < count; i++ {
			require.Equal(t, messages[count-(i+1)].MessageID, gotMessages[i].MessageID)
		}
	})

	t.Run("get latest 100 messages in desc order", func(t *testing.T) {
		count := 105
		messages, game := CreateRandomMessagesInGame(t, count)

		gotMessages, err := testQueries.GetGameMessages(context.Background(), game.GameID)
		require.NoError(t, err)
		require.NotEmpty(t, gotMessages)
		require.Equal(t, 100, len(gotMessages))

		for i, hardLimit := 0, 100; i < hardLimit; i++ {
			require.Equal(t, messages[count-(i+1)].MessageID, gotMessages[i].MessageID)
		}
	})
}

func TestQueries_GetGameMessagesByGameRoom(t *testing.T) {
	t.Run("get empty set if no messages in game room", func(t *testing.T) {
		gameRoom := CreateRandomGameRoomTest(t)

		params := PaginatedGameMessageByGameRoomParams{
			GameRoomID: gameRoom.GameRoomID,
			Limit:      10,
			Offset:     0,
		}

		messages, err := testQueries.PaginatedGameMessageByGameRoom(context.Background(), params)
		require.NoError(t, err)
		require.Empty(t, messages)
	})

	t.Run("get paginated game messages in desc order", func(t *testing.T) {
		count := 30
		messages, game := CreateRandomMessagesInGame(t, count)

		limit := 10
		params := PaginatedGameMessageByGameRoomParams{
			GameRoomID: game.GameRoomID,
			Limit:      int32(limit),
			Offset:     0,
		}
		gotMessages, err := testQueries.PaginatedGameMessageByGameRoom(context.Background(), params)
		require.NoError(t, err)
		require.NotEmpty(t, gotMessages)
		require.Equal(t, len(gotMessages), limit)

		for i := 0; i < limit; i++ {
			require.Equal(t, messages[count-(i+1)].MessageID, gotMessages[i].MessageID)
		}
	})

	t.Run("get different offset game messages in desc order", func(t *testing.T) {
		count := 30
		messages, game := CreateRandomMessagesInGame(t, count)

		limit := 10
		offset := 10
		params := PaginatedGameMessageByGameRoomParams{
			GameRoomID: game.GameRoomID,
			Limit:      int32(limit),
			Offset:     int32(offset),
		}
		gotMessages, err := testQueries.PaginatedGameMessageByGameRoom(context.Background(), params)
		require.NoError(t, err)
		require.NotEmpty(t, gotMessages)
		require.Equal(t, len(gotMessages), limit)

		for i := 0; i < limit; i++ {
			require.Equal(t, messages[count-(i+offset+1)].MessageID, gotMessages[i].MessageID)
		}
	})
}

func TestQueries_PaginatedGameMessages(t *testing.T) {
	t.Run("get empty set if no messages in game room", func(t *testing.T) {
		game := CreateRandomGameTest(t)

		params := PaginatedGameMessagesParams{
			GameID: game.GameID,
			Limit:  10,
			Offset: 0,
		}

		messages, err := testQueries.PaginatedGameMessages(context.Background(), params)
		require.NoError(t, err)
		require.Empty(t, messages)
	})

	t.Run("get paginated game messages in desc order", func(t *testing.T) {
		count := 30
		messages, game := CreateRandomMessagesInGame(t, count)

		limit := 10
		params := PaginatedGameMessagesParams{
			GameID: game.GameID,
			Limit:  int32(limit),
			Offset: 0,
		}
		gotMessages, err := testQueries.PaginatedGameMessages(context.Background(), params)
		require.NoError(t, err)
		require.NotEmpty(t, gotMessages)
		require.Equal(t, len(gotMessages), limit)

		for i := 0; i < limit; i++ {
			require.Equal(t, messages[count-(i+1)].MessageID, gotMessages[i].MessageID)
		}
	})

	t.Run("get different offset game messages in desc order", func(t *testing.T) {
		count := 30
		messages, game := CreateRandomMessagesInGame(t, count)

		limit := 10
		offset := 10
		params := PaginatedGameMessagesParams{
			GameID: game.GameID,
			Limit:  int32(limit),
			Offset: int32(offset),
		}
		gotMessages, err := testQueries.PaginatedGameMessages(context.Background(), params)
		require.NoError(t, err)
		require.NotEmpty(t, gotMessages)
		require.Equal(t, len(gotMessages), limit)

		for i := 0; i < limit; i++ {
			require.Equal(t, messages[count-(i+offset+1)].MessageID, gotMessages[i].MessageID)
		}
	})
}
