package gridscale

import (
	"log"
	"github.com/parce-iot/gridscale"
	"time"
)

type Config struct {
	*gridscale.Client
	Endpoint    string
	AuthToken string
	UserId    string
	Timeout   time.Duration
}


// Create creates a generic gridscale client
func (c *Config) CreateClient() error {
	if c.Endpoint == "" || c.UserId == "" || c.AuthToken == "" {
		return nil
	}

	client, err := gridscale.NewClient(c.UserId, c.AuthToken, c.Endpoint)
	if err != nil {
		return err
	}
	c.Client = client

	log.Printf("[INFO] Rancher Client configured for url: %s", c.Endpoint)


	return nil
}
