package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"errors"
	"fmt"

    "bytes"
	"github.com/gin-gonic/gin"
    "github.com/jung-kurt/gofpdf"
    qrcode "github.com/skip2/go-qrcode" 
    "image/png"
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
	pdf.Cell(0, 10, "แบบตรวจสอบคุณวุฒิการศึกษาออนไลน์ ระดับบัณฑิตศึกษา")
	pdf.Ln(12)

	    // 6. ส่ง PDF กลับไป
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename=student_error.pdf`)
	_ = pdf.Output(c.Writer)
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
		PDFContentError("สิทธิ์ " + role + " ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้.", c)
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
		PDFContentError("ไม่พบข้อมูลรับรองคุณวุฒิการศึกษา " + std_code , c)
		return
	}

	tokenResponse, err := h.studentService.Certificate(token) 
	if err != nil {
		c.Error(errors.New(err.Error() + ", " + token))
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		//c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		//c.Abort()
		PDFContentError("ไม่สามารถสร้าง Certificate ของ " + std_code, c)
		return
	}

	verifyURL := fmt.Sprintf("http://10.3.65.186:4300/certificate/?id=%s" , tokenResponse.CertificateToken)

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

    // 3. เตรียม PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("THSarabun", "", "fonts/THSarabunNew.ttf")
	pdf.AddUTF8Font("THSarabunBold", "", "fonts/THSarabunNew Bold.ttf")
	pdf.AddPage()
	pdf.SetFont("THSarabun", "", 16) // อย่าลืม register font ด้วย

	// ส่วนหัว
	// แทรกโลโก้ที่มุมบนซ้าย
	logoOpt := gofpdf.ImageOptions{
		ImageType:             "PNG",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}
	pdf.ImageOptions("images/logo.png", 10, 10, 15, 0, false, logoOpt, 0, "")

	pdf.SetFont("THSarabunBold", "", 20)
	pdf.SetXY(45, 10)
	pdf.Cell(0, 10, "แบบตรวจสอบคุณวุฒิการศึกษาออนไลน์ ระดับบัณฑิตศึกษา")
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
	pdf.SetXY(30, 110)
	pdf.Cell(0, 8, "QR-code ข้อมูลนักศึกษา")
	// คำอธิบายท้ายกระดาษ
	pdf.SetFont("THSarabun", "", 14)
	pdf.SetXY(10, 120)
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
	pdf.SetXY(72, 10)
	pdf.Cell(0, 8, "รายงานผลการตรวจสอบและรับรองคุณวุฒิการศึกษา")
	pdf.SetXY(60, 20)
	pdf.Cell(0, 8, "(Report on the Educational Qualification and Certification)")
	pdf.SetXY(70, 30)
	pdf.Cell(0, 8, "มหาวิทยาลัยรามคำแหง (Ramkhamhaeng University)")
	pdf.SetXY(90, 40)
	pdf.Cell(0, 8, "ประเทศไทย (Thailand)")

	gpa := fmt.Sprintf("%.2f", studentSuccessResponse.GPA)

	headers := []string{"ข้อมูลผู้สำเร็จการศึกษา Graduate Information Inquiry ", "THAI", "ENGLISH "}
	rows := [][]string{
        {"ชื่อ-สกุล (Name-Surname)", studentSuccessResponse.NAME_THAI, studentSuccessResponse.NAME_ENG},
        {"รหัสประจำตัวนักศึกษา (Student Code)", studentSuccessResponse.STD_CODE, studentSuccessResponse.STD_CODE},
        {"วันที่เข้าศึกษา (Date of Admission)", studentSuccessResponse.ADMIT_DATE, studentSuccessResponse.ADMIT_DATE},
        {"วันที่สำเร็จการศึกษา (Date of Graduation)", studentSuccessResponse.GRADUATED_DATE, studentSuccessResponse.GRADUATED_DATE},
        {"คุณวุฒิที่สำเร็จการศึกษา (Degree Awarded)", studentSuccessResponse.CURR_NAME, studentSuccessResponse.CURR_ENG},
        {"สาขาวิชา (Field of Study)", studentSuccessResponse.MAJOR_NAME , studentSuccessResponse.MAJOR_ENG},
        {"วิชาเอก (Major)", studentSuccessResponse.MAIN_MAJOR_THAI, studentSuccessResponse.MAIN_MAJOR_ENG},
        {"เกรดเฉลี่ยสะสม (GPA)", gpa, gpa},
    }

	pdf.SetXY(10, 60)
	pdf.SetFontSize(14)
	// ความกว้างของแต่ละ column (หน่วย: mm)
	colWidths := []float64{80, 55, 55}

	// Header
	for i, header := range headers {
		pdf.CellFormat(colWidths[i], 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1) // ขึ้นบรรทัดใหม่

	// Rows
	for _, row := range rows {
		for i, col := range row {
			pdf.CellFormat(colWidths[i], 10, col, "1", 0, "L", false, 0, "")
		}
		pdf.Ln(-1)
	}
	signOpt := gofpdf.ImageOptions{
		ImageType:             "JPG",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}
	pdf.ImageOptions("images/sign_long.jpg", 130, 158, 25, 0, false, signOpt, 0, "")
	pdf.SetXY(90, 160)
	pdf.Cell(0, 8, "ลงชื่อ (Signature)……………………………………………………ผู้รับรอง (Certifier)")
	pdf.SetXY(120, 170)
	pdf.Cell(0, 8, "( รศ.ดร.กฤษดา ตั้งชัยศักดิ์ )")
	pdf.SetXY(90, 180)
	pdf.Cell(0, 8, "ตำแหน่ง (Position) คณบดีบัณฑิตวิทยาลัย (Dean of Graduate School)")
	pdf.SetXY(10, 190)
	pdf.SetFontSize(12)
	pdf.MultiCell(0, 6, `*** เอกสารฉบับนี้ใช้ลายมือชื่ออิเล็กทรอนิกส์ตามพระราชบัญญัติว่าด้วยธุรกรรมทางอิเล็กทรอนิกส์ พ.ศ.2544 พระราชบัญญัตินี้กำหนดให้ ลายมือชื่ออิเล็กทรอนิกส์มีผลทางกฎหมายเทียบเท่ากับการลงลายมือชื่อ บนเอกสารราชการ`, "", "L", false)
	pdf.SetXY(10, 210)
	pdf.SetFontSize(12)
	pdf.MultiCell(0, 6, `หมายเหตุ
1.ระบบนี้จัดทำขึ้นเพื่อให้หน่วยงานภายนอกสามารถตรวจสอบคุณวุฒิการศึกษาของผู้สำเร็จการศึกษาจาก มหาวิทยาลัยรามคำแหง ระดับปริญญาโทและปริญญาเอก
2. หากต้องการตรวจสอบข้อมูลนอกเหนือจากที่ปรากฏ หรือมีปัญหา ข้อสงสัย โปรดติดต่อหน่วยตรวจสอบการสำเร็จการศึกษา ฝ่ายบริการการศึกษา บัณฑิตวิทยาลัย มหาวิทยาลัยรามคำแหง โทร.0-2310-8000 ต่อ 3708 หรือ 0-2310-8561 
หรือ E-Mail: rugrad_verify@ru.ac.th`, "", "L", false)
pdf.SetXY(10, 240)
pdf.MultiCell(0, 6, `Note 
1.This system designed to allow external agencies to verify the education qualifications of graduates from Ramkhamhaeng University Master's and Doctoarate level
2.If you want to check information other than waht is shown or have any questions or problem, please contact the Graduation Verification Unit, Educational Service Division, Graduate School, Ramkhamhaeng University Tel. 02310-8000 ext 3708 or 0-2310-8561 or E-Mail: rugrad_verify@ru.ac.th`, "", "L", false)

    // 6. ส่ง PDF กลับไป
    c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename=student_`+std_code+`.pdf`)
    _ = pdf.Output(c.Writer)
}