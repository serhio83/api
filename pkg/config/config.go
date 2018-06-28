package config

import (
	"log"
	"os"

	"git.topcloud.ru/topcloud.ru/api/pkg"
	"github.com/BurntSushi/toml"
)

type Cfg struct {
	ApiMongoUrl		string
	ApiMongoDB		string
	ApiServerPort	string
}

func GetConfig() *root.Config {
	var c Cfg
	ReadConfig(&c)
  return &root.Config {
    Mongo: &root.MongoConfig { Url: c.ApiMongoUrl, DbName: c.ApiMongoDB},
    Server: &root.ServerConfig { Port: c.ApiServerPort},
    Auth: &root.AuthConfig { Secret: envOrDefaultString("API_AUTH_SECRET", "mysecret")}}
}

func envOrDefaultString(envVar string, defaultValue string) string {
  value := os.Getenv(envVar)
  if value == "" {
    return defaultValue
  }
  
  return value
}

func ReadConfig(c *Cfg) *Cfg {
	if _, err := toml.DecodeFile("../../cfg/config.toml", &c); err != nil {
		log.Println(err)
	}
	return c
}
