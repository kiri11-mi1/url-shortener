package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log"
	"strings"
	"sync"
)

type Config struct {
	BaseURL string `env:"BASE_URL, default=http://localhost:8080"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		lookuper := UpcaseLookuper(envconfig.OsLookuper())
		if err := envconfig.ProcessWith(context.Background(), &cfg, lookuper); err != nil {
			log.Fatal(err)
		}
	})
	return cfg
}

type upcaseLookuper struct {
	Next envconfig.Lookuper
}

func (l *upcaseLookuper) Lookup(key string) (string, bool) {
	return l.Next.Lookup(strings.ToUpper(key))
}

func UpcaseLookuper(next envconfig.Lookuper) *upcaseLookuper {
	return &upcaseLookuper{
		Next: next,
	}
}
