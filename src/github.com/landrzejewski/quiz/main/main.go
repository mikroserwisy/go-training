package main

import (
	"github.com/landrzejewski/quiz"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/landrzejewski/quiz/common"
)

func parseData(bytes []byte) *quiz.Test {
	var test quiz.Test
	common.PanicIfError(json.Unmarshal(bytes, &test))
	return &test
}

func main() {
	data := common.LoadData("test.json")
	test := parseData(data)


	db := common.InitDb()
	db.AutoMigrate(&quiz.UserTest{}, &quiz.UserAnswer{}, &quiz.UserAnswerValue{})

	repository := quiz.UserTestRepository{Db: db}
	engine := quiz.TestEngine{Test: test, UserTestRepository: &repository}

	spew.Dump(engine.StartTest(1))
}
