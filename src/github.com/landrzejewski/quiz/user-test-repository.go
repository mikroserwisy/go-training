package quiz

import "github.com/jinzhu/gorm"

type UserTestRepository struct {

	Db *gorm.DB

}

func (repository *UserTestRepository) save(userTest *UserTest) {
	repository.Db.Create(userTest)
}