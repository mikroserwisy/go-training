package quiz

import "github.com/jinzhu/gorm"

type UserTestRepository struct {

	Db *gorm.DB

}

func (repository *UserTestRepository) save(userTest *UserTest) {
	repository.Db.Create(userTest)
}

func (repository *UserTestRepository) getByUserId(userId int) *UserTest {
	userTest := UserTest{}
	repository.Db.First(&userTest, "user_Id = ?", userId)
	return &userTest
}

func (repository *UserTestRepository) update(userTest *UserTest)  {
	repository.Db.Save(userTest)
}