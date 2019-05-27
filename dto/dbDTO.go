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
