package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type PushId string
type ConfigurationKey string

type newPushIdPostPayload struct {
	PushId   PushId `json:"pushId"`
	UserId   UserId `json:"userId"`
	Username Username `json:"username"`
}

type authenticationAcceptPostPayload struct {
	Answer   string `json:"answer"`
	PushId   PushId `json:"pushId"`
	Code     ConfigurationKey `json:"code"`
	Username Username `json:"username"`
	UserId   UserId `json:"userId"`
}

type authenticationRejectPostPayload struct {
	Answer   string `json:"answer"`
	PushId   PushId `json:"pushId"`
	Username Username `json:"username"`
	UserId   UserId `json:"userId"`
}

func postAuthenticationAccept(url *url.URL, pushId PushId, confKey ConfigurationKey, username Username, userId UserId) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(authenticationAcceptPostPayload{
		Answer:   "yes",
		PushId:   pushId,
		Code:     confKey,
		Username: username,
		UserId:   userId,
	})
	resp, _ := http.Post(url.String()+SUFFIX_INBOUNDCLIENT, APPLICATION_TYPE_JSON, b)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Print(string(body))
}

func postAuthenticationReject(url *url.URL, pushId PushId, username Username, userId UserId) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(authenticationRejectPostPayload{
		Answer:   "no",
		PushId:   pushId,
		Username: username,
		UserId:   userId,
	})
	resp, _ := http.Post(url.String()+SUFFIX_INBOUNDCLIENT, APPLICATION_TYPE_JSON, b)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Print(string(body))
}

func postNewPushId(url *url.URL, pushId PushId, username Username, userId UserId) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(newPushIdPostPayload{
		PushId:   pushId,
		UserId:   userId,
		Username: username,
	})

	resp, _ := http.Post(url.String()+SUFFIX_INBOUNDCLIENT+SUFFIX_UPDATEPUSHID, APPLICATION_TYPE_JSON, b)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Print(string(body))
}
