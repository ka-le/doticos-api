package main

import (
	"bufio"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

const playerSummariesURL = "http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/"

func getPlayerInfoHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	steamReq, err := http.NewRequest(http.MethodGet, playerSummariesURL, nil)
	if err != nil {
		//TODO improve error handling
		http.Error(w, err.Error(), 500)
	}

	query := url.Values{}
	query.Add("key", "2029F8A592609ADA5E680CE08A2F128F")
	query.Add("steamids", p.ByName("accountId"))

	steamReq.URL.RawQuery = query.Encode()

	steamResp, err := http.DefaultClient.Do(steamReq)
	if err != nil {
		//TODO improve error handling
		w.WriteHeader(404)
	}

	w.Header().Add("Content-Type", "application/json")
	bufio.NewReader(steamResp.Body).WriteTo(w)
}
