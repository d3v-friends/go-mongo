package fnMongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const CtxMongoClientKey = "CTX_MONGO_CLIENT"
const CtxMongoDbKey = "CTX_MONGO_DB"

var (
	ErrNotFoundMongoClient = errors.New("not found *mongo.Client")
)

func GetClient(ctx context.Context) (client *mongo.Client, err error) {
	var has bool
	if client, has = ctx.Value(CtxMongoClientKey).(*mongo.Client); !has {
		err = ErrNotFoundMongoClient
		return
	}
	return
}

func GetClientP(ctx context.Context) (client *mongo.Client) {
	var err error
	if client, err = GetClient(ctx); err != nil {
		panic(err)
	}
	return
}

func SetClient(ctx context.Context, client *mongo.Client) context.Context {
	return context.WithValue(ctx, CtxMongoClientKey, client)
}

func GetDB(ctx context.Context) (db *mongo.Database, err error) {
	var has bool
	if db, has = ctx.Value(CtxMongoDbKey).(*mongo.Database); !has {
		err = ErrNotFoundMongoClient
		return
	}

	return
}

func GetDBP(ctx context.Context) (db *mongo.Database) {
	var err error
	if db, err = GetDB(ctx); err != nil {
		panic(err)
	}
	return
}

func SetDB(ctx context.Context, db *mongo.Database) context.Context {
	return context.WithValue(ctx, CtxMongoDbKey, db)
}

func GetCol(ctx context.Context, colNm string) (col *mongo.Collection, err error) {
	var db *mongo.Database
	if db, err = GetDB(ctx); err != nil {
		return
	}
	col = db.Collection(colNm)
	return
}

func GetColP(ctx context.Context, colNm string) (col *mongo.Collection) {
	var err error
	if col, err = GetCol(ctx, colNm); err != nil {
		panic(err)
	}
	return
}
