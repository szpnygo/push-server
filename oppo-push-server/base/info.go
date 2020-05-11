package base

import (
	"os"
)

func GetAppName() string {
	return os.Getenv("app_name")
}

func GetAppKey() string {
	return os.Getenv("app_key")
}

func GetAppMasterSecret() string {
	return os.Getenv("app_master_secret")
}

func GetPushApi() string {
	return "https://api.push.oppomobile.com"
}

func GetFeedBackApi() string {
	return "https://feedback.push.oppomobile.com"
}
