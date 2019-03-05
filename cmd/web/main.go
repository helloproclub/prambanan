package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/helloproclub/prambanan"
	"github.com/helloproclub/prambanan/database/mysql"
	"github.com/helloproclub/prambanan/handler"
	"github.com/prometheus/common/log"
	"github.com/subosito/gotenv"
	"net/http"
	"os"
	"path"
)

func main() {
	err := gotenv.Load(path.Join(
		os.Getenv("GOPATH"),
		"/src/github.com/helloproclub/prambanan/.env",
	))
	if err != nil {
		log.Errorln(err)
	}

	prbDatabase, err := mysql.NewMySql(os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalln(err)
	}

	prb, err := prambanan.NewPrambanan(prbDatabase)
	if err != nil {
		log.Fatalln(err)
	}

	handler, err := handler.NewHandler(prb)
	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.Ping).Methods("GET")
	router.HandleFunc("/perform", handler.DecodeNik).Methods("GET")

	log.Infof("Ready at %v:%v", "", os.Getenv("PORT"))
	log.Fatalln(http.ListenAndServe(
		fmt.Sprintf(":%v", os.Getenv("PORT")),
		router,
	))
}
