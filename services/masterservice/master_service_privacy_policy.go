package masterservice

import "fmt"

func (s *studentServices) SetPrivacyPolicy(std_code, version string) (*PrivacyPolicyResponse, error) {

	privacy := PrivacyPolicyResponse{}

	_, err := s.studentRepo.GetPrivacyPolicy(std_code)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			fmt.Println("ไม่พบข้อมูล " + std_code + " ทำการเพิ่มข้อมูลใหม่")
			err := s.studentRepo.AddPrivacyPolicy(std_code, version)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	fmt.Println("พบข้อมูล ทำการปรับค่าข้อมูล version ใหม่")
	err = s.studentRepo.UpdatePrivacyPolicy(std_code, version)
	if err != nil {
		return nil, err
	}

	sp, err := s.studentRepo.GetPrivacyPolicy(std_code)

	if err != nil {
		return nil, err
	}

	privacy = PrivacyPolicyResponse{
		STD_CODE: sp.STD_CODE,
		VERSION:  sp.VERSION,
		CREATED:  sp.CREATED,
		MODIFIED: sp.MODIFIED,
	}

	return &privacy, nil
}
