package router

import (
	"encoding/json"
	"net/http"

	"github.com/devbytom/bnb/pkg/controllers"
	"github.com/devbytom/bnb/pkg/utils"
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
