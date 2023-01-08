package database

import (
	"database/sql"
	"github.com/crew_0/poker/internal/config"
	"github.com/crew_0/poker/internal/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	c, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

func randomPlayerNameForTest() string {
	return utils.RandomString(int(utils.RandomInt(5, 20)))
}
