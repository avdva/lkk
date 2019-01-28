// Copyright 2019 Aleksandr Demakin. All rights reserved.

package lkk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Public is a client for public API
type Public struct {
	client *http.Client
}

// NewPublic returns new Public client.
func NewPublic() *Public {
	return &Public{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// AssetPairs returns '/AssetPairs' EP accessor.
func (p *Public) AssetPairs() *assetPairs {
	return &assetPairs{client: p.client}
}

// Candles returns '/Candles/history' EP accessor.
func (p *Public) Candles() *candles {
	return &candles{client: p.client}
}

func get(client *http.Client, url string, obj interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(obj)
}
