package masterhandlers

import (
	"fmt"
    "bytes"
	"github.com/gin-gonic/gin"
    "github.com/jung-kurt/gofpdf"
    qrcode "github.com/skip2/go-qrcode" 
    "image/png"
)

func (h *studentHandlers) GeneratePDFWithQR(c *gin.Context) {
	studentID := c.Param("id")
	studentName := "นายสมชาย ใจดี"
	studentCode := "6512345678"
	degree := "วิทยาศาสตรบัณฑิต (คอมพิวเตอร์)"

	verifyURL := fmt.Sprintf("http://ruconnext-dev.ru.ac.th:9100/master/student/successcheck/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW4iOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKaFkyTmxjM05mZEc5clpXNWZhMlY1SWpvaU5UTXhNamMwTURJd09EbzZZV05qWlhOek9qb3laV0V6WkdSbU9DMWpNMkV5TFRReFl6Z3RZamhqTkMwek5HSTVZalJtTmpNNE5EUWlMQ0psZUhCcGNtVnpYM1J2YTJWdUlqb3hOelE1TnpBeU5ERTRMQ0pwYzNOMVpYSWlPaUpTZFMxamIyNXVaWGgwSWl3aWNtVm1jbVZ6YUY5MGIydGxibDlyWlhraU9pSTFNekV5TnpRd01qQTRPanB5WldaeVpYTm9Pam8yTmpJeVlXRmxZeTFrT1dJd0xUUTVNR1l0WVdVeE1TMDBObVpsTjJaaE16QXdZVEFpTENKeWIyeGxJam9pVFdGemRHVnlJaXdpYzNSa1gyTnZaR1VpT2lJMU16RXlOelF3TWpBNElpd2ljM1ZpYW1WamRDSTZJbEoxTFVOdmJtNWxlSFE2T2pVek1USTNOREF5TURnaWZRLkNyUVJqelZKM2tEdkNrRGhYcmZkb0U1RllwN29JUm56bDNJV2xnem9TdHMiLCJhY2Nlc3NfdG9rZW5fa2V5IjoiNTMxMjc0MDIwODo6Y2VydGlmaWNhdGU6OmM0YzJjMjdlLTEwMTctNGFhMC1iM2Q4LTFkNjliMTYwMjk5ZSIsImNlcnRpZmljYXRlIjoiZWdyYWR1YXRlIiwiZXhwaXJlX2RhdGUiOiIyMDI1LTA2LTExVDEyOjM5OjMxLjQxNTM5Njc3NyswNzowMCIsImV4cGlyZXNfdG9rZW4iOjE3NDk2MjAzNzEsImlzc3VlciI6IlJ1LWNvbm5leHRlZ3JhZHVhdGUiLCJzdGFydF9kYXRlIjoiMjAyNS0wNi0xMVQxMTozOTozMS40MTUzOTY3NzcrMDc6MDAiLCJzdGRfY29kZSI6IjUzMTI3NDAyMDgiLCJzdWJqZWN0IjoiQ2VydGlmaWNhdGU6OjUzMTI3NDAyMDgifQ.PA2uzyStdCGVHu0A4rHsRrejm4ndfy-7Vd_xNQREaOU")

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
	pdf.AddPage()
	pdf.SetFont("THSarabun", "", 14) // อย่าลืม register font ด้วย

	// ส่วนหัว
	// แทรกโลโก้ที่มุมบนซ้าย
	logoOpt := gofpdf.ImageOptions{
		ImageType:             "PNG",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}
	pdf.ImageOptions("images/logo.png", 20, 10, 25, 0, false, logoOpt, 0, "")

	pdf.SetFontSize(18)
	pdf.SetXY(55, 10)
	pdf.Cell(0, 10, "แบบตรวจสอบคุณวุฒิการศึกษาออนไลน์ ระดับบัณฑิตศึกษา")
	pdf.Ln(12)

	pdf.SetFontSize(16)
	pdf.SetXY(110, 30)
	pdf.CellFormat(0, 10, "ข้อมูลผู้สำเร็จการศึกษา (Graduate Information)", "", 1, "", false, 0, "")
	pdf.Ln(3)

	pdf.SetFont("THSarabun", "", 16)
	pdf.SetXY(100, 40)
	pdf.Cell(0, 8, fmt.Sprintf("ชื่อ - สกุล: %s", studentName))
	pdf.SetXY(100, 50)
	pdf.Cell(0, 8, fmt.Sprintf("Name - Surname: %s", studentName))
	pdf.SetXY(100, 60)
	pdf.Cell(0, 8, fmt.Sprintf("รหัสประจำตัวนักศึกษา: %s", studentCode))
	pdf.SetXY(100, 70)
	pdf.Cell(0, 8, fmt.Sprintf("Student Code: %s", studentCode))
	pdf.SetXY(100, 80)
	pdf.Cell(0, 8, fmt.Sprintf("ปริญญาที่สำเร็จการศึกษา: %s", degree))
	pdf.SetXY(100, 90)
	pdf.Cell(0, 8, fmt.Sprintf("Degree Awarded: %s", degree))
	pdf.Ln(10)

    // 4. Register image จาก memory
    opt := gofpdf.ImageOptions{
        ImageType:             "PNG",
        ReadDpi:               false,
        AllowNegativePosition: false,
    }

	qrFile := fmt.Sprintf("tmp_qr_%s.png", studentID)


	pdf.RegisterImageOptionsReader(qrFile, opt, &buf)
	pdf.ImageOptions(qrFile, 10, 50, 80, 80, false, opt, 0, "")
	pdf.SetXY(30, 130)
	pdf.Cell(0, 8, "QR-code เฉพาะบุคคล")
	// คำอธิบายท้ายกระดาษ
	pdf.SetXY(10, 140)
	pdf.MultiCell(0, 6, `ผู้สำเร็จการศึกษา 
โปรดนำส่งแบบตรวจสอบคุณวุฒิการศึกษาออนไลน์ระดับบัณฑิตศึกษามหาวิทยาลัยรามคำแหงฉบับนี้พร้อมสำเนาหนังสือสำคัญ
แสดงคุณวุฒิ(ใบรับรองสภามหาวิทยาลัยฯ,ใบปริญญาบัตรหรือใบรับรองผลการศึกษา(Transcript) ยังหน่วยงานภาครัฐ
หรือภาคเอกชนที่ท่านประสงค์ ปรับคุณวุฒิการศึกษาสมัครงานศึกษาต่ออื่นๆเพื่อให้หน่วยงานของท่านสามารถตรวจสอบ 
และขอหนังสือรับรองคุณวุฒิการศึกษาโดยตรงกับ มหาวิทยาลัยเป็นไปด้วยความถูกต้องรวดเร็ว

หน่วยงานภาครัฐหรือภาคเอกชน 
สามารถตรวจสอบและขอหนังสือรับรองคุณวุฒิการศึกษาออนไลน์ระดับบัณฑิตศึกษา(ปริญญาโท-ปริญญาเอก) 
ของมหาวิทยาลัยรามคำแหง โดยดำเนินการ ดังนี้
1. Scan QR-code ในแบบตรวจสอบคุณวุฒิการศึกษาออนไลน์ของผู้สำเร็จการศึกษารายบุคคลข้างต้น 
2. กรอกข้อมูลที่เกี่ยวข้องในเว็บไซต์ ..................(E-Mail,ชื่อบริษัท/หน่วยงาน,
ชื่อผู้รับผิดชอบในการตรวจสอบคุณวุฒิการศึกษา)
3. ระบบจะแสดงผลการตรวจสอบและรับรองคุณวุฒิการศึกษาซึ่งหน่วยงานสามารถสั่งพิมพ์หนังสือรับรอง
คุณวุฒิการศึกษา ได้ผ่านทางเว็บไซต์ www.e-regis.ru.ac.th 

หมายเหตุ : 
1.ระบบนี้จัดทำขึ้นเพื่อให้หน่วยงานภายนอกสามารถตรวจสอบคุณวุฒิการศึกษาของผู้สำเร็จการศึกษา
จากมหาวิทยาลัยรามคำแหงระดับปริญญาโทและปริญญาเอก 
2. QR-code มีอายุการใช้งานไม่เกิน 120 วันนับจากวันที่ออกหนังสือ
3. หากต้องการตรวจสอบข้อมูลนอกเหนือจากที่ปรากฏหรือมีปัญหาข้อสงสัยโปรดติดต่อหน่วยตรวจสอบการสำเร็จ
การศึกษาฝ่ายบริการการศึกษาบัณฑิตวิทยาลัยมหาวิทยาลัยรามคำแหง โทร.0-2310-8000 ต่อ 3708 
หรือ 0-2310-8561 หรือ E-Mail: rugrad_verify@ru.ac.th`, "", "L", false)

    // 6. ส่ง PDF กลับไป
    c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename=student_`+studentID+`.pdf`)
    _ = pdf.Output(c.Writer)
}
