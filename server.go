package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/rjelierse/competence-server/competence"
	"github.com/rjelierse/competence-server/mongo"
)

func main() {
	err := mongo.Db.Connect("localhost")
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	// Set up CORS
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods:                []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:                []string{"Accept", "Content-Type", "Origin", "Authorization"},
		AccessControlAllowCredentials: false,
		AccessControlExposeHeaders:    []string{"WWW-Authenticate"},
		AccessControlMaxAge:           3600,
	})

	router, err := rest.MakeRouter(
		rest.Get("/profiles", competence.GetAllProfiles),
		rest.Post("/profiles", competence.CreateProfile),
		rest.Get("/profiles/:id", competence.GetProfile),
		rest.Put("/profiles/:id", competence.UpdateProfile),
		rest.Delete("/profiles/:id", competence.DeleteProfile),
	)
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	fmt.Println("Accepting incoming connections at http://localhost:9212")
	log.Fatal(http.ListenAndServe(":9212", api.MakeHandler()))
}
