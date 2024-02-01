package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-tools/fn/fnReflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type (
	Mango struct {
		Id        primitive.ObjectID `bson:"_id"`
		ColNm     string             `bson:"colNm"`
		NextIdx   int                `bson:"nextIdx"`
		History   []*MangoHistory    `bson:"history"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	MangoHistory struct {
		Memo       string    `bson:"memo"`
		MigratedAt time.Time `bson:"migratedAt"`
	}
)

var MangoModel = &MigrateModel{
	ColNm: "mango",
	Migrate: []Migrate{
		func(ctx context.Context, col *mongo.Collection) (memo string, err error) {
			memo = "init indexing"
			_, err = col.Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{
							Key:   "colNm",
							Value: 1,
						},
					},
					Options: &options.IndexOptions{
						Unique: fnReflect.Pointer(true),
					},
				},
			})
			return
		},
	},
}
