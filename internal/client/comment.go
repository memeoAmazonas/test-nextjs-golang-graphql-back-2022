package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func GetCommentByPost(id string) ([]*model.Comment, error) {
	var response []*model.Comment
	log.Info("get comments by post client")
	res, err := http.Get(fmt.Sprintf("http://localhost:3002/post/%s/comments", id))

	if err != nil {
		log.Error("Get comments by post call", err.Error())
		return nil, err
	}
	log.Info("success call get comments by post")

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error("Parse data comments by post", err.Error())
		return nil, err
	}
	return response, nil

}
func CreateComment(input *model.NewComment) (string, error) {

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(input)
	req, _ := http.NewRequest("POST", "http://localhost:3002/comment", buff)
	client := &http.Client{}
	log.Info("create comment")
	res, err := client.Do(req)
	if err != nil {
		log.Error("create comments get")
		return "", err
	}
	defer res.Body.Close()
	send, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("create comments parser body")
		return "", err
	}
	log.Info("create comment success")
	return string(send), err
}
