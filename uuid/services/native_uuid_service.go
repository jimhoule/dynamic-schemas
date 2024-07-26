package services

import "github.com/google/uuid"

type NativeUuidService struct{}

func (*NativeUuidService) Generate() string {
	return uuid.New().String()
}