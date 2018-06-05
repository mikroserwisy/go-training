package quiz

type TestEngine struct {

	Test Test

	UserTestRepository *UserTestRepository

}

func (engine *TestEngine) StartTest(userId int) Test  {
	currentQuestionId := engine.Test.Questions[0].Id
	userTest := UserTest{UserId: userId, CurrentQuestionId: currentQuestionId}
	engine.UserTestRepository.save(&userTest)
	testCopy := engine.Test
	for _, question := range testCopy.Questions {
		for _, answer := range question.Answers {
			answer.Value = ""
		}
	}
	return testCopy
}