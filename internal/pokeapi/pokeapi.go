package pokeapi

import (
	"net/http"
	"time"
	"github.com/lucasSSalgado/pokedex/internal/pokecache"
)

type Client struct {
	client http.Client
	memory pokecache.Cache
}

func NewClient() Client {
	customClient := Client{
		client: http.Client{
			Timeout: time.Minute,
		},
		memory: pokecache.NewCache(),
	}
	return customClient
}