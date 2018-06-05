package quiz

type TestEngine struct {

	Test *Test

	UserTestRepository *UserTestRepository

}

func (engine *TestEngine) StartTest(userId int) Test  {
	currentQuestionId := engine.Test.Questions[0].Id
	userTest := UserTest{UserId: userId, CurrentQuestionId: currentQuestionId}
	engine.UserTestRepository.save(&userTest)


	questions := make([]*Question, len(engine.Test.Questions))
	for _, question := range engine.Test.Questions {
		userQuestion := Question{
			Id:question.Id, Text:question.Text,
			Answers:make([]*Answer,len(question.Answers)),
		}
		for _, answer := range question.Answers {
			userQuestion.Answers = append(userQuestion.Answers, &Answer{
				Id:answer.Id,
				Text:answer.Text,
			})
		}
	}
	return Test{
		Name:engine.Test.Name,
		TimeLimit:engine.Test.TimeLimit,
		Questions: questions,
		Categories:engine.Test.Categories,
		}
}