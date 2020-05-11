package auth

import (
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"github.com/szpnygo/go-tools"
	"neobaran.com/server/push/oppo/base"
	"net/url"
	"strconv"
	"time"
)

var bm, cacheErr = cache.NewCache("memory", `{"interval":60}`)

func RequestAuthToken() (base.AuthToken, error) {
	var authToken base.AuthTokenResult

	appKey := base.GetAppKey()
	appMasterSecret := base.GetAppMasterSecret()
	timestamp := time.Now().UnixNano() / 1e6

	logs.Info(strconv.Itoa(int(timestamp)))

	data := url.Values{
		"app_key":   {appKey},
		"sign":      {base.GetSha256([]byte(appKey + strconv.Itoa(int(timestamp)) + appMasterSecret))},
		"timestamp": {strconv.Itoa(int(timestamp))},
	}

	requestUrl := base.GetPushApi() + "/server/v1/auth"

	body, err := neo.Post(requestUrl, data, nil)
	if err != nil {
		logs.Error(err)
		return authToken.Data, err
	}
	logs.Info("request auth result : " + body)

	jsonErr := json.Unmarshal([]byte(body), &authToken)

	if jsonErr != nil {
		logs.Error(jsonErr)
		return authToken.Data, jsonErr
	}

	if authToken.Code != 0 {
		return authToken.Data, errors.New(authToken.Message)
	}

	if cacheErr != nil {
		logs.Info(cacheErr)
		return authToken.Data, nil
	}

	bm.Put("auth_token", authToken.Data, 24*time.Hour)
	logs.Info("get auth token : " + authToken.Data.AuthToken)

	return authToken.Data, nil

}

func GetAuthToken() (base.AuthToken, error) {
	var accessToken base.AuthToken

	if cacheErr != nil {
		accessToken, err := RequestAuthToken()
		if err != nil {
			return accessToken, err
		}
		return accessToken, err
	}
	cacheData := bm.Get("auth_token")
	if cacheData == nil {
		accessToken, err := RequestAuthToken()
		if err != nil {
			return accessToken, err
		}
		return accessToken, err
	}
	accessToken = cacheData.(base.AuthToken)
	return accessToken, nil

}
