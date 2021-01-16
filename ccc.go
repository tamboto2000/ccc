package ccc

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// CardInfo contains card info
type CardInfo struct {
	BIN     string `json:"bin"`
	Country string `json:"country"`
	Vendor  string `json:"vendor"`
	Type    string `json:"type"`
	Level   string `json:"level"`
	Bank    string `json:"bank"`
}

const host = "https://ccbins.pro"

// Check check card info
func Check(bins string) (CardInfo, error) {
	return req(bins, nil)
}

// CheckWithProx check card info with proxy.
// proxy format {protoc}://{domain or ip}:{port}. Example:
// http://123.456.78.9:123
func CheckWithProx(bins, prox string) (CardInfo, error) {
	proxURL, err := url.Parse(prox)
	if err != nil {
		return CardInfo{}, err
	}

	return req(bins, proxURL)
}

func req(bins string, prox *url.URL) (CardInfo, error) {
	cinfo := CardInfo{}
	req, err := http.NewRequest("GET", host+"/?bins="+bins, nil)
	if err != nil {
		return cinfo, err
	}

	header := make(http.Header)
	header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	header.Set("Accept-Language", "en-US,en;q=0.5")
	header.Set("Connection", "keep-alive")
	header.Set("Referer", host+"/?bins="+bins)
	header.Set("Upgrade-Insecure-Requests", "1")

	req.Header = header

	cl := http.DefaultClient
	if prox != nil {
		cl.Transport = &http.Transport{Proxy: http.ProxyURL(prox)}
	}

	resp, err := cl.Do(req)
	if err != nil {
		return cinfo, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cinfo, err
	}

	// extract card info
	split := strings.Split(string(respBody), `<table style="">`)
	if len(split) == 0 {
		return cinfo, errors.New("unknown error")
	}

	table := strings.Split(split[1], `</table>`)[0]

	// extract BIN
	bin := strings.Split(table, `<tr><th>BIN:&nbsp;&nbsp;</th><td>`)
	bin = strings.Split(bin[1], `</td></tr>`)
	cinfo.BIN = bin[0]

	// extract country
	cn := strings.Split(table, `<tr><th>Country:&nbsp;&nbsp;</th><td>`)
	cn = strings.Split(cn[1], `</td></tr>`)
	cinfo.Country = cn[0]

	// extract vendor
	vn := strings.Split(table, `<tr><th>Vendor:&nbsp;&nbsp;</th><td>`)
	vn = strings.Split(vn[1], `</td></tr>`)
	cinfo.Vendor = vn[0]

	// extract type
	typ := strings.Split(table, `<tr><th>Type:&nbsp;&nbsp;</th><td>`)
	typ = strings.Split(typ[1], `</td></tr>`)
	cinfo.Type = typ[0]

	// extract level
	lvl := strings.Split(table, `<tr><th>Level:&nbsp;&nbsp;</th><td>`)
	lvl = strings.Split(lvl[1], `</td></tr>`)
	cinfo.Level = lvl[0]

	// extract bank
	bank := strings.Split(table, `<tr><th>Bank:&nbsp;&nbsp;</th><td>`)
	bank = strings.Split(bank[1], `</td></tr>`)
	cinfo.Bank = bank[0]

	return cinfo, nil
}
