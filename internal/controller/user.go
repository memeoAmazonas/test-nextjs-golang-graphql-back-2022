package controller

import (
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/internal/client"
	log "github.com/sirupsen/logrus"
)

func CreateUser(input *model.NewUser) int {
	res, err := client.CreateUser(input)
	if err != nil {
		log.Error("Create user")
	}
	return res
}
func GetUserByEmail(email string) *model.User {
	res, err := client.GetUserByEmail(email)
	if err != nil {
		log.Error("Get user")
	}
	return res
}
