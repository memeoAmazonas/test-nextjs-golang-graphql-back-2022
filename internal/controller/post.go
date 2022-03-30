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
func CreatePost(input *model.NewPost) int {
	res, err := client.CreatePost(input)
	if err != nil {
		log.Error("Create post")
	}
	return res
}
