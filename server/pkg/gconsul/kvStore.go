package gconsul

import (
	"reflect"
	"strings"
	"time"

	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	"github.com/kevin-luvian/goauth/server/pkg/util"
)

type KVStore struct {
	App struct {
		JWTSecret string `json:"jwt_secret"`
		CORS      string
	}
	GoogleOauth struct {
		SecretID string `json:"secret_id"`
	} `json:"google_oauth"`
	Redis struct {
		Password string
	}
}

var store KVStore

func FetchKV() (err error) {
	store, err = instance.fetchKVStore()
	if err != nil {
		return err
	}

	if store.App.JWTSecret != "" {
		setting.App.JWTSecret = store.App.JWTSecret
	}
	if store.App.CORS != "" {
		setting.App.CORS = store.App.CORS
	}

	if store.GoogleOauth.SecretID != "" {
		setting.GoogleOAuth.SecretID = store.GoogleOauth.SecretID
	}

	if store.Redis.Password != "" {
		setting.Redis.Password = store.Redis.Password
	}

	return nil
}

func WatchKV(f func()) {
	checkKV := func() bool {
		logging.Infoln("checking KV Store", setting.App.JWTSecret)

		newStore, err := instance.fetchKVStore()
		if err != nil {
			logging.Errorln("fetching kv failed", err.Error())
			return false
		}

		return reflect.DeepEqual(store, newStore)
	}

	go func() {
		ticker := time.NewTicker(instance.WatchTTL)
		for range ticker.C {
			if !checkKV() {
				f()
			}
		}
	}()
}

func (c *Consul) fetchKVStore() (KVStore, error) {
	kvs := KVStore{}

	m, err := util.StructToMap(kvs)
	if err != nil {
		return kvs, err
	}

	m = fillInMap(m, c.Name, func(path string) interface{} {
		kvpair, _, _ := c.ConsulKV.Get(path, nil)
		if kvpair != nil {
			return string(kvpair.Value)
		}

		return ""
	}).(map[string]interface{})

	err = util.MapToStruct(m, &kvs)
	if err != nil {
		return kvs, err
	}

	return kvs, nil
}

func fillInMap(i interface{}, path string, getVal func(path string) interface{}) interface{} {
	switch v := i.(type) {
	case map[string]interface{}:
		for key, val := range v {
			v[key] = fillInMap(val, path+"/"+strings.ToLower(key), getVal)
		}
		return v
	default:
		return getVal(path)
	}
}