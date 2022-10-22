package mongoqueryengine

import (
	"fmt"
	"testing"

	"slashbase.com/backend/internal/models"
)

func TestEngineConnection(t *testing.T) {
	mqueryengine := InitMongoQueryEngine()
	ping := mqueryengine.TestConnection(nil, &models.DBConnection{
		Type:           models.DBTYPE_MONGO,
		UseSSH:         models.DBUSESSH_NONE,
		DBName:         "testdb",
		DBHost:         "localhost",
		DBPort:         "27888",
		ConnectionUser: &models.DBConnectionUser{},
	})
	if !ping {
		t.Errorf("ping failed")
	} else {
		fmt.Println("pong")
	}
}
