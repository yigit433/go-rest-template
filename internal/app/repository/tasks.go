package repository

import (
	"context"
	"time"

	"github.com/yigit433/go-rest-template/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) *TaskRepository {
	return &TaskRepository{collection: collection}
}

func (r *TaskRepository) FindAll(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) FindByID(ctx context.Context, id string) (*model.Task, error) {
	var task model.Task
	if err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Create(ctx context.Context, task *model.Task) (primitive.ObjectID, error) {
	task.CreatedAt = time.Now()

	inserted, err := r.collection.InsertOne(ctx, task)
	return inserted.InsertedID.(primitive.ObjectID), err
}

func (r *TaskRepository) Update(ctx context.Context, id string, task *model.Task) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": task})

	return err
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}