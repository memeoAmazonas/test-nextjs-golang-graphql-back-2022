package client

import (
	"encoding/json"
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	log "github.com/sirupsen/logrus"
	"net/http"
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
