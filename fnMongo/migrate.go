package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-snippet/fn/fnMatch"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type (
	MigrateModel struct {
		ColNm   string
		Migrate []Migrate
	}

	Migrate func(ctx context.Context, col *mongo.Collection) (memo string, err error)
)

func (x *MigrateModel) AppendMigrate(ls []Migrate) {
	x.Migrate = append(x.Migrate, ls...)
}

func RunMigrate(
	ctx context.Context,
	models ...*MigrateModel,
) (err error) {
	var modelList = make([]*MigrateModel, 0)
	modelList = append(modelList, MangoModel)
	modelList = append(modelList, models...)

	// create collection
	var db = GetDBP(ctx)
	if err = createCollection(ctx, modelList); err != nil {
		return
	}

	// exec migrate
	var colMango = db.Collection(MangoModel.ColNm)
	var now = time.Now()

	for _, model := range modelList {
		var count int64
		if count, err = colMango.CountDocuments(ctx, bson.M{
			"colNm": model.ColNm,
		}); err != nil {
			return
		}

		if count == 0 {
			if _, err = colMango.InsertOne(
				ctx,
				&Mango{
					Id:        primitive.NewObjectID(),
					ColNm:     model.ColNm,
					NextIdx:   0,
					History:   make([]*MangoHistory, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
			); err != nil {
				return
			}
		}

		var cur *mongo.SingleResult
		if cur = colMango.FindOne(
			ctx,
			bson.M{
				"colNm": model.ColNm,
			},
		); cur.Err() != nil {
			err = cur.Err()
			return
		}

		var doc = new(Mango)
		if err = cur.Decode(doc); err != nil {
			return
		}

		var colModel = db.Collection(model.ColNm)
		var migrateList = model.Migrate

		for i := doc.NextIdx; i < len(migrateList); i++ {
			var fn = migrateList[i]
			var memo string
			if memo, err = fn(ctx, colModel); err != nil {
				return
			}

			if _, err = colMango.UpdateOne(
				ctx,
				bson.M{
					"colNm": model.ColNm,
				},
				bson.M{
					"$push": bson.M{
						"history": &MangoHistory{
							Memo:       memo,
							MigratedAt: time.Now(),
						},
					},
					"$inc": bson.M{
						"nextIdx": 1,
					},
				}); err != nil {
				return err
			}
		}
	}

	return
}

func createCollection(ctx context.Context, models []*MigrateModel) (err error) {
	var db = GetDBP(ctx)
	var names []string
	if names, err = db.ListCollectionNames(ctx, bson.M{}); err != nil {
		return
	}

	for _, model := range models {
		if _, err = fnMatch.Get(names, func(v string) bool {
			return v == model.ColNm
		}); err != nil {
			continue
		}

		if err = db.CreateCollection(ctx, model.ColNm); err != nil {
			return
		}
	}
	return
}
