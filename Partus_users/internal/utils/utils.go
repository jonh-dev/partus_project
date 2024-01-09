package utils

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ValidAreaCodes = map[string]bool{
	"11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true,
	"21": true, "22": true, "24": true, "27": true, "28": true, "31": true, "32": true, "33": true, "34": true,
	"35": true, "37": true, "38": true, "41": true, "42": true, "43": true, "44": true, "45": true, "46": true,
	"47": true, "48": true, "49": true, "51": true, "53": true, "54": true, "55": true, "61": true, "62": true,
	"63": true, "64": true, "65": true, "66": true, "67": true, "68": true, "69": true, "71": true, "73": true,
	"74": true, "75": true, "77": true, "79": true, "81": true, "82": true, "83": true, "84": true, "85": true,
	"86": true, "87": true, "88": true, "89": true, "91": true, "92": true, "93": true, "94": true, "95": true,
	"96": true, "97": true, "98": true, "99": true,
}

func ConvertToObjectId(userId string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("Erro ao converter UserId para ObjectID: %v", err)
		return primitive.NilObjectID, status.Errorf(codes.Internal, "Erro ao converter UserId para ObjectID: %v", err)
	}

	return objectId, nil
}

func IsValidAreaCode(areaCode string) bool {
	_, ok := ValidAreaCodes[areaCode]
	return ok
}

func GetCurrentTimestamp() *timestamppb.Timestamp {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now().In(location)
	return timestamppb.New(now)
}

func AdjustToSaoPaulo(t *timestamppb.Timestamp) *timestamppb.Timestamp {
	adjustedTime := t.AsTime().Add(-3 * time.Hour)
	return timestamppb.New(adjustedTime)
}

func ReadjustToSaoPaulo(t time.Time) time.Time {
	return t.Add(3 * time.Hour)
}
