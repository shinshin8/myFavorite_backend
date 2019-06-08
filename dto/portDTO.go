package dto

// configuration file
type PortConfig struct {
	Port PortInfo
}

// a port part in configuration file
type PortInfo struct {
	Port string
}

type RedisConfig struct {
	Redis RedisInfo
}

type RedisInfo struct {
	RedisAddress string
}