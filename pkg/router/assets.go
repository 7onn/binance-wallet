package router

import (
	"net/http"

	"github.com/devbytom/binance-wallet/pkg/controllers"
	"github.com/devbytom/binance-wallet/pkg/utils"
	"github.com/sirupsen/logrus"
)

//GetAssets !
func GetAssets(w http.ResponseWriter, r *http.Request) {
	res := controllers.GetAssets()

	if !res.Success {
		logrus.Error(res.Message)
		utils.Respond(w, utils.Message(false, res.Message))
		return
	}

	utils.Respond(w, res)
}
