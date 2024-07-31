package repository

import (
	"context"
	"github.com/Inf85/ticket-booking/models"
	"gorm.io/gorm"
)

type EventRepositiry struct {
	db *gorm.DB
}

func (e *EventRepositiry) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}
	res := e.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (e *EventRepositiry) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {
	event := &models.Event{}

	res := e.db.Model(event).Where("id = ?", eventId).First(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (e *EventRepositiry) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := e.db.Model(event).Create(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (e *EventRepositiry) UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}

	updateRes := e.db.Model(event).Where("id = ?", eventId).Updates(updateData)
	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := e.db.Where("id = ?", eventId).First(event)
	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return event, nil
}

func (e *EventRepositiry) DeleteOne(ctx context.Context, eventId uint) error {
	res := e.db.Delete(&models.Event{}, eventId)

	return res.Error
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepositiry{
		db: db,
	}
}
