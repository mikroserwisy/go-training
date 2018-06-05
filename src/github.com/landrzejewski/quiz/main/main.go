package main

import (
	"github.com/landrzejewski/quiz"
	"encoding/json"
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
	db.LogMode(true)
	db.AutoMigrate(&quiz.UserTest{}, &quiz.UserAnswer{}, &quiz.UserAnswerValue{})

	repository := quiz.UserTestRepository{Db: db}

	engine := quiz.TestEngine{Test: test, UserTestRepository: &repository}
	engine.StartTest(1)

	engine.AnswerQuestion(1, &quiz.UserAnswer{QuestionId:1, Values:[]*quiz.UserAnswerValue{&quiz.UserAnswerValue{Value:"tetetet"}}})
}
