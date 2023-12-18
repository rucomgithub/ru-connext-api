package services_test

import (
	"RU-Smart-Workspace/ru-smart-api/databases"
	"RU-Smart-Workspace/ru-smart-api/repositories"
	"RU-Smart-Workspace/ru-smart-api/services"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestGetRegister(t *testing.T) {
	t.Run("Get Register Success", func(t *testing.T) {
		//Arrage
		var redis_cache *redis.Client
		redis_cache = databases.NewDatabases().RedisInint()
		promoRepo := repositories.NewRegisterRepoMock()
		registers := []repositories.RegisterRepo{}
		registers = append(registers, repositories.RegisterRepo{
			STD_CODE:  "6299999991",
			YEAR:      "2565",
			SEMESTER:  "1",
			COURSE_NO: "ACC1101",
			CREDIT:    "3",
		})
		registerRecord := []services.RegisterRecord{}
		registerRecord = append(registerRecord, services.RegisterRecord{
			YEAR:      "2565",
			SEMESTER:  "1",
			COURSE_NO: "ACC1101",
			CREDIT:    "3",
		})
		expected := services.RegisterResponse{
			STD_CODE: "6299999991",
			YEAR:     "2565",
			RECORD:   registerRecord,
		}
		promoRepo.On("GetRegister", "6299999991", "2565").Return(&registers, nil)

		registerService := services.NewRegisterServices(promoRepo, redis_cache)

		registerReq := services.RegisterRequest{
			STD_CODE: "6299999991",
			YEAR:     "2565",
		}

		//Act
		result, _ := registerService.GetRegister(registerReq)

		//Assert
		assert.Equal(t, &expected, result)
		promoRepo.AssertNotCalled(t, "GetRegister")
	})
}
