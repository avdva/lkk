// Copyright 2019 Aleksandr Demakin. All rights reserved.

package lkk

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

// CandlesHistoryInfo contains common fields for candles requests and responses.
type CandlesHistoryInfo struct {
	AssetPair string    `json:"assetPair"`
	Period    string    `json:"period"`
	DateFrom  time.Time `json:"dateFrom"`
	DateTo    time.Time `json:"dateTo"`
	Type      string    `json:"type"`
}

// CandlesHistoryResp is a response for '/Candles/history' EP.
type CandlesHistoryResp struct {
	CandlesHistoryInfo
	Data []CandleData `json:"data"`
}

// CandleData contains candle info.
type CandleData struct {
	Ts             time.Time       `json:"dateTime"`
	Open           decimal.Decimal `json:"open"`
	Low            decimal.Decimal `json:"low"`
	High           decimal.Decimal `json:"high"`
	Close          decimal.Decimal `json:"close"`
	Volume         float64         `json:"volume"`
	OppositeVolume float64         `json:"oppositeVolume"`
}

type candles struct {
	client *http.Client
}

// History returns candles for given market and request.
func (c *candles) History(market string, req *CandlesHistoryInfo) (*CandlesHistoryResp, error) {
	var resp CandlesHistoryResp
	reqURL := fmt.Sprintf("Candles/history/%s/%s/%s/%s/%s/%s",
		market, req.AssetPair, req.Period, req.Type, req.DateFrom.Format(time.RFC3339Nano), req.DateTo.Format(time.RFC3339Nano))
	if err := get(c.client, requestPath+reqURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
