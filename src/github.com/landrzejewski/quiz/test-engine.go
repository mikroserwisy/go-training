package quiz

import "time"

type TestEngine struct {

	Test *Test

	UserTestRepository *UserTestRepository

}

func (engine *TestEngine) StartTest(userId int) *Test  {
	engine.prepareUserTest(userId)
	return engine.buildTestDefinition()
}

func (engine *TestEngine) prepareUserTest(userId int)  {
	userTest := engine.UserTestRepository.getByUserId(userId)
	if userTest.ID == 0 {
		currentQuestionId := engine.Test.Questions[0].Id
		userTest := UserTest{UserId: userId, CurrentQuestionId: currentQuestionId}
		engine.UserTestRepository.save(&userTest)
	}
}

func (engine *TestEngine) buildTestDefinition() *Test {
	userQuestions := make([]*Question, 0)
	for _, question := range engine.Test.Questions {
		userQuestion := Question{Id:question.Id, Text:question.Text, Answers:make([]*Answer,0),
		}
		for _, answer := range question.Answers {
			userQuestion.Answers = append(userQuestion.Answers, &Answer{Id:answer.Id, Text:answer.Text})
		}
		userQuestions = append(userQuestions, &userQuestion)
	}
	return &Test{Name:engine.Test.Name, TimeLimit:engine.Test.TimeLimit, Questions: userQuestions, Categories:engine.Test.Categories}
}

func (engine *TestEngine) AnswerQuestion(userId int, answer *UserAnswer) {
	userTest := engine.UserTestRepository.getByUserId(userId)
	if engine.isActive(userTest) {
		if engine.shouldEnd(userTest) {
			userTest.Finished = true
		} else {
			engine.addAnswer(userTest, answer)
		}
		engine.UserTestRepository.update(userTest)
	}
}

func (engine *TestEngine) isActive(userTest *UserTest) bool{
	return userTest.ID != 0 && userTest.Finished == false
}

func (engine *TestEngine) addAnswer(userTest *UserTest, answer *UserAnswer) {
	var question *Question
	for _, currentQuestion := range engine.Test.Questions {
		if currentQuestion.Id == answer.QuestionId {
			question = currentQuestion
			break
		}
	}
	if question != nil {
		valid := true
		for index := range question.Answers {
			if question.Answers[index].Value != answer.Values[index].Value {
				valid = false
				break
			}
		}
		answer.Valid = valid
		answerIndex := engine.indexOf(userTest, answer)
		if answerIndex == -1 {
			userTest.Answers = append(userTest.Answers, answer)
		} else {
			userTest.Answers[answerIndex] = answer
		}
	}
}

func (engine *TestEngine) shouldEnd(userTest *UserTest) bool {
	return int(time.Since(userTest.CreatedAt).Seconds()) - engine.Test.TimeLimit >= 0
}

func (engine *TestEngine) indexOf(userTest *UserTest, answer *UserAnswer) int {
	answerIndex := -1
	for index, currentAnswer := range userTest.Answers {
		if currentAnswer.QuestionId == answer.QuestionId {
			answerIndex = index
			break
		}
	}
	return answerIndex
}

func (engine *TestEngine) generateReport(userId int) {

}