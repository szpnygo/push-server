package base

type ApiBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AuthToken struct {
	AuthToken  string `json:"auth_token"`
	CreateTime int    `json:"create_time"`
}

type AuthTokenResult struct {
	ApiBase
	Data AuthToken `json:"data"`
}
