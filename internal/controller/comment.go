package controller

import (
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/internal/client"
	log "github.com/sirupsen/logrus"
)

func GetCommentByPost(id string) []*model.Comment {
	res, err := client.GetCommentByPost(id)
	if err != nil {
		log.Error("Get comment by post")
	}
	return res
}
func CreateComment(input *model.NewComment) string {
	res, err := client.CreateComment(input)
	if err != nil {
		log.Error("Create comment")
	}
	return res
}
