package repository

import (
	"context"
	"github.com/Inf85/ticket-booking/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (t *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := t.db.Model(&models.Ticket{}).Preload("Event").Order("updated_at desc").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (t *TicketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := t.db.Model(ticket).Preload("Event").First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (t *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := t.db.Model(ticket).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return t.GetOne(ctx, ticket.ID)
}

func (t *TicketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := t.db.Model(ticket).Where("id = ?", ticketId).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return t.GetOne(ctx, ticket.ID)
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
