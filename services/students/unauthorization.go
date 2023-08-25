package students

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
)

func (s *studentServices) Unauthorization(token string) bool {

	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	isToken := middlewares.RevokeToken(token, s.redis_cache)
	if isToken {
		return isToken
	}

	return false
}

func (s *studentServices) CheckExistsToken(token string) bool {
	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	_, err := middlewares.CheckExistsToken(token, s.redis_cache)
	if err != nil {
		return false
	}

	return true
}
