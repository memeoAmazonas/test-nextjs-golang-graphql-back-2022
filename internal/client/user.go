package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func CreateUser(input *model.NewUser) (int, error) {

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(input)
	req, _ := http.NewRequest("POST", "http://localhost:3002/user", buff)
	client := &http.Client{}
	log.Info("create user")
	res, err := client.Do(req)
	if err != nil {
		log.Error("create user")
		return -1, err
	}
	defer res.Body.Close()
	send, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("create user parser body")
		return -1, err
	}
	log.Info("create user success")
	re := strings.ReplaceAll(string(send), "\n", "")
	result, _ := strconv.Atoi(re)
	return result, err
}
func GetUserByEmail(email string) (*model.User, error) {
	var response *model.User
	log.Info("get user")
	res, err := http.Get(fmt.Sprintf("http://localhost:3002/user/%s", email))

	if err != nil {
		log.Error("Get user", err.Error())
		return nil, err
	}
	log.Info("success get user")

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error("Parse get user", err.Error())
		return nil, err
	}
	return response, nil
}
