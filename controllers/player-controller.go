package controllers

import (
	"github.com/dghubble/sling"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)


type Profile struct {
	Response struct {
				 Players []struct {
					 Steamid                  string `json:"steamid"`
					 Communityvisibilitystate int `json:"communityvisibilitystate"`
					 Profilestate             int `json:"profilestate"`
					 Personaname              string `json:"personaname"`
					 Lastlogoff               int `json:"lastlogoff"`
					 Profileurl               string `json:"profileurl"`
					 Avatar                   string `json:"avatar"`
					 Avatarmedium             string `json:"avatarmedium"`
					 Avatarfull               string `json:"avatarfull"`
					 Personastate             int `json:"personastate"`
					 Realname                 string `json:"realname"`
					 Primaryclanid            string `json:"primaryclanid"`
					 Timecreated              int `json:"timecreated"`
					 Personastateflags        int `json:"personastateflags"`
					 Loccountrycode           string `json:"loccountrycode"`
					 Locstatecode             string `json:"locstatecode"`
				 } `json:"players"`
			 } `json:"response"`
}

type (
	// PlayerController represents the controller for operating on the Player resource
	PlayerController struct {
	}
)

func NewPlayerController() *PlayerController {
	return &PlayerController{}
}

func (pc PlayerController) GetPlayerInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	accountId := p.ByName("accountId")
	profile := new(Profile)
	req, err := sling.New().Get("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=2029F8A592609ADA5E680CE08A2F128F&steamids=" + accountId).ReceiveSuccess(profile)
	if (err != nil) {
		w.WriteHeader(404)
		return
	} else {
		fmt.Println(req)
		// Marshal provided interface into JSON structure
		profile_to_json, _ := json.Marshal(profile)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", profile_to_json)
	}
}
