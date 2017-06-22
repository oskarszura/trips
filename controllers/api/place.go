package api

import (
	"log"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/oskarszura/trips/utils"
)


func CtrPlace(w http.ResponseWriter, r *http.Request, options struct{Params map[string]string}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ds := utils.GetDataSource()
	c := ds.C("places")

	switch r.Method {
	case "DELETE":
		placeId := options.Params["id"]
		err := c.Remove(bson.M{"_id": bson.ObjectIdHex(placeId)})

		if err != nil {
			log.Fatalln(err)
		}

		output := &utils.HalResponse{
			Status: 200,
		}

		json.NewEncoder(w).Encode(output)
	default:
	}
}

