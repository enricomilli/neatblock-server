package viabtc

import (
	"fmt"
	"net/http"
	"net/url"
)

// independent pool example: https://www.viabtc.com/observer/dashboard?access_key=d1cab8ae8c6fedcc2cd3d370cc1fd212&coin=BTC
// multiple user pool example: https://www.viabtc.com/observer/dashboard?access_key=fec491fdb7fdc2726c357200a798a041&coin=BTC&user_id=172951

func (provider *ViaBTC) CompanyName() string {
	return "VIABTC"
}

func (provider *ViaBTC) ValidateURL(observerURL string) error {

	url, err := url.Parse(observerURL)
	if err != nil {
		return fmt.Errorf("This URL could not be parsed, please try a different one.")
	}
	queries := url.Query()

	accessKey := queries.Get("access_key")
	if accessKey == "" {
		return fmt.Errorf("No access key found in this URL, please paste the complete URL.")
	}

	coin := queries.Get("coin")
	if coin == "" {
		return fmt.Errorf("No coin parameter found in this URL, please paste the complete URL.")
	}

	return nil
}

func (provider *ViaBTC) GetTotalsEndpoint(userID, accessKey, coin string) string {
	queryString := "?" + "access_key=" + accessKey + "&coin=" + coin

	if userID != "" {
		queryString += "&user_id=" + userID
	}

	return "https://www.viabtc.com/res/observer/home" + queryString
}

func (provider *ViaBTC) GetRewardsEndpoint(userID, accessKey, coin string) string {
	queryString := "?" + "access_key=" + accessKey + "&coin=" + coin

	if userID != "" {
		queryString += "&user_id=" + userID
	}

	return "https://www.viabtc.com/observer/profit" + queryString
}

func addViaBTCHeaders(req *http.Request) {
	req.Header.Set("authority", "www.viabtc.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en_US")
	req.Header.Set("cookie", "lang=en_US")
	req.Header.Set("platform", "web")
	req.Header.Set("sec-ch-ua", "'Not_A Brand';v='8', 'Chromium';v='120', 'Brave';v='120'")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "'macOS'")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
}
