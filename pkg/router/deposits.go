package router

import (
	"encoding/json"
	"net/http"

	"github.com/devbytom/binance-wallet/pkg/controllers"
	"github.com/devbytom/binance-wallet/pkg/utils"
	"github.com/sirupsen/logrus"
)

type getDepositReq struct {
	Asset string `json:"asset"`
	Days  int    `json:"days"`
}

//GetDeposits !
func GetDeposits(w http.ResponseWriter, r *http.Request) {
	body := getDepositReq{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	res := controllers.GetDeposits(body.Days, body.Asset)

	if !res.Success {
		logrus.Error(res.Message)
		utils.Respond(w, utils.Message(false, res.Message))
		return
	}

	utils.Respond(w, res)
}

//GetDepositAddress !
func GetDepositAddress(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	asset := qs.Get("asset")
	if asset == "" {
		utils.Respond(w, utils.Message(false, "must provide 'asset' parameter on query string"))
		return
	}

	res := controllers.GetDepositAddress(asset)

	if !res.Success {
		logrus.Error(res.Message)
		utils.Respond(w, utils.Message(false, res.Message))
		return
	}

	utils.Respond(w, res)
}
