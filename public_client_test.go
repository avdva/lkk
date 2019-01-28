// Copyright 2019 Aleksandr Demakin. All rights reserved.

package lkk

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPublicAssetPairsDictionary(t *testing.T) {
	r := require.New(t)

	p := NewPublic()
	dict, err := p.AssetPairs().Dictionary(MarketMt)
	r.NoError(err)
	r.NotEmpty(dict)

	dict, err = p.AssetPairs().Dictionary(MarketSpot)
	r.NoError(err)
	r.NotEmpty(dict)
}

func TestPublicAssetPairsRates(t *testing.T) {
	r := require.New(t)

	p := NewPublic()

	rate, err := p.AssetPairs().Rate("EURUSD")
	r.NoError(err)
	r.NotEmpty(rate.ID)
	r.True(rate.Ask.Sign() > 0)
	r.True(rate.Bid.Sign() > 0)

	rates, err := p.AssetPairs().Rates()
	r.NoError(err)
	r.True(len(rates) > 0)

}

func TestPublicCandlesHistory(t *testing.T) {
	r := require.New(t)

	p := NewPublic()

	from := time.Date(2019, 1, 25, 10, 0, 0, 0, time.UTC)
	req := CandlesHistoryInfo{
		AssetPair: "EURUSD",
		Period:    TmHour,
		DateFrom:  from,
		DateTo:    from.Add(8 * time.Hour),
		Type:      PriceMid,
	}

	resp, err := p.Candles().History(MarketSpot, &req)
	r.NoError(err)
	r.NotEmpty(resp.Data)

	r.Equal(8, len(resp.Data))
}
