package controller

import (
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/internal/client"
	log "github.com/sirupsen/logrus"
)

func GetPost() []*model.Post {
	res, err := client.GetPost()
	if err != nil {
		log.Error("Get list post")
	}
	return res
}
func GetCommentByPost(id string) []*model.Comment {
	res, err := client.GetCommentByPost(id)
	if err != nil {
		log.Error("Get comment by post")
	}
	return res
}
