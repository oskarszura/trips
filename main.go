package main

import (
    "os"
    "log"
    "fmt"
    "gopkg.in/mgo.v2"
    "github.com/oskarszura/trips/controllers"
    "github.com/oskarszura/trips/controllers/api"
    "github.com/oskarszura/trips/utils"
    gws "github.com/oskarszura/gowebserver"
)

func getServerAddress() (string, error) {
    port := os.Getenv("PORT")

    if port == "" {
        return "", fmt.Errorf("$PORT not set")
    }
    return ":" + port, nil
}

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./version.go

func main() {
    dbUri := os.Getenv("MONGOLAB_URI")
    addr, err := getServerAddress()
    if err != nil {
        panic(err)
    }

    serverOptions := gws.WebServerOptions{
        addr,
        "/static/",
        "public",
    }

    server := gws.New(serverOptions, controllers.NotFound)

    utils.VERSION = VERSION
    log.Println("Starting trips version:", utils.VERSION)

    log.Println("Connecting to mgo with URI = " + dbUri)
    dbSession, err := mgo.Dial(dbUri)
    if err != nil {
        panic(err)
    }
    defer dbSession.Close()
    dbSession.SetMode(mgo.Monotonic, true)
    utils.SetSession(dbSession)

    server.Router.AddRoute("/login/register", controllers.Register)
    server.Router.AddRoute("/login/logout", controllers.AuthenticateLogout)
    server.Router.AddRoute("/login", controllers.Authenticate)
    server.Router.AddRoute("/trips", controllers.Trips)
    server.Router.AddRoute("/travel-map", controllers.TravelMap)
    server.Router.AddRoute("/", controllers.Front)
    server.Router.AddRoute("/api/trips", api.CtrTrips)
    server.Router.AddRoute("/api/trips/{id}", api.CtrTrip)
    server.Router.AddRoute("/api/places", api.CtrPlaces)
    server.Router.AddRoute("/api/places/{id}", api.CtrPlace)

    server.Run()
}
