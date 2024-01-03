package utils

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertToObjectId(userId string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("Erro ao converter UserId para ObjectID: %v", err)
		return primitive.NilObjectID, status.Errorf(codes.Internal, "Erro ao converter UserId para ObjectID: %v", err)
	}

	return objectId, nil
}
