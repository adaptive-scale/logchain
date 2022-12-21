package internal

type Configuration struct {
	DatabaseType string `yaml:"database_type" mapstructure:"DATABASE_TYPE"`
	DSN          string `yaml:"dsn"  mapstructure:"DSN"`
	Port         int    `yaml:"port"  mapstructure:"PORT"`
	Protected    bool   `yaml:"protected"  mapstructure:"PROTECTED"`
	CertLocation string `yaml:"cert_location"  mapstructure:"CERT_LOCATION"`
	KeyLocation  string `yaml:"key_location"  mapstructure:"KEY_LOCATION"`
}
