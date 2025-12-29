package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"bytes"
	"fmt"
	"image/png"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	qrcode "github.com/skip2/go-qrcode"
)

func (h *studentHandlers) GeneratePDFWithQRCertificate(c *gin.Context) {
	token := c.Param("id")
	fmt.Println(token)
	if token == "" {
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		PDFContentError("ไม่พบ token certificate.", c)
		return
	}

	studentSuccessResponse, err := h.studentService.GetStudentSuccessCheck(token)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		PDFContentError(err.Error()+" โปรดติดต่อนักศึกษาให้สร้างเอกสารสารใหม่.", c)
		return
	}

	std_code := studentSuccessResponse.STD_CODE

	verifyURL := fmt.Sprintf("%s", std_code)

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

	// ตั้งค่าสำหรับ Watermark
	pdf.SetFont("THSarabun", "", 50)
	pdf.SetTextColor(200, 200, 200) // สีเทาอ่อน
	pdf.SetXY(30, 140)

	// บิดหมุนข้อความ 45 องศา
	pdf.TransformBegin()
	pdf.TransformRotate(45, 105, 148)
	pdf.Text(10, 150, "สำเนาเอกสารใช้เพื่อตรวจสอบเอกสารเท่านั้น") // หรือ "CONFIDENTIAL"
	pdf.TransformEnd()

	// คืนค่า Text และ Font ปกติ
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("THSarabun", "", 16)
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
	pdf.Cell(0, 6, fmt.Sprintf("ชื่อ สกุล: %s", studentSuccessResponse.NAME_THAI))
	pdf.SetXY(100, 50)
	pdf.Cell(0, 6, fmt.Sprintf("Name Surname: %s", studentSuccessResponse.NAME_ENG))
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
	
	// ตั้งค่าสำหรับ Watermark
	pdf.SetFont("THSarabun", "", 50)
	pdf.SetTextColor(200, 200, 200) // สีเทาอ่อน
	pdf.SetXY(30, 140)

	// บิดหมุนข้อความ 45 องศา
	pdf.TransformBegin()
	pdf.TransformRotate(45, 105, 148)
	pdf.Text(10, 150, "สำเนาเอกสารใช้เพื่อตรวจสอบเอกสารเท่านั้น") // หรือ "CONFIDENTIAL"
	pdf.TransformEnd()

	// คืนค่า Text และ Font ปกติ
	pdf.SetTextColor(0, 0, 0)

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

headers := []string{"ข้อมูลผู้สำเร็จการศึกษา (Graduate Information Inquiry)",}

rows := [][]string{
	{"1.รหัสประจำตัวนักศึกษา (Student Code)", studentSuccessResponse.STD_CODE},
	{"2.ชื่อ สกุล", studentSuccessResponse.NAME_THAI},
	{"  Name Surname", studentSuccessResponse.NAME_ENG},
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
colWidths := []float64{190} 


// Header
for i, header := range headers {
	pdf.SetFont("THSarabunBold", "", 14)
	pdf.CellFormat(colWidths[i], 14, header, "1", 0, "C", false, 0, "")
}
pdf.Ln(-1) // ขึ้นบรรทัดใหม่

colWidths = []float64{85, 105} 
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
pdf.MultiCell(0, 6, `เอกสารฉบับนี้ใช้ลายมือชื่ออิเล็กทรอนิกส์ตามพระราชบัญญัติว่าด้วยธุรกรรมทางอิเล็กทรอนิกส์ พ.ศ.2544 พระราชบัญญัตินี้กำหนดให้ลายมือชื่ออิเล็กทรอนิกส์ มีผลทางกฎหมายเทียบเท่ากับการลงลายมือชื่อ บนเอกสารราชการ`, "", "L", false)

pdf.SetXY(10, 250)
pdf.SetFont("THSarabunBold", "", 12)
pdf.MultiCell(0, 6, `This document is signed using an electronic signature in accordance with the Electronic Transactions Act B.E.2544 (2001), which recognizes electronic signatures as having the same legal effect as handwritten signatures on official documents.`, "", "L", false)


pdf.AddPage()
pdf.SetXY(10, 10)
pdf.SetFontSize(12)
pdf.MultiCell(0, 6, `หมายเหตุ
1.ระบบนี้จัดทำขึ้นเพื่อให้หน่วยงานภายนอกสามารถตรวจสอบคุณวุฒิการศึกษาของผู้สำเร็จการศึกษาจาก มหาวิทยาลัยรามคำแหง ระดับปริญญาโทและปริญญาเอก
2.หากต้องการตรวจสอบข้อมูลนอกเหนือจากที่ปรากฏ หรือมีปัญหา ข้อสงสัย โปรดติดต่อหน่วยตรวจสอบการสำเร็จการศึกษา ฝ่ายบริการการศึกษา บัณฑิตวิทยาลัย มหาวิทยาลัยรามคำแหง โทร.0-2310-8000 ต่อ 3708 หรือ 0-2310-8561 หรือ E-Mail: rugrad_verify@ru.ac.th`, "", "L", false)

pdf.SetXY(10, 35)
pdf.MultiCell(0, 6, `Note 
1.This system designed to allow external agencies to verify the education qualifications of graduates from Ramkhamhaeng University Master’s and Doctoarate level.
2.If you want to check information other than want is shown or have any questions of problem, please contact the Graduation Verification Unit, Educational Service Division, Graduate School, Ramkhamhaeng University Tel. 0-2310-8000 ext 3708 or 0-2310-8561 or E-Mail: rugrad_verify@ru.ac.th`, "", "L", false)

// 6. ส่ง PDF กลับไป
c.Header("Content-Type", "application/pdf")
c.Header("Content-Disposition", `attachment; filename=student_`+std_code+`.pdf`)
_ = pdf.Output(c.Writer)
}
