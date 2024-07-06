package user

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-model/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindOne(id string, user *model.User) error
	FindByEmail(email string, user *model.User) error
	Create(user *model.User, stamp *model.Stamp, group *model.Group) error
	Update(id string, user *model.User) error
	AssignGroup(id string, groupID *uuid.UUID) error
}

type repositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{Db: db}
}

func (r *repositoryImpl) FindOne(id string, user *model.User) error {
	return r.Db.Model(user).Preload("Stamp").First(user, "id = ?", id).Error
}

func (r *repositoryImpl) FindByEmail(email string, user *model.User) error {
	return r.Db.Model(user).Preload("Stamp").First(user, "email = ?", email).Error
}

func (r *repositoryImpl) Create(user *model.User, stamp *model.Stamp, group *model.Group) error {
	return r.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		stamp.UserID = &user.ID
		if err := tx.Create(stamp).Error; err != nil {
			return err
		}
		user.Stamp = stamp

		group.LeaderID = &user.ID
		if err := tx.Create(group).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *repositoryImpl) Update(id string, user *model.User) error {
	return r.Db.Model(user).Where("id = ?", id).Updates(user).Error
}

func (r *repositoryImpl) AssignGroup(id string, groupID *uuid.UUID) error {
	return r.Db.Model(&model.User{}).Where("id = ?", id).Update("group_id", groupID).Error
}
