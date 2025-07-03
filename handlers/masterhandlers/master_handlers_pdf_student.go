package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"errors"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"

	"bytes"
	"image/jpeg"
	"image/png"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	qrcode "github.com/skip2/go-qrcode"
)

func PDFContentError(strerr string, c *gin.Context) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("THSarabun", "", "fonts/THSarabunNew.ttf")
	pdf.AddUTF8Font("THSarabunBold", "", "fonts/THSarabunNew Bold.ttf")
	pdf.AddPage()
	// ตั้งค่าสำหรับ Watermark
	pdf.SetFont("THSarabun", "", 25)
	pdf.SetTextColor(200, 200, 200) // สีเทาอ่อน
	pdf.SetXY(30, 140)

	// บิดหมุนข้อความ 45 องศา
	pdf.TransformBegin()
	pdf.TransformRotate(45, 105, 148)
	pdf.Text(10, 150, strerr) // หรือ "CONFIDENTIAL"
	pdf.TransformEnd()

	// แทรกโลโก้ที่มุมบนซ้าย
	logoOpt := gofpdf.ImageOptions{
		ImageType:             "PNG",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}
	pdf.ImageOptions("images/logo.png", 10, 10, 15, 0, false, logoOpt, 0, "")

	pdf.SetFont("THSarabunBold", "", 20)
	pdf.SetXY(45, 10)
	pdf.Cell(0, 10, "แบบคำร้องขอรับรองคุณวุฒิการศึกษาแบบออนไลน์ ระดับบัณฑิตศึกษา")
	pdf.Ln(12)

	// 6. ส่ง PDF กลับไป
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename=student_error.pdf`)
	_ = pdf.Output(c.Writer)
}

func (h *studentHandlers) GeneratePicture(std_code string) (*bytes.Buffer, error) {

	token, err := middlewares.GenerateToken(std_code, "admin", h.redis_cache)

	if err != nil {
		return nil, err
	}

	service_token := viper.GetString("token.eservice")
	url := "http://10.2.1.155:9100/student/photograduate"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}

	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("id_token", token.AccessToken)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+service_token)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		err = errors.New("Invalid image format.")
		return nil, err
	}

	fmt.Println(contentType)

	// Decode the image
	var img image.Image
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			err = errors.New("Get decode jpeg error.")
			return nil, err
		}
	case "image/png":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			err = errors.New("Get decode png error.")
			return nil, err
		}
	default:
		err = errors.New("Unsupported image format.")
		return nil, err
	}

	// outputImg := new(bytes.Buffer)
	outputImg := bytes.NewBuffer(nil)

	if err := jpeg.Encode(outputImg, img, nil); err != nil {
		err = errors.New("Get resize error" + err.Error())
		return nil, err
	}

	return outputImg, nil
}

func (h *studentHandlers) GeneratePDFWithQR(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	fmt.Println(token)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		//c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ token login."})
		//c.Abort()
		PDFContentError("ไม่พบ token login.", c)
		return
	}

	fmt.Println(token)

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		err = errors.New("ไม่พบ claims user." + err.Error())
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		//c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ claims user."})
		//c.Abort()
		PDFContentError("ไม่พบ claims user.", c)
		return
	}

	role := claim.Role

	fmt.Println(role)

	if role == "Bachelor" {
		err = errors.New("สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้...")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		//c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		//c.Abort()
		PDFContentError("สิทธิ์ "+role+" ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้.", c)
		return
	}

	std_code := claim.StudentCode

	studentSuccessResponse, err := h.studentService.GetStudentSuccess(std_code)
	if err != nil {
		err = errors.New("ไม่พบข้อมูลรับรองคุณวุฒิการศึกษา " + std_code + ".")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลรับรองคุณวุฒิการศึกษา " + std_code + "."})
		//c.Abort()
		PDFContentError("ไม่พบข้อมูลรับรองคุณวุฒิการศึกษา "+std_code, c)
		return
	}

	tokenResponse, err := h.studentService.Certificate(token)
	if err != nil {
		c.Error(errors.New(err.Error() + ", " + token))
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		//c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		//c.Abort()
		PDFContentError("ไม่สามารถสร้าง Certificate ของ "+std_code, c)
		return
	}

	verifyURL := fmt.Sprintf("https://backend.ru.ac.th/egraduate/certificate/?id=%s", tokenResponse.CertificateToken)

	// 1. สร้าง QR Code เป็น image.Image
	qrImg, err := qrcode.New(verifyURL, qrcode.Medium)
	if err != nil {
		panic(err)
	}

	// 2. แปลงเป็น []byte ผ่าน io.Reader (PNG ในหน่วยความจำ)
	var buf bytes.Buffer
	err = png.Encode(&buf, qrImg.Image(400)) // 256 คือขนาด
	if err != nil {
		panic(err)
	}
	// ส่วนหัว
	// แทรกโลโก้ที่มุมบนซ้าย
	logoOpt := gofpdf.ImageOptions{
		ImageType:             "PNG",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}
	// 3. เตรียม PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("THSarabun", "", "fonts/THSarabunNew.ttf")
	pdf.AddUTF8Font("THSarabunBold", "", "fonts/THSarabunNew Bold.ttf")
	pdf.AddPage()
	pdf.SetFont("THSarabun", "", 16) // อย่าลืม register font ด้วย


	pdf.ImageOptions("images/logo.png", 10, 10, 15, 0, false, logoOpt, 0, "")

	pdf.SetFont("THSarabunBold", "", 20)
	pdf.SetXY(45, 10)
	pdf.Cell(0, 10, "แบบคำร้องขอรับรองคุณวุฒิการศึกษาแบบออนไลน์ ระดับบัณฑิตศึกษา")
	pdf.Ln(12)

	pdf.SetFontSize(18)
	pdf.SetXY(110, 25)
	pdf.CellFormat(0, 10, "ข้อมูลผู้สำเร็จการศึกษา (Graduate Information)", "", 1, "", false, 0, "")

	pdf.SetFont("THSarabunBold", "", 16)
	pdf.SetXY(100, 40)
	pdf.Cell(0, 6, fmt.Sprintf("ชื่อ - สกุล: %s", studentSuccessResponse.NAME_THAI))
	pdf.SetXY(100, 50)
	pdf.Cell(0, 6, fmt.Sprintf("Name - Surname: %s", studentSuccessResponse.NAME_ENG))
	pdf.SetXY(100, 60)
	pdf.Cell(0, 6, fmt.Sprintf("รหัสประจำตัวนักศึกษา: %s", studentSuccessResponse.STD_CODE))
	pdf.SetXY(100, 70)
	pdf.Cell(0, 6, fmt.Sprintf("Student Code: %s", studentSuccessResponse.STD_CODE))
	pdf.SetXY(100, 80)
	pdf.MultiCell(0, 6, fmt.Sprintf("ปริญญาที่สำเร็จการศึกษา: %s", studentSuccessResponse.THAI_NAME), "", "L", false)
	pdf.SetXY(100, 95)
	pdf.MultiCell(0, 6, fmt.Sprintf("Degree Awarded: %s", studentSuccessResponse.ENG_NAME), "", "L", false)

	// 4. Register image จาก memory
	opt := gofpdf.ImageOptions{
		ImageType:             "PNG",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}

	qrFile := fmt.Sprintf("tmp_qr_%s.png", std_code)

	pdf.RegisterImageOptionsReader(qrFile, opt, &buf)
	pdf.ImageOptions(qrFile, 10, 30, 80, 80, false, opt, 0, "")
	pdf.SetXY(27, 110)
	pdf.Cell(0, 8, "QR Code ข้อมูลนักศึกษา")
	pdf.SetXY(30, 115)
	pdf.SetFont("THSarabun", "", 12)
	pdf.Cell(0, 8, "(อายุการใช้งานไม่เกิน 120 วัน)")
	// คำอธิบายท้ายกระดาษ
	pdf.SetFont("THSarabun", "", 14)
	pdf.SetXY(10, 130)
	pdf.MultiCell(0, 6, `ผู้สำเร็จการศึกษา
โปรดนำส่งแบบตรวจสอบคุณวุฒิการศึกษาออนไลน์ระดับบัณฑิตศึกษามหาวิทยาลัยรามคำแหงฉบับนี้พร้อมสำเนาหนังสือสำคัญ
แสดงคุณวุฒิ (ใบรับรองสภามหาวิทยาลัยฯ, ใบปริญญาบัตรหรือใบรับรองผลการศึกษา(Transcript) ยังหน่วยงานภาครัฐ
หรือภาคเอกชนที่ท่านประสงค์ปรับคุณวุฒิการศึกษา สมัครงาน ศึกษาต่อ หรืออื่นๆ เพื่อให้หน่วยงานของท่านสามารถตรวจสอบ 
และขอหนังสือรับรองคุณวุฒิการศึกษาโดยตรงกับมหาวิทยาลัยด้วยความถูกต้องรวดเร็ว

หน่วยงานภาครัฐหรือภาคเอกชน 
สามารถตรวจสอบและขอหนังสือรับรองคุณวุฒิการศึกษาออนไลน์ระดับบัณฑิตศึกษา (ปริญญาโท-ปริญญาเอก) ของมหาวิทยาลัยรามคำแหง โดยดำเนินการ ดังนี้
1.Scan QR-code ในแบบตรวจสอบคุณวุฒิการศึกษาออนไลน์ของผู้สำเร็จการศึกษารายบุคคลข้างต้น 
2.กรอกข้อมูลที่เกี่ยวข้องในเว็บไซต์ ..................(E-Mail, ชื่อบริษัท/หน่วยงาน, ชื่อผู้รับผิดชอบในการตรวจสอบคุณวุฒิการศึกษา)
3.ระบบจะแสดงผลการตรวจสอบและรับรองคุณวุฒิการศึกษาซึ่งหน่วยงานสามารถสั่งพิมพ์หนังสือรับรองคุณวุฒิการศึกษาได้ผ่านทางเว็บไซต์ www.e-regis.ru.ac.th 

หมายเหตุ : 
1.ระบบนี้จัดทำขึ้นเพื่อให้หน่วยงานภายนอกสามารถตรวจสอบคุณวุฒิการศึกษาของผู้สำเร็จการศึกษาจากมหาวิทยาลัยรามคำแหง ระดับปริญญาโทและปริญญาเอก 
2.QR-code มีอายุการใช้งานไม่เกิน 120 วัน นับจากวันที่ออกหนังสือ
3.หากต้องการตรวจสอบข้อมูลนอกเหนือจากที่ปรากฏหรือมีปัญหาข้อสงสัยโปรดติดต่อหน่วยตรวจสอบการสำเร็จการศึกษา ฝ่ายบริการการศึกษา บัณฑิตวิทยาลัยมหาวิทยาลัยรามคำแหง โทร.0-2310-8000 ต่อ 3708 หรือ 0-2310-8561 หรือ E-Mail: rugrad_verify@ru.ac.th`, "", "L", false)

	pdf.AddPage()
	// ส่วนหัว
	// แทรกโลโก้ที่มุมบนซ้าย

	pdf.ImageOptions("images/logo.png", 10, 10, 15, 0, false, logoOpt, 0, "")
	pdf.SetFont("THSarabunBold", "", 16)
	pdf.SetXY(50, 10)
	pdf.Cell(0, 8, "รายงานผลการตรวจสอบและรับรองคุณวุฒิการศึกษา")
	pdf.SetXY(40, 15)
	pdf.Cell(0, 8, "(Report on the Educational Qualification and Certification)")
	pdf.SetXY(50, 20)
	pdf.Cell(0, 8, "มหาวิทยาลัยรามคำแหง (Ramkhamhaeng University)")
	pdf.SetXY(70, 25)
	pdf.Cell(0, 8, "ประเทศไทย (Thailand)")

	// ต้องตั้งชื่อให้ภาพแม้จะไม่ใช่ไฟล์ (ชื่อสมมติ)
	// ระบุว่าเป็นภาพ JPEG จาก memory
	imageOpts := gofpdf.ImageOptions{
		ImageType:             "JPEG",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}

	var imgReader io.Reader
	var imgName string
	mainImg, err := h.GeneratePicture(std_code)
	if err != nil {
		// แสดงรูปสำรองแทน
		fallbackImgFile, err2 := os.Open("images/person.jpg")
		if err2 != nil {
			panic("ไม่สามารถโหลด fallback image ได้")
		}
		defer fallbackImgFile.Close()
		imgReader = fallbackImgFile
		imgName = "person.jpg"
	} else {
		imgReader = mainImg
		imgName = "main.jpg"
	}

	pdf.RegisterImageOptionsReader(imgName, imageOpts, imgReader)
	pdf.ImageOptions(imgName, 160, 10, 30, 30, false, imageOpts, 0, "")

	gpa := fmt.Sprintf("%.2f", studentSuccessResponse.GPA)

	headers := []string{"ข้อมูลผู้สำเร็จการศึกษา (Graduate Information Inquiry)", "คำอธิบาย (DESCRIPTION)",}

	rows := [][]string{
		{"1.รหัสประจำตัวนักศึกษา (Student Code)", studentSuccessResponse.STD_CODE},
		{"2.ชื่อ-สกุล", studentSuccessResponse.NAME_THAI},
		{"  Name-Surname", studentSuccessResponse.NAME_ENG},
		{"3.วันที่เข้าศึกษา", studentSuccessResponse.ADMIT_DATE},
		{"  Date of Admission", studentSuccessResponse.ADMIT_DATE_EN},
		{"4.วันที่สำเร็จการศึกษา", studentSuccessResponse.GRADUATED_DATE},
		{"  Date of Graduation)",studentSuccessResponse.GRADUATED_DATE_EN},
		{"5.คุณวุฒิที่สำเร็จการศึกษา", studentSuccessResponse.CURR_NAME},
		{"  Degree Awarded", studentSuccessResponse.CURR_ENG},
		{"6.สาขาวิชา", studentSuccessResponse.MAJOR_NAME},
		{"  Field of Study", studentSuccessResponse.MAJOR_ENG},
		{"7.วิชาเอก", studentSuccessResponse.MAIN_MAJOR_THAI},
		{"  Major", studentSuccessResponse.MAIN_MAJOR_ENG},
		{"8.เกรดเฉลี่ยสะสม (GPA)", gpa},
	}

	pdf.SetXY(10, 45)
	pdf.SetFont("THSarabun", "", 14)
	// ความกว้างของแต่ละ column (หน่วย: mm)
	colWidths := []float64{85, 105} 

	// Header
	for i, header := range headers {
		pdf.SetFont("THSarabunBold", "", 14)
		pdf.CellFormat(colWidths[i], 14, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1) // ขึ้นบรรทัดใหม่

	// Rows
	for _, row := range rows {
		for i, col := range row {
			if i == 0 {
				pdf.SetFont("THSarabunBold", "", 14)
			} else {
				pdf.SetFont("THSarabun", "", 14)
			}
			pdf.CellFormat(colWidths[i], 10, col, "1", 0, "L", false, 0, "")
		}
		pdf.Ln(-1)
	}
	signOpt := gofpdf.ImageOptions{
		ImageType:             "JPG",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}
	pdf.SetFont("THSarabun", "", 14)
	pdf.ImageOptions("images/sign_long.jpg", 120, 205, 25, 0, false, signOpt, 0, "")
	pdf.SetXY(90, 205)
	pdf.Cell(0, 8, "ลงชื่อ (Signature)……………………………………………………ผู้รับรอง (Certifier)")
	pdf.SetXY(110, 210)
	pdf.Cell(0, 8, "( รองศาสตราจารย์กฤษดา ตั้งชัยศักดิ์ )")
	pdf.SetXY(110, 215)
	pdf.Cell(0, 8, "( Assoc.Prof. Krisda Tanchaisak )")
	pdf.SetXY(90, 220)
	pdf.Cell(0, 8, "คณบดีบัณฑิตวิทยาลัย ปฏิบัติราชการแทนอธิการบดีมหาวิทยาลัยรามคำแหง")
	pdf.SetXY(85, 225)
	pdf.Cell(0, 8, "Dean of Graduate School for the President of Ramkhamhaeng University")
	
	pdf.SetXY(10, 235)
	pdf.SetFont("THSarabunBold", "", 12)
	pdf.MultiCell(0, 6, `*** เอกสารฉบับนี้ใช้ลายมือชื่ออิเล็กทรอนิกส์ตามพระราชบัญญัติว่าด้วยธุรกรรมทางอิเล็กทรอนิกส์ พ.ศ.2544 พระราชบัญญัตินี้กำหนดให้ลายมือชื่ออิเล็กทรอนิกส์ มีผลทางกฎหมายเทียบเท่ากับการลงลายมือชื่อ บนเอกสารราชการ ***`, "", "L", false)
	
	pdf.SetXY(10, 248)
	pdf.SetFontSize(12)
	pdf.MultiCell(0, 6, `หมายเหตุ
1.ระบบนี้จัดทำขึ้นเพื่อให้หน่วยงานภายนอกสามารถตรวจสอบคุณวุฒิการศึกษาของผู้สำเร็จการศึกษาจาก มหาวิทยาลัยรามคำแหง ระดับปริญญาโทและปริญญาเอก
2.หากต้องการตรวจสอบข้อมูลนอกเหนือจากที่ปรากฏ หรือมีปัญหา ข้อสงสัย โปรดติดต่อหน่วยตรวจสอบการสำเร็จการศึกษา ฝ่ายบริการการศึกษา บัณฑิตวิทยาลัย มหาวิทยาลัยรามคำแหง โทร.0-2310-8000 ต่อ 3708 หรือ 0-2310-8561 หรือ E-Mail: rugrad_verify@ru.ac.th`, "", "L", false)
	
	pdf.AddPage()
	pdf.SetXY(10, 10)
	pdf.MultiCell(0, 6, `Note 
1.This system designed to allow external agencies to verify the education qualifications of graduates from Ramkhamhaeng University Master's and Doctoarate level
2.If you want to check information other than waht is shown or have any questions or problem, please contact the Graduation Verification Unit, Educational Service Division, Graduate School, Ramkhamhaeng University Tel. 02310-8000 ext 3708 or 0-2310-8561 or E-Mail: rugrad_verify@ru.ac.th`, "", "L", false)

	// 6. ส่ง PDF กลับไป
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename=student_`+std_code+`.pdf`)
	_ = pdf.Output(c.Writer)
}
