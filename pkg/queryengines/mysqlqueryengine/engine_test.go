package mysqlqueryengine

// ----------------
// TEMP UNIT TESTS
// ----------------

// var dbConn *models.DBConnection = &models.DBConnection{
// 	Type:       models.DBTYPE_MYSQL,
// 	UseSSH:     models.DBUSESSH_NONE,
// 	DBName:     "",
// 	DBHost:     "localhost",
// 	DBPort:     "3306",
// 	DBScheme:   "mysql",
// 	DBUser:     "",
// 	DBPassword: "",
// }

// func TestRunQueryMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.RunQuery(dbConn, "select * from Persons;", &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestConnectionMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	ping := mqueryengine.TestConnection(dbConn, &models.QueryConfig{})
// 	if !ping {
// 		t.Errorf("error")
// 	} else {
// 		fmt.Println("pong")
// 	}
// }

// func TestGetDataModelsMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.GetDataModels(dbConn, &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestGetSingleDataModelMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.GetSingleDataModelFields(dbConn, "Persons", &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestGetIndexesMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.GetSingleDataModelIndexes(dbConn, "Persons", &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestGetDataMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.GetData(dbConn, "temp", 20, 0, true, []string{}, []string{}, &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestUpdateSingleDataMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.UpdateSingleData(dbConn, "Persons", "", "", "", &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestAddDataMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.AddData(dbConn, "temp", map[string]interface{}{"c1": "test1", "c2": "text1"}, &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestDeleteDataMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.DeleteData(dbConn, "Persons", []string{}, &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestAddSingleDataModelIndexMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.AddSingleDataModelIndex(dbConn, "Persons", "idx_name", []string{"FirstName", "LastName"}, true, &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }

// func TestDeleteSingleDataModelIndexMySQLEngineConnection(t *testing.T) {
// 	mqueryengine := InitMysqlQueryEngine()
// 	data, err := mqueryengine.DeleteSingleDataModelIndex(dbConn, "Persons", "idx_name", &models.QueryConfig{})
// 	if err != nil {
// 		t.Errorf("error: %s", err.Error())
// 	} else {
// 		fmt.Println("data: ", data)
// 	}
// }
