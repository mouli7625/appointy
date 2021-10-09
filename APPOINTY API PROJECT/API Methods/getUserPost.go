package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func GetUserPost(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	
	var params = mux.Vars(r)

	id, _ := params["id"]

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filterCursor, err := postCollection.Find(ctx, bson.M{"userid": id})
	if err != nil {
		log.Fatal(err)
	}
	var allUserPost []bson.M
	if err = filterCursor.All(ctx, &allUserPost); err != nil {
		log.Fatal(err)
	}
	fmt.Println(allUserPost)
	

    filter:=bson.M{}
	findOptions := options.Find()
 	page, _ := postCollection.CountDocuments(ctx, bson.M{"page":"1"})
	  
		var perPage int64 = 10

		total, _ := postCollection.CountDocuments(ctx, filter)
        fmt.Println(total)
		findOptions.SetSkip((int64(page) - 1) * perPage)
		findOptions.SetLimit(perPage)

		
	json.NewEncoder(w).Encode(allUserPost)
}
