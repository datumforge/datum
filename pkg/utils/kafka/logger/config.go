package logger

type Config struct {
	// App name
	App string `json:"app"`

	// App Version
	AppVer string `json:"appVer"`

	// Log environment (development or production)
	Env string `json:"env"`

	// Location where the system log will be saved
	FileLocation string `json:"fileLocation"`

	// Location where the tdr log will be saved
	FileTDRLocation string `json:"fileTDRLocation"`

	// Maximum size of a single log file.
	// If the capacity reach, file will be saved but it will be renamed
	// with suffix the current date
	FileMaxSize int `json:"fileMaxSize"`

	// Maximum number of backup file that will not be deleted
	FileMaxBackup int `json:"fileMaxBackup"`

	// Number of days where the backup log will not be deleted
	FileMaxAge int `json:"fileMaxAge"`

	// Log will be printed in console if the value is true
	Stdout bool `json:"stdout"`
}
