package controller

import (
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/internal/client"
	log "github.com/sirupsen/logrus"
)

func CreateLike(input *model.NewLike) string {
	res, err := client.CreateLike(input)
	if err != nil {
		log.Error("Create like")
	}
	return res
}
