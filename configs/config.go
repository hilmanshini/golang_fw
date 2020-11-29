package configs

type config struct {
	JwtKey []byte
}

var appconfig *config

// GetConfig() asdasd
func GetConfig() *config {
	if appconfig == nil {
		appconfig = &config{[]byte("1234123412341234")}
	}
	return appconfig
}
