package main

const APIBackend = "https://partner.steam-api.com"

var (
	AppVersion     = "Undefined"
	ConfigFilepath = "steam-config.yaml"
	LogLevel       = "info"
	AppConfig      Config
)

type Config struct {
	Rpc   RpcConfig   `yaml:"rpc"`
	Steam SteamConfig `yaml:"steam"`
}

type RpcConfig struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

type SteamConfig struct {
	PublisherId string `yaml:"publisher_id"`
	AppId       uint32 `yaml:"app_id"`
	Key         string `yaml:"key"`
}
