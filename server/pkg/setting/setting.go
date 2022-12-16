package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

const (
	ENV_DEVELOPMENT = "development"
	ENV_STAGING     = "staging"
	ENV_PRODUCTION  = "production"
)

const (
	OS_MAC   = "mac"
	OS_LINUX = "linux"
)

type AppSetting struct {
	ENV       string
	CORS      string
	JWTSecret string
	OS        string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type GoogleOAuthSetting struct {
	ClientID    string
	SecretID    string
	RedirectURL string
}

type DatabaseSetting struct {
	LocalURL  string
	DockerURL string
}

type ConsulSetting struct {
	Address     string
	ServiceName string
	RootFolder  string
	HealthTTL   time.Duration
	WatchTTL    time.Duration
}

type RedisSetting struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Salt        string
}

var (
	App           = &AppSetting{}
	GoogleOAuth   = &GoogleOAuthSetting{}
	ServerSetting = &Server{}
	Consul        = &ConsulSetting{}
	Redis         = &RedisSetting{}

	Database = &DatabaseSetting{}
)

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", App)
	mapTo("server", ServerSetting)
	mapTo("consul", Consul)
	mapTo("database", Database)
	mapTo("redis", Redis)
	mapTo("google-oauth", GoogleOAuth)

	App.ENV = getEnv(App.ENV)
	App.OS = getOS(App.OS)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	Redis.IdleTimeout = Redis.IdleTimeout * time.Second

	Consul.HealthTTL = Consul.HealthTTL * time.Second
	Consul.WatchTTL = Consul.WatchTTL * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func getEnv(env string) string {
	switch env {
	case ENV_PRODUCTION:
		return ENV_PRODUCTION
	case ENV_STAGING:
		return ENV_STAGING
	default:
		return ENV_DEVELOPMENT
	}
}

func getOS(os string) string {
	switch os {
	case OS_MAC:
		return OS_MAC
	default:
		return OS_LINUX
	}
}
