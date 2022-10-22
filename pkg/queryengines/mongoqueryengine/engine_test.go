package mongoqueryengine

import (
	"fmt"
	"testing"

	"slashbase.com/backend/internal/models"
)

func TestEngineConnection(t *testing.T) {
	mqueryengine := InitMongoQueryEngine()
	data, err := mqueryengine.RunQuery(nil, &models.DBConnection{
		Type:           models.DBTYPE_MONGO,
		UseSSH:         models.DBUSESSH_NONE,
		DBName:         "testdb",
		DBHost:         "localhost",
		DBPort:         "27888",
		ConnectionUser: &models.DBConnectionUser{},
	}, "db.user.findOne()", false)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(data)
}
