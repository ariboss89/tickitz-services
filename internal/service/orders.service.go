package service

import (
	"context"
	"log"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

//var ErrInvalidGender = errors.New("invalid gender")

type OrderService struct {
	orderRepository *repository.OrdersRepository
	db              *pgxpool.Pool
}

func NewOrderService(orderRepository *repository.OrdersRepository, db *pgxpool.Pool) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		db:              db,
	}
}

func (o OrderService) GetSchedules(ctx context.Context, date, time, location string, movieId int) ([]dto.Schedules, error) {
	data, err := o.orderRepository.GetSchedule(ctx, o.db, date, time, location, movieId)
	if err != nil {
		return nil, err
	}
	var response []dto.Schedules

	for _, v := range data {
		response = append(response, dto.Schedules{
			Id:             v.Id,
			Show_Date:      v.Show_Date,
			Show_Time:      v.Show_Time,
			Price:          v.Price,
			Movie_Name:     v.Movie_Name,
			Cinema_Name:    v.Cinema_Name,
			Cinema_City:    v.Cinema_City,
			Cinema_Address: v.Cinema_Address,
			Studio_Name:    v.Studio_Name,
		})
	}
	return response, nil
}

func (o OrderService) CreateOrder(ctx context.Context, userId int, createOrder dto.CreateOrder) (dto.OrderResponse, error) {
	//begin trx
	tx, err := o.db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return dto.OrderResponse{}, err
	}

	data, err := o.orderRepository.CreateOrder(ctx, tx, userId, createOrder)
	if err != nil {
		return dto.OrderResponse{}, err
	}
	defer tx.Rollback(ctx)

	for i := range len(createOrder.Seats) {
		_, err := o.orderRepository.CreateDetailOrder(ctx, tx, data.Id, createOrder.Seats[i])
		if err != nil {
			return dto.OrderResponse{}, err
		}
	}

	if e := tx.Commit(ctx); e != nil {
		log.Println("failed to commit", e.Error())
		return dto.OrderResponse{}, e
	}

	response := dto.OrderResponse{
		Id: data.Id,
	}

	return response, nil
}
