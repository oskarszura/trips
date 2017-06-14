package api

import (
	"log"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/oskarszura/trips/utils"
	. "github.com/oskarszura/trips/models"
)

type TripList []Trip

func Trips(w http.ResponseWriter, r *http.Request, params struct{Params map[string]string}) {
	var trips []Trip

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ds := utils.GetDataSource()
	c := ds.C("trips")

	switch r.Method {
	case "GET":
		pipe := c.Pipe([]bson.M{
			{"$match": bson.M{"userid": utils.GetUser().Id}},
			{"$lookup": bson.M{
			"from":"places",
			"localField": "id",
			"foreignField": "tripid",
			"as": "places",
		}}})

		trips = TripList{}
		err := pipe.All(&trips)

		if err != nil {
			log.Fatal(err)
		}

		if trips == nil {
			trips = TripList{}
		}

		output := trips

		json.NewEncoder(w).Encode(output)
	case "POST":
		var newTrip Trip

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newTrip)

		if err != nil {
			panic(err)
		}

		newTrip = Trip{
			Id: bson.NewObjectId(),
			UserId: utils.GetUser().Id,
			Name: newTrip.Name,
			Places: PlaceList{},
		}

		err = c.Insert(newTrip)

		if err != nil {
			log.Fatal(err)
		}

		output := newTrip

		json.NewEncoder(w).Encode(output)
	default:
	}
}

