package setting

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
	"github.com/kevin-luvian/goauth/server/pkg/util"
)

type AppSetting struct {
	ENV  string
	CORS string
	OS   string

	JWTAccessSecret  string
	JWTRefreshSecret string

	TickerTTL time.Duration

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type ServerSetting struct {
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
	LocalURL    string
	DockerURL   string
	Retries     int
	MaxActive   int
	MaxIdle     int
	MaxLifetime time.Duration
}

type ConsulSetting struct {
	Address     string
	ServiceName string
	RootFolder  string
}

type RedisSetting struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Salt        string
}

const (
	ENV_DEVELOPMENT = "development"
	ENV_STAGING     = "staging"
	ENV_PRODUCTION  = "production"

	OS_MAC   = "mac"
	OS_LINUX = "linux"
)

var (
	App         = &AppSetting{}
	GoogleOAuth = &GoogleOAuthSetting{}
	Server      = &ServerSetting{}
	Consul      = &ConsulSetting{}
	Redis       = &RedisSetting{}

	Database = &DatabaseSetting{}
)

var (
	filepath = "conf/app.ini"
	checksum string
)

// Setup initialize the configuration instance
func Setup() {
	cfg, err := ini.Load(filepath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	checksum = getFileChecksum()

	mapTo(cfg, "app", App)
	mapTo(cfg, "server", Server)
	mapTo(cfg, "consul", Consul)
	mapTo(cfg, "database", Database)
	mapTo(cfg, "redis", Redis)
	mapTo(cfg, "google-oauth", GoogleOAuth)

	App.ENV = getEnv(App.ENV)
	App.OS = getOS(App.OS)
	App.TickerTTL = App.TickerTTL * time.Second

	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.WriteTimeout * time.Second

	Redis.IdleTimeout = Redis.IdleTimeout * time.Second

	Database.MaxLifetime = Database.MaxLifetime * time.Second
}

func HasSettingChanged() bool {
	return checksum != getFileChecksum()
}

func getFileChecksum() string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to open 'conf/app.ini': %v", err)
	}

	defer file.Close()

	return util.EncodeMD5File(file)
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
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
