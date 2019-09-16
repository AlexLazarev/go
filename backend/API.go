package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	api API
)

func init() {
	api.address = "http://127.0.0.1:8082/api"
	api.user = api.address + "/user"
}

func (api *API)createUser(user user) error {
	buf, err := json.Marshal(user)

	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", api.user + "/create", bytes.NewReader(buf))
	if err != nil {
		return err
	}
	client := http.Client{}
	responce, err := client.Do(request)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		return err
	}
	return check(body)
}

func (api *API)getUser(login string, user *user) error {
	body, err := getQuery(api.user + "/findUser?login=" + login)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, user)
}

func (api *API)findLogin(login string) error {
	_, err := getQuery(api.user + "/findUser?login=" + login)
	return err
}

func (api *API)updateAge(id uint64, age uint64) error {
	_, err := getQuery(api.user + "/updateAge?id=" + strconv.FormatUint(id, 10) + "&" + "age=" + strconv.FormatUint(age, 10))
	return err
}

func getQuery(url string) ([]byte, error) {
	logf(url)
	responce, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		return nil, err
	}
	if err = check(body); err != nil {
		return nil, err
	}
	return body, nil
}


func check(body []byte) error {
	ret := obj{}
	if err := json.Unmarshal(body, &ret); err != nil {
		return err
	}
	_, exist := ret["err"]
	if exist {
		return errors.New(ret["err"].(string))
	}
	return nil
}
