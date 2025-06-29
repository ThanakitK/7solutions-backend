package repositories

import (
	"7solutions/backend/core/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	db         *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) UserRepository {
	return &userRepo{
		db:         db,
		collection: collection,
	}
}

func (r *userRepo) CreateUser(payload models.RepoCreateUserModel) (result models.RepoResUserModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	after := options.After
	statusTrue := true
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &statusTrue,
	}

	res := r.db.Collection(r.collection).FindOneAndUpdate(ctx, bson.M{"user_id": payload.ID}, bson.M{"$set": payload}, &opt)
	if res.Err() != nil {
		return result, res.Err()
	}

	if err := res.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (r *userRepo) GetUserByID(id string) (result models.RepoResUserModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := r.db.Collection(r.collection).FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return result, res.Err()
	}

	if err := res.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (r *userRepo) GetUserByEmail(email string) (result models.RepoResUserModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := r.db.Collection(r.collection).FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		return result, res.Err()
	}

	if err := res.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (r *userRepo) GetUsers() (result []models.RepoResUserModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user models.RepoResUserModel
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

func (r *userRepo) UpdateUser(id string, payload models.RepoUpdateUserModel) (result models.RepoResUserModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": payload})
	if err != nil {
		return result, err
	}

	res := r.db.Collection(r.collection).FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return result, res.Err()
	}

	if err := res.Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

func (r *userRepo) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) CountUser() (result int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.db.Collection(r.collection).CountDocuments(ctx, bson.M{})
	if err != nil {
		return result, err
	}
	result = res
	return result, nil
}
