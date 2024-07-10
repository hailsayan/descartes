package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ExchangeResponse struct {
	Status string          `json:"status"`
	Stats  map[string]Stat `json:"stats"`
}

type Stat struct {
	BestSell  string `json:"best_sell"`
	IsClosed  bool   `json:"is_closed"`
	DayOpen   string `json:"day_open"`
	DayHigh   string `json:"day_high"`
	BestBuy   string `json:"best_buy"`
	VolumeSrc string `json:"volume_src"`
	DayLow    string `json:"day_low"`
	Latest    string `json:"latest"`
	VolumeDst string `json:"volume_dst"`
	DayChange string `json:"day_change"`
	DayClose  string `json:"day_close"`
}

func GetExchangeRate(src, dst string) (string, error) {
	if dst == "" {
		dst = "rls"
	} else {
		dst = strings.ToLower(dst)
	}
	src = strings.ToLower(src)

	resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:4001/rates?srcCurrency=%s&dstCurrency=%s", src, dst))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response ExchangeResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		return "", err
	}

	if response.Status != "OK" {
		return "", errors.New(fmt.Sprintf("bad status, expected ok got %s", response.Status))
	}

	return response.Stats[fmt.Sprintf("%s-%s", src, dst)].Latest, nil
}
