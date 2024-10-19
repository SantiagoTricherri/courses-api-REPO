package dao

import (
	"context"
	"errors"
	domain "inscriptions-api/domain/inscriptions"

	"gorm.io/gorm"
)

type InscriptionDAO struct {
	db *gorm.DB
}

type InscriptionModel struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserID   uint `gorm:"not null;index"`
	CourseID uint `gorm:"not null;index"`
}

func NewInscriptionDAO(db *gorm.DB) *InscriptionDAO {
	return &InscriptionDAO{db: db}
}
func (dao *InscriptionDAO) CreateInscription(ctx context.Context, userID, courseID uint) (*domain.Inscription, error) {
	// Verificar si la inscripción ya existe
	var inscription InscriptionModel
	if err := dao.db.WithContext(ctx).Where("user_id = ? AND course_id = ?", userID, courseID).
		First(&inscription).Error; err == nil {
		return nil, errors.New("inscription already exists")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// Crear la inscripción
	newInscription := InscriptionModel{UserID: userID, CourseID: courseID}
	if err := dao.db.WithContext(ctx).Create(&newInscription).Error; err != nil {
		return nil, err
	}

	return &domain.Inscription{
		ID:       newInscription.ID,
		UserID:   newInscription.UserID,
		CourseID: newInscription.CourseID,
	}, nil
}

func (dao *InscriptionDAO) GetInscriptions(ctx context.Context) ([]domain.Inscription, error) {
	var inscriptionsModel []InscriptionModel
	if err := dao.db.WithContext(ctx).Find(&inscriptionsModel).Error; err != nil {
		return nil, err
	}

	inscriptions := make([]domain.Inscription, len(inscriptionsModel))
	for i, model := range inscriptionsModel {
		inscriptions[i] = domain.Inscription{
			ID:       model.ID,
			UserID:   model.UserID,
			CourseID: model.CourseID,
		}
	}
	return inscriptions, nil
}

func (dao *InscriptionDAO) GetInscriptionsByUser(ctx context.Context, userID uint) ([]domain.Inscription, error) {
	var inscriptionsModel []InscriptionModel
	if err := dao.db.WithContext(ctx).Where("user_id = ?", userID).Find(&inscriptionsModel).Error; err != nil {
		return nil, err
	}

	inscriptions := make([]domain.Inscription, len(inscriptionsModel))
	for i, model := range inscriptionsModel {
		inscriptions[i] = domain.Inscription{
			ID:       model.ID,
			UserID:   model.UserID,
			CourseID: model.CourseID,
		}
	}
	return inscriptions, nil
}
