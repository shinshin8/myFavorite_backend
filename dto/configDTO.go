package dto

// DbConfig is a struct for a configuration for database.
type DbConfig struct {
	Database DbInfo
}

// DbInfo struct is the configuration for database information.
type DbInfo struct {
	DriverName string
	User       string
	Password   string
	Database   string
	Host       string
	DbPort     string
}

// configuration file
type PortConfig struct {
	Port PortInfo
}

// a port part in configuration file
type PortInfo struct {
	Port string
}

// LogConfig is a struct for configuration of a log file path.
type LogConfig struct {
	LogFile LogFile
}

// LogFile is a sturct for a log file path.
type LogFile struct {
	LogFile string
}

// TokenStringConfig is a struct for configuration of a toke string.
type TokenStringConfig struct {
	TokenString TokenString
}

// TokenString is a struct for token string.
type TokenString struct {
	TokenString string
}
