package masterservice

import "fmt"

func (s *studentServices) SetPrivacyPolicy(privacyPolicyRequest PrivacyPolicyRequest) (*PrivacyPolicyResponse, error) {

	privacy := PrivacyPolicyResponse{}

	_, err := s.studentRepo.GetPrivacyPolicy(privacyPolicyRequest.STD_CODE)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			fmt.Println("ไม่พบข้อมูล ทำการเพิ่มข้อมูลใหม่")
			err := s.studentRepo.AddPrivacyPolicy(privacyPolicyRequest.STD_CODE, privacyPolicyRequest.VERSION)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	fmt.Println("พบข้อมูล ทำการปรับค่าข้อมูล Version ใหม่")
	err = s.studentRepo.UpdatePrivacyPolicy(privacyPolicyRequest.STD_CODE, privacyPolicyRequest.VERSION)
	if err != nil {
		return nil, err
	}

	sp, err := s.studentRepo.GetPrivacyPolicy(privacyPolicyRequest.STD_CODE)

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
