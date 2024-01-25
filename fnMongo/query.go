package fnMongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/*------------------------------------------------------------------------------------------------*/

type Time struct {
	Equal    *time.Time
	NotEqual *time.Time
	GT       *time.Time
	GTE      *time.Time
	LT       *time.Time
	LTE      *time.Time
}

func (x Time) Query() (res bson.M, err error) {
	res = make(bson.M)

	if x.Equal != nil {
		res["$equal"] = *x.Equal
	}

	if x.NotEqual != nil {
		res["$ne"] = *x.NotEqual
	}

	if x.GT != nil {
		res["$gt"] = *x.GT
	}

	if x.GTE != nil {
		res["$gte"] = *x.GTE
	}

	if x.LT != nil {
		res["$lt"] = *x.LT
	}

	if x.LTE != nil {
		res["$lte"] = *x.LTE
	}

	if len(res) == 0 {
		err = fmt.Errorf("no has time query")
		return
	}

	return
}

/*------------------------------------------------------------------------------------------------*/
