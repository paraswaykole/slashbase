package mongoqueryengine

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClientInstance struct {
	mongoClientInstance *mongo.Client
	LastUsed            time.Time
}

func createMongoConnectionURI(scheme string, host string, port uint16, user, password string) string {
	usernamePassword := ""
	if user != "" && password != "" {
		usernamePassword = user + ":" + password + "@"
	}
	if scheme == "mongodb" {
		return "mongodb://" + usernamePassword + host + ":" + strconv.Itoa(int(port))
	} else if scheme == "mongodb+srv" {
		return "mongodb+srv://" + usernamePassword + host
	}
	return ""
}

func (mEngine *MongoQueryEngine) getConnection(dbConnectionId, scheme, host string, port uint16, user, password string) (c *mongo.Client, err error) {
	if mClientInstance, exists := mEngine.openClients[dbConnectionId]; exists {
		mEngine.openClients[dbConnectionId] = mongoClientInstance{
			mongoClientInstance: mClientInstance.mongoClientInstance,
			LastUsed:            time.Now(),
		}
		return mClientInstance.mongoClientInstance, nil
	}
	connectionURI := createMongoConnectionURI(scheme, host, port, user, password)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionURI))
	if err != nil {
		err = fmt.Errorf("unable to connect to database: %v", err)
		return
	}
	if dbConnectionId != "" {
		mEngine.openClients[dbConnectionId] = mongoClientInstance{
			mongoClientInstance: client,
			LastUsed:            time.Now(),
		}
	}
	return client, err
}

func (mEngine *MongoQueryEngine) RemoveUnusedConnections() {
	for dbConnID, instance := range mEngine.openClients {
		now := time.Now()
		diff := now.Sub(instance.LastUsed)
		if diff.Minutes() > 20 {
			delete(mEngine.openClients, dbConnID)
			go func() {
				err := instance.mongoClientInstance.Disconnect(context.Background())
				if err != nil {
					return
				}
			}()
		}
	}
}
