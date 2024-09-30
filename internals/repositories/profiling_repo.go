package repositories

import (
	"context"
	model "crud-hex/internals/core/domain"
	"crud-hex/internals/core/ports"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfilingRepository struct {
	collection *mongo.Collection
   }
   
   func NewProfilingRepository(db *mongo.Database) ports.IProfilingRepository {
	return &ProfilingRepository{
	 collection: db.Collection("profiling"),
	}
   }
   
   func (r *ProfilingRepository) Create(profiling model.Profiling) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
   
	_, err := r.collection.InsertOne(ctx, profiling)
	return err
   }