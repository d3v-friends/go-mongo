package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-tools/fn/fnParam"
	"go.mongodb.org/mongo-driver/mongo"
)

func IndexCollection(
	ctx context.Context,
	colNm string,
	idx []mongo.IndexModel,
	isInits ...bool,
) (_ []string, err error) {
	var col = GetColP(ctx, colNm)

	if fnParam.Get(isInits) {
		if _, err = col.Indexes().DropAll(ctx); err != nil {
			return
		}
	}

	return col.Indexes().CreateMany(ctx, idx)
}

/*------------------------------------------------------------------------------------------------*/

type IndexResp struct {
	Key    map[string]int64 `bson:"key"`
	Name   string           `bson:"name"`
	V      int64            `bson:"v"`
	Unique *bool            `bson:"unique"`
}

type IndexRespList []*IndexResp

func (x IndexRespList) StringList() (ls []string) {
	ls = make([]string, len(x))
	for i, resp := range x {
		ls[i] = resp.Name
	}
	return
}

func GetIndexList(
	ctx context.Context,
	colNm string,
) (ls IndexRespList, err error) {

	var col = GetColP(ctx, colNm)

	var cur *mongo.Cursor
	if cur, err = col.Indexes().List(ctx); err != nil {
		return
	}

	ls = make(IndexRespList, 0)
	if err = cur.All(ctx, &ls); err != nil {
		return
	}

	return
}
