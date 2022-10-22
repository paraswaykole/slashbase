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

func (mEngine *MongoQueryEngine) getConnection(dbConnectionId, host string, port uint16, database, user, password string) (c *mongo.Client, err error) {
	if mClientInstance, exists := mEngine.openClients[dbConnectionId]; exists {
		mEngine.openClients[dbConnectionId] = mongoClientInstance{
			mongoClientInstance: mClientInstance.mongoClientInstance,
			LastUsed:            time.Now(),
		}
		return mClientInstance.mongoClientInstance, nil
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://"+host+":"+strconv.Itoa(int(port))))
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
	for {
		time.Sleep(time.Minute * time.Duration(5))
		for dbConnID, instance := range mEngine.openClients {
			now := time.Now()
			diff := now.Sub(instance.LastUsed)
			if diff.Minutes() > 20 {
				delete(mEngine.openClients, dbConnID)
				go instance.mongoClientInstance.Disconnect(context.Background())
			}
		}
	}
}
