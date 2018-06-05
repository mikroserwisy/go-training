package quiz

import "github.com/jinzhu/gorm"

type Test struct {

	Name string

	TimeLimit int

	Questions []*Question

	Categories []*Category

}

type Category struct {

	Id int

	Name string

}

type Question struct {

	Id int

	Text string

	Answers []*Answer

	CategoryId int

}

type Answer struct {

	Id int

	Text string

	Value string

}

//---------------------------------------------------------

type UserTest struct {

	gorm.Model

	Finished bool

	UserId int

	CurrentQuestionId int

	Answers []*UserAnswer `gorm:"foreignkey:user_test_id"`

}

type UserAnswer struct {

	gorm.Model

	QuestionId int

	Values []*UserAnswerValue `gorm:"foreignkey:user_answer_value_id"`

	UserTestID uint

}

type UserAnswerValue struct {

	gorm.Model

	Value string

	UserAnswerValueID uint

}