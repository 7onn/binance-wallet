package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/devbytom/binance-wallet/pkg/utils"
)

type getAssetsReq struct {
	RecvWindow int   `json:"recvWindow"`
	Timestamp  int64 `json:"timestamp"`
}

func (r getAssetsReq) new() getAssetsReq {
	ms := time.Now().Unix() * 1000
	r.RecvWindow = 5000
	r.Timestamp = ms
	return r
}

//Assets !
type Assets map[string]AssetData

//AssetData !
type AssetData struct {
	MinWithdrawAmount string `json:"minWithdrawAmount"`
	DepositStatus     bool   `json:"depositStatus"`
	WithdrawFee       int    `json:"withdrawFee"`
	WithdrawStatus    bool   `json:"withdrawStatus"`
	DepositTip        string `json:"depositTip"`
}

//GetAssetsRes !
type GetAssetsRes struct {
	Success     bool   `json:"success"`
	AssetDetail Assets `json:"assetDetail"`
	Message     string `json:"msg,omitempty"`
}

//GetAssets !
func GetAssets() GetAssetsRes {
	p := "/wapi/v3/assetDetail.html"
	qs := utils.QueryString(getAssetsReq{}.new())
	sig := utils.GetHmac256(qs)
	url := os.Getenv("API_URL") + p + "?" + qs + "&signature=" + sig

	logrus.Info(url)
	r, err := http.NewRequest("GET", url, nil)
	r.Header.Set("X-MBX-APIKEY", os.Getenv("API_KEY"))

	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		logrus.Error(url)
		return GetAssetsRes{}
	}
	defer resp.Body.Close()

	res := &GetAssetsRes{}
	bs, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bs, res)

	return *res
}
