package libs

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	mongoClientOpts *options.ClientOptions
	defaultDB       string
)

func InitMongoDB(hosts []string, user, password string, database string) {
	log.Info("初始化Mongo配置")
	defaultDB = database
	opts := options.Client()
	opts.SetHosts(hosts)
	opts.SetAuth(options.Credential{
		Username: user,
		Password: password,
	})
	opts.SetMaxPoolSize(20)
	mongoClientOpts = opts
}

func GetContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

// GetMongoDB 使用 mongo-driver 获取示例
func GetMongoDB() (db *mongo.Database, err error) {
	opts := mongoClientOpts
	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatal("请检查mongodb配置", err)
	}

	ctx, cancel := GetContext()
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("连接mongodb出错", err)
		return
	}
	db = client.Database(defaultDB)
	return
}

func MongoExec(f func(ctx context.Context, db *mongo.Database) (err error)) (err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("mongo panic: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	db, err := GetMongoDB()
	if err != nil {
		return
	}
	ctx, cancel := GetContext()
	defer cancel()
	defer db.Client().Disconnect(ctx)

	// exec
	err = f(ctx, db)
	return
}

func MongoFindOne(collectionName string, filter interface{}, opts ...*options.FindOneOptions) (result *mongo.SingleResult, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("mongo panic: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	if collectionName == "" {
		err = errors.New("未指定collectionName")
		return
	}

	db, err := GetMongoDB()
	if err != nil {
		return
	}
	ctx, cancel := GetContext()
	defer cancel()
	defer db.Client().Disconnect(ctx)

	// findOne
	coll := db.Collection(collectionName)
	result = coll.FindOne(ctx, filter, opts...)
	return
}

func MongoInsertOne(collectionName string, doc interface{}, opts ...*options.InsertOneOptions) (result *mongo.InsertOneResult, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("mongo panic: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	if collectionName == "" {
		err = errors.New("未指定collectionName")
		return
	}

	db, err := GetMongoDB()
	if err != nil {
		return
	}
	ctx, cancel := GetContext()
	defer cancel()
	defer db.Client().Disconnect(ctx)

	// insert
	coll := db.Collection(collectionName)
	result, err = coll.InsertOne(ctx, doc, opts...)
	return
}

func MongoUpdateOne(collectionName string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("mongo panic: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	if collectionName == "" {
		err = errors.New("未指定collectionName")
		return
	}

	db, err := GetMongoDB()
	if err != nil {
		return
	}
	ctx, cancel := GetContext()
	defer cancel()
	defer db.Client().Disconnect(ctx)

	// update
	coll := db.Collection(collectionName)
	result, err = coll.UpdateOne(ctx, filter, update, opts...)
	return
}
