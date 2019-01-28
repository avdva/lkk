// Copyright 2019 Aleksandr Demakin. All rights reserved.

package lkk

import "net/http"

type Dictionary []DictEntry

type DictEntry struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Accuracy         int    `json:"accuracy"`
	InvertedAccuracy int    `json:"invertedAccuracy"`
	BaseAssetID      string `json:"baseAssetId"`
	QuotingAssetId   string `json:"quotingAssetId"`
}

type assetPairs struct {
	client *http.Client
}

func (a *assetPairs) Dictionary(market string) (*Dictionary, error) {
	var d Dictionary
	if err := get(a.client, requestPath+"AssetPairs/dictionary/"+market, &d); err != nil {
		return nil, err
	}
	return &d, nil
}
