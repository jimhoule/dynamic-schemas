package uuid

import "main/uuid/services"

func GetService() services.UuidService {
	return &services.NativeUuidService{}
}