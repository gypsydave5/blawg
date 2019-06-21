package blawg

import "os"
import "github.com/BurntSushi/toml"

func GetSiteMeta() Config {
	f, _ := os.Open("config.toml")
	var c Config

	toml.DecodeReader(f, &c)
	return c
}

type Config struct {
	URL         string `toml:"URL"`
	Description string
	Title       string
}
