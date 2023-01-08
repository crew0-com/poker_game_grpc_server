package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomPlayerTest(t *testing.T) Player {
	name := randomPlayerNameForTest()

	arg := CreatePlayerParams{
		Name:     name,
		PlayerID: uuid.New(),
	}

	player, err := testQueries.CreatePlayer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, player)

	require.Equal(t, arg.Name, player.Name)
	require.Equal(t, arg.PlayerID, player.PlayerID)

	return player
}

func CreateRandomPlayersTest(t *testing.T, n int) []Player {
	if n <= 0 {
		return nil
	}

	players := make([]Player, n)

	for i := 0; i < n; i++ {
		players[i] = CreateRandomPlayerTest(t)
	}

	return players
}

func TestQueries_CreatePlayer(t *testing.T) {
	CreateRandomPlayerTest(t)
}

func TestQueries_GetPlayer(t *testing.T) {
	player1 := CreateRandomPlayerTest(t)

	player2, err := testQueries.GetPlayer(context.Background(), player1.PlayerID)
	require.NoError(t, err)
	require.NotEmpty(t, player2)

	require.Equal(t, player1.PlayerID, player2.PlayerID)
	require.Equal(t, player1.Name, player2.Name)
}
