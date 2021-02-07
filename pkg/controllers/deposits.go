package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
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

func (r getDepositReq) parseQueryString() string {
	o := reflect.ValueOf(r)
	s := ""

	for i := 0; i < o.NumField(); i++ {
		p := utils.LowerFirstChar(o.Type().Field(i).Name)
		s = s + p

		f := o.Field(i)
		switch f.Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			s = s + "=" + strconv.FormatInt(f.Int(), 10)
		default:
			s = s + "=" + f.String()
		}

		s = s + "&"

	}
	return strings.TrimSuffix(s, "&")
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
	qs := getDepositReq{}.new(days, asset).parseQueryString()
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
