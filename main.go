package main

import (
	"net/http"
	"os"

	"github.com/devbytom/binance-wallet/pkg/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	godotenv.Load()

	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", router.Healthcheck)
	r.HandleFunc("/deposits", router.GetDeposits)

	http.Handle("/", r)
	p := os.Getenv("PORT")
	logrus.Info("listening at port: " + p)
	logrus.Fatal(http.ListenAndServe(":"+p, r))
}
