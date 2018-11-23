package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	DbConfig     = &Db{}
	ServerConfig = &Server{}
	AppConfig    = &App{}
	LogConfig    = &Log{}
)

type Db struct {
	Type        string
	User        string
	PassWord    string
	Host        string
	Name        string
	TablePrefix string
	Debug       bool
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type App struct {
	PageSize        int
	JwtSecret       string
	RuntimeRootPath string
	ImagePrefixUrl  string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExts  []string
}

type Log struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	LogLevel    string
	TimeFormat  string
	Debug       bool
}

func InitSetUp(config_file_path string) {
	cfg, err := ini.Load(config_file_path)
	if err != nil {
		log.Fatalf("error:cant not load 'conf/app.ini':%s", err)
	}

	err = cfg.Section("database").MapTo(DbConfig)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %s", err)
	}

	err = cfg.Section("server").MapTo(ServerConfig)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerConfig err: %s", err)
	}
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second

	err = cfg.Section("app").MapTo(AppConfig)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	AppConfig.ImageMaxSize = AppConfig.ImageMaxSize * 1024 * 1024

	err = cfg.Section("log").MapTo(LogConfig)
	if err != nil {
		log.Fatalf("Cfg.MapTo LogConfig err: %v", err)
	}

}
