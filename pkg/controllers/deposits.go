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

type getDepositReq struct {
	Asset      string `json:"asset"`
	StartTime  int64  `json:"startTime"`
	EndTime    int64  `json:"endTime"`
	RecvWindow int    `json:"recvWindow"`
	Timestamp  int64  `json:"timestamp"`
}

func (r getDepositReq) new(days int, asset string) getDepositReq {
	ms := time.Now().Unix() * 1000
	r.Asset = asset
	r.StartTime = ms - 1000*60*60*24*int64(days)
	r.EndTime = ms
	r.RecvWindow = 5000
	r.Timestamp = ms
	return r
}

//Deposit !
type Deposit struct {
	InsertTime int64   `json:"insertTime"`
	Amount     float64 `json:"amount"`
	Asset      string  `json:"asset"`
	Address    string  `json:"address"`
	TxID       string  `json:"txId"`
	Status     int     `json:"status"`
}

//GetDepositRes !
type GetDepositRes struct {
	Success     bool      `json:"success"`
	DepositList []Deposit `json:"depositList"`
	Message     string    `json:"msg,omitempty"`
}

//GetDeposits !
func GetDeposits(days int, asset string) GetDepositRes {
	p := "/wapi/v3/depositHistory.html"
	qs := utils.QueryString(getDepositReq{}.new(days, asset))
	sig := utils.GetHmac256(qs)
	url := os.Getenv("API_URL") + p + "?" + qs + "&signature=" + sig

	logrus.Info(url)
	r, err := http.NewRequest("GET", url, nil)
	r.Header.Set("X-MBX-APIKEY", os.Getenv("API_KEY"))

	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		logrus.Error(url)
		return GetDepositRes{}
	}
	defer resp.Body.Close()

	res := &GetDepositRes{}
	bs, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bs, res)

	return *res
}

//GetDepositAddress !
func GetDepositAddress(days int, asset string) GetDepositRes {
	p := "/wapi/v3/depositAddress.html"
	qs := utils.QueryString(getDepositReq{}.new(days, asset))
	sig := utils.GetHmac256(qs)
	url := os.Getenv("API_URL") + p + "?" + qs + "&signature=" + sig

	logrus.Info(url)
	r, err := http.NewRequest("GET", url, nil)
	r.Header.Set("X-MBX-APIKEY", os.Getenv("API_KEY"))

	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		logrus.Error(url)
		return GetDepositRes{}
	}
	defer resp.Body.Close()

	res := &GetDepositRes{}
	bs, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bs, res)

	return *res
}
