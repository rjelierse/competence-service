package competence

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/rjelierse/competence-server/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Profile contains a user's competence profile
type Profile struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Competences []Competence  `bson:"competences" json:"competences"`
	TotalPoints int           `bson:"totalPoints" json:"total_points"`
	Locked      bool          `bson:"locked" json:"locked"`
	// student User
	// coach User
}

func getCollection() *mgo.Collection {
	return mongo.Db.Collection("go-test", "profiles")
}

// GetAllProfiles handles 'GET /profiles'
func GetAllProfiles(response rest.ResponseWriter, request *rest.Request) {
	profiles := []Profile{}

	if err := getCollection().Find(bson.M{}).All(&profiles); err != nil {
		rest.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.WriteJson(&profiles)
}

// CreateProfile handles 'POST /profiles'
func CreateProfile(response rest.ResponseWriter, request *rest.Request) {
	profile := Profile{}

	if err := request.DecodeJsonPayload(&profile); err != nil {
		rest.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	if err := getCollection().Insert(&profile); err != nil {
		rest.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.WriteJson(&profile)
}

// GetProfile handles 'GET /profiles/:id'
func GetProfile(response rest.ResponseWriter, request *rest.Request) {
	id := request.PathParam("id")
	profile := Profile{}

	if err := getCollection().FindId(bson.ObjectIdHex(id)).One(&profile); err != nil {
		rest.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.WriteJson(&profile)
}

// UpdateProfile handles 'PUT /profiles/:id'
func UpdateProfile(response rest.ResponseWriter, request *rest.Request) {
	id := request.PathParam("id")
	profile := Profile{}

	if err := request.DecodeJsonPayload(&profile); err != nil {
		rest.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	if err := getCollection().UpdateId(bson.ObjectIdHex(id), &profile); err != nil {
		rest.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.WriteJson(&profile)
}

// DeleteProfile handles 'DELETE /profiles/:id'
func DeleteProfile(response rest.ResponseWriter, request *rest.Request) {
	id := request.PathParam("id")

	if err := getCollection().RemoveId(bson.ObjectIdHex(id)); err != nil {
		rest.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
