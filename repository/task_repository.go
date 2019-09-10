package repository

import (
	"api/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type TaskRepository struct {

}

func (repo *TaskRepository) Store(ctx context.Context, task *domain.Task) error {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}

	errCon := client.Connect(ctx)

	if errCon != nil {
		return errCon
	}

	return nil
}

func (repo *TaskRepository) Load(ctx context.Context, id uint64) (*domain.Task, error) {
	return &domain.Task{}, nil
}

func (repo *TaskRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}

func (repo *TaskRepository) Update(ctx context.Context, task *domain.Task) error {
	return nil
}
