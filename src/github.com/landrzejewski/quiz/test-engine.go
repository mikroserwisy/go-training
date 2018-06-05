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
	if userTest.ID != 0 && userTest.Finished == false {
		if engine.isTimeLeft(userTest) {
			userTest.Finished = true
		} else {
			answerIndex := engine.indexOf(userTest, answer)
			if answerIndex == -1 {
				userTest.Answers = append(userTest.Answers, answer)
			} else {
				userTest.Answers[answerIndex] = answer
			}
		}
		engine.UserTestRepository.update(userTest)
	}
}

func (engine *TestEngine) isTimeLeft(userTest *UserTest) bool {
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