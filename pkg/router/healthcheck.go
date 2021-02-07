package router

import (
	"net/http"

	"github.com/devbytom/binance-wallet/pkg/utils"
)

//Healthcheck !
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, utils.Message(true, "alive"))
}
