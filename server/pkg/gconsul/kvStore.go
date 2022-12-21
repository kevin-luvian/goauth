package gconsul

import (
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	"github.com/kevin-luvian/goauth/server/pkg/util"
)

type KVStore struct {
	App struct {
		JWTAccessSecret  string `json:"jwt_access_secret"`
		JWTRefreshSecret string `json:"jwt_refresh_secret"`
		CORS             string `json:"cors"`
	} `json:"app"`
	GoogleOAuth struct {
		SecretID string `json:"secret_id"`
	} `json:"google_oauth"`
	Redis struct{} `json:"redis"`
}

var storeHash string

func FetchKV() error {
	store, err := instance.fetchKVStore()
	if err != nil {
		return err
	}

	storeHash = util.EncodeMD5(store)

	setting.App.JWTAccessSecret = nEmptyFill(setting.App.JWTAccessSecret, store.App.JWTAccessSecret)
	setting.App.JWTRefreshSecret = nEmptyFill(setting.App.JWTRefreshSecret, store.App.JWTRefreshSecret)
	setting.App.CORS = nEmptyFill(setting.App.CORS, store.App.CORS)

	setting.GoogleOAuth.SecretID = nEmptyFill(setting.GoogleOAuth.SecretID, store.GoogleOAuth.SecretID)

	return nil
}

func HasKVChanged() bool {
	newStore, err := instance.fetchKVStore()
	if err != nil {
		logging.Errorln("fetching kv failed", err.Error())
		return false
	}

	newStoreHash := util.EncodeMD5(newStore)

	return storeHash != newStoreHash
}

func (c *Consul) fetchKVStore() (KVStore, error) {
	kvs := KVStore{}

	m, err := util.StructToMap(kvs)
	if err != nil {
		return kvs, err
	}

	m = fillInMap(m, c.RootFolder, func(path string) interface{} {
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
			v[key] = fillInMap(val, path+"/"+key, getVal)
		}
		return v
	default:
		return getVal(path)
	}
}

func nEmptyFill(base, new string) string {
	if new != "" {
		return new
	}
	return base
}
