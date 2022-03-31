package client

import (
	"bytes"
	"encoding/json"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetPost() ([]*model.Post, error) {
	var response []*model.Post
	log.Info("get post list")
	res, err := http.Get("http://localhost:3002/post")

	if err != nil {
		log.Error("Get post list call", err.Error())
		return nil, err
	}
	log.Info("success call get post list")

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error("Parse data comments by post", err.Error())
		return nil, err
	}
	return response, nil
}
func CreatePost(input *model.NewPost) (int, error) {

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(input)
	req, _ := http.NewRequest("POST", "http://localhost:3002/post", buff)
	client := &http.Client{}
	log.Info("create new post")
	res, err := client.Do(req)
	if err != nil {
		log.Error("create new post")
		return -1, err
	}
	defer res.Body.Close()
	send, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("create new post parser body")
		return -1, err
	}
	log.Info("create new post success")
	re := strings.ReplaceAll(string(send), "\n", "")
	result, _ := strconv.Atoi(re)
	return result, err
}
