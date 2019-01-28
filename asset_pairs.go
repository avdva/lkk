// Copyright 2019 Aleksandr Demakin. All rights reserved.

package lkk

import (
	"net/http"

	"github.com/shopspring/decimal"
)

// Dictionary is an array if assets on some market.
type Dictionary []DictEntry

// DictEntry is a description of an asset.
type DictEntry struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Accuracy         int    `json:"accuracy"`
	InvertedAccuracy int    `json:"invertedAccuracy"`
	BaseAssetID      string `json:"baseAssetId"`
	QuotingAssetId   string `json:"quotingAssetId"`
}

// RateEntry is an ask/bid pair for some asset.
type RateEntry struct {
	ID  string          `json:"id"`
	Ask decimal.Decimal `json:"ask"`
	Bid decimal.Decimal `json:"bid"`
}

type assetPairs struct {
	client *http.Client
}

// Dictionary returns discionary for given market.
func (a *assetPairs) Dictionary(market string) (*Dictionary, error) {
	var d Dictionary
	if err := get(a.client, requestPath+"AssetPairs/dictionary/"+market, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// Rates returns current asset rates for all available pairs.
func (a *assetPairs) Rates() ([]RateEntry, error) {
	var rates []RateEntry
	if err := get(a.client, requestPath+"AssetPairs/rate", &rates); err != nil {
		return nil, err
	}
	return rates, nil
}

// Rates returns current asset rate for given pair.
func (a *assetPairs) Rate(assetPair string) (*RateEntry, error) {
	var rate RateEntry
	path := requestPath + "AssetPairs/rate/" + assetPair
	if err := get(a.client, path, &rate); err != nil {
		return nil, err
	}
	return &rate, nil
}
