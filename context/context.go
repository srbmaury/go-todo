package contextapi

import (
	"context"
)

type MyContextKey string

var (
	appContext context.Context
)

func CreateContext() {
	appContext = context.Background()
}

func SetValue(key MyContextKey, value string) {
	appContext = context.WithValue(appContext, key, value)
}

func GetValue(key MyContextKey) string {
	val := appContext.Value(key)
	if val == nil {
		return ""
	}
	return val.(string)
}
