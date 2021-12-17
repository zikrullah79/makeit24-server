package mongo

import (
	"context"
	"fmt"
	"os"
	"time"
	"zikrullah79/makeit24-server/pkg/mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient     *mongo.Client
	mongoCollection *mongo.Collection
)

func InitDB() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "hostKey", os.Getenv("MONGO_HOST"))
	ctx = context.WithValue(ctx, "port", os.Getenv("MONGO_PORT"))
	// ctx = context.WithValue(ctx, "usernameKey", os.Getenv("MONGO_USERNAME"))
	// ctx = context.WithValue(ctx, "passwordKey", os.Getenv("MONGO_PASSWORD"))
	// ctx = context.WithValue(ctx, "databaseKey", os.Getenv("MONGO_DATABASE"))

	uri := fmt.Sprintf(`mongodb://%s:%s`,
		ctx.Value("hostKey").(string),
		ctx.Value("port").(string),
		// ctx.Value("usernameKey").(string),
		// ctx.Value("passwordKey").(string),
		// ctx.Value("databaseKey").(string),
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return err
	}
	// err = client.Connect(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	db := client.Database("dev")
	// defer client.Disconnect(ctx)
	mongoClient = client
	mongoCollection = db.Collection("score")
	return nil
}
func GetAllScores() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	c, err := mongoCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var t []bson.M
	if err := c.All(ctx, &t); err != nil {
		fmt.Println(err)
	}
	return t, nil
}

func GetOneById(id string) (*model.Score, error) {
	c := mongoCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	s := &model.Score{}
	err := c.Decode(&s)

	if err != nil {
		return nil, err
	}
	return s, nil
}
func InsertScores(score *model.Score) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := mongoCollection.InsertOne(ctx, score)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}
