package mongoqueryengine

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/slashbaseide/slashbase/pkg/queryengines/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClientInstance struct {
	mongoClientInstance *mongo.Client
	LastUsed            time.Time
}

func createMongoConnectionURI(scheme string, host string, port uint16, user, password string, useSSL bool) string {
	usernamePassword := ""
	if user != "" && password != "" {
		usernamePassword = user + ":" + password + "@"
	}
	if scheme == "mongodb" {
		// Adding support to connect to Azure CosmosDB using MongoDB API.
		// According to official docs, the connection string should pass
		// ssl=true param to connect.
		if useSSL {
			return "mongodb://" + usernamePassword + host + ":" + strconv.Itoa(int(port)) + "/?ssl=true"
		} else {
			return "mongodb://" + usernamePassword + host + ":" + strconv.Itoa(int(port))
		}
	} else if scheme == "mongodb+srv" {
		return "mongodb+srv://" + usernamePassword + host
	}
	return ""
}

func (mEngine *MongoQueryEngine) getConnection(dbConnectionId, scheme, host string, port uint16, user, password string, useSSL bool) (c *mongo.Client, err error) {
	if mClientInstance, exists := mEngine.openClients[dbConnectionId]; exists {
		mEngine.mutex.Lock()
		mEngine.openClients[dbConnectionId] = mongoClientInstance{
			mongoClientInstance: mClientInstance.mongoClientInstance,
			LastUsed:            time.Now(),
		}
		mEngine.mutex.Unlock()
		return mClientInstance.mongoClientInstance, nil
	}
	err = utils.CheckTcpConnection(host, strconv.Itoa(int(port)))
	if err != nil {
		return
	}
	connectionURI := createMongoConnectionURI(scheme, host, port, user, password, useSSL)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionURI))
	if err != nil {
		err = fmt.Errorf("unable to connect to database: %v", err)
		return
	}
	if dbConnectionId != "" {
		mEngine.mutex.Lock()
		mEngine.openClients[dbConnectionId] = mongoClientInstance{
			mongoClientInstance: client,
			LastUsed:            time.Now(),
		}
		mEngine.mutex.Unlock()
	}
	return client, err
}

func (mEngine *MongoQueryEngine) RemoveUnusedConnections() {
	for dbConnID, instance := range mEngine.openClients {
		now := time.Now()
		diff := now.Sub(instance.LastUsed)
		if diff.Minutes() > 20 {
			mEngine.mutex.Lock()
			delete(mEngine.openClients, dbConnID)
			mEngine.mutex.Unlock()
			go instance.mongoClientInstance.Disconnect(context.Background())
		}
	}
}
