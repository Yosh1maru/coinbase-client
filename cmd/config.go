package main

// Config app configuration file
type Config struct {
	// DbHost db host
	DbHost string `env:"MYSQL_HOST"`
	// DbName database for connection
	DbName string `env:"MYSQL_DATABASE"`
	// DbUser db user
	DbUser string `env:"MYSQL_ROOT_USER"`
	// DbPassword db user password
	DbPassword string `env:"MYSQL_ROOT_PASSWORD"`
	// SourceUrl Source url
	SourceUrl string `env:"SOURCE_URL"`
	// SourceUrl Source url
	Type string `env:"TYPE"`
	// SourceUrl Source url
	ProductIds []string `env:"PRODUCT_IDS"`
	// SourceUrl Source url
	Channels []string `env:"CHANNELS"`
	// SleepDurationTillEnd waiting duration for ending of application in seconds
	SleepDurationTillEnd int `env:"SLEEP_DURATION_TILL_END"`
}
