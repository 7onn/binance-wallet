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

func (r getAssetsReq) queryString() string {
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

//Asset !
type Asset map[string]AssetData

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
	AssetDetail Asset  `json:"assetDetail"`
	Message     string `json:"msg,omitempty"`
}

//GetAssets !
func GetAssets() GetAssetsRes {
	p := "/wapi/v3/assetDetail.html"
	qs := getAssetsReq{}.new().queryString()
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
