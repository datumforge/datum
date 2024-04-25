package logger

type Config struct {
	// App name
	App string `json:"app" koanf:"app" default:"tdr"`

	// App Version
	AppVer string `json:"appVer" koanf:"appVer" default:"v0.0.1"`

	// Log environment (development or production)
	Env string `json:"env" koanf:"env" default:"development"`

	// Location where the system log will be saved
	FileLocation string `json:"fileLocation" koanf:"fileLocation" default:"system.log"`

	// Location where the tdr log will be saved
	FileTDRLocation string `json:"fileTDRLocation" koanf:"fileTDRLocation" default:"tdr.log"`

	// Maximum size of a single log file - when log capacity is reached the file will be rotated
	FileMaxSize int `json:"fileMaxSize" koanf:"fileMaxSize" default:"100"`

	// Maximum number of backup files that will not be deleted - this is useful when you want to rotate but let the older logs fall off
	FileMaxBackup int `json:"fileMaxBackup" koanf:"fileMaxBackup" default:"10"`

	// Number of days where the backup log will not be deleted - controls age of files vs. size
	FileMaxAge int `json:"fileMaxAge" koanf:"fileMaxAge" default:"30"`

	// Log will be printed in console if the value is true
	Stdout bool `json:"stdout" koanf:"stdout" default:"false"`
}
