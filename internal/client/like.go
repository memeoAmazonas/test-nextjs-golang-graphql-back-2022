package client

import (
	"bytes"
	"encoding/json"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func CreateLike(input *model.NewLike) (string, error) {

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(input)
	req, _ := http.NewRequest("POST", "http://localhost:3002/like", buff)
	client := &http.Client{}
	log.Info("create new like")
	res, err := client.Do(req)
	if err != nil {
		log.Error("create like post")
		return "", err
	}

	defer res.Body.Close()
	send, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("create like")
		return "", err
	}
	log.Info("create like success")
	return string(send), err
}
