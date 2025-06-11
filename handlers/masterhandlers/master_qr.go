package masterhandlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
)

func (h *studentHandlers) GenerateStudentPDF(c *gin.Context) {
	studentID := c.Param("id")

	// (ตัวอย่างข้อมูล)
	studentName := "สมชาย ใจดี"
	major := "วิศวกรรมคอมพิวเตอร์"
	year := "2567"
	verifyURL := fmt.Sprintf("http://ruconnext-dev.ru.ac.th:9100/master/student/successcheck/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW4iOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKaFkyTmxjM05mZEc5clpXNWZhMlY1SWpvaU5UTXhNamMwTURJd09EbzZZV05qWlhOek9qb3laV0V6WkdSbU9DMWpNMkV5TFRReFl6Z3RZamhqTkMwek5HSTVZalJtTmpNNE5EUWlMQ0psZUhCcGNtVnpYM1J2YTJWdUlqb3hOelE1TnpBeU5ERTRMQ0pwYzNOMVpYSWlPaUpTZFMxamIyNXVaWGgwSWl3aWNtVm1jbVZ6YUY5MGIydGxibDlyWlhraU9pSTFNekV5TnpRd01qQTRPanB5WldaeVpYTm9Pam8yTmpJeVlXRmxZeTFrT1dJd0xUUTVNR1l0WVdVeE1TMDBObVpsTjJaaE16QXdZVEFpTENKeWIyeGxJam9pVFdGemRHVnlJaXdpYzNSa1gyTnZaR1VpT2lJMU16RXlOelF3TWpBNElpd2ljM1ZpYW1WamRDSTZJbEoxTFVOdmJtNWxlSFE2T2pVek1USTNOREF5TURnaWZRLkNyUVJqelZKM2tEdkNrRGhYcmZkb0U1RllwN29JUm56bDNJV2xnem9TdHMiLCJhY2Nlc3NfdG9rZW5fa2V5IjoiNTMxMjc0MDIwODo6Y2VydGlmaWNhdGU6OmM0YzJjMjdlLTEwMTctNGFhMC1iM2Q4LTFkNjliMTYwMjk5ZSIsImNlcnRpZmljYXRlIjoiZWdyYWR1YXRlIiwiZXhwaXJlX2RhdGUiOiIyMDI1LTA2LTExVDEyOjM5OjMxLjQxNTM5Njc3NyswNzowMCIsImV4cGlyZXNfdG9rZW4iOjE3NDk2MjAzNzEsImlzc3VlciI6IlJ1LWNvbm5leHRlZ3JhZHVhdGUiLCJzdGFydF9kYXRlIjoiMjAyNS0wNi0xMVQxMTozOTozMS40MTUzOTY3NzcrMDc6MDAiLCJzdGRfY29kZSI6IjUzMTI3NDAyMDgiLCJzdWJqZWN0IjoiQ2VydGlmaWNhdGU6OjUzMTI3NDAyMDgifQ.PA2uzyStdCGVHu0A4rHsRrejm4ndfy-7Vd_xNQREaOU")

	// สร้าง QR code ลงไฟล์
	qrFile := fmt.Sprintf("tmp_qr_%s.png", studentID)
	err := qrcode.WriteFile(verifyURL, qrcode.Medium, 400, qrFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "QR code error")
		return
	}
	defer os.Remove(qrFile) // ลบไฟล์หลังใช้

	// เริ่มสร้าง PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("THSarabun", "", "fonts/THSarabunNew.ttf")
	pdf.SetFont("THSarabun", "", 16)
	pdf.AddPage()

	pdf.Cell(40, 10, "ข้อมูลนักศึกษา")
	pdf.Ln(12)
	pdf.Cell(40, 10, "รหัสนักศึกษา: "+studentID)
	pdf.Ln(10)
	pdf.Cell(40, 10, "ชื่อ: "+studentName)
	pdf.Ln(10)
	pdf.Cell(40, 10, "สาขา: "+major)
	pdf.Ln(10)
	pdf.Cell(40, 10, "ปีการศึกษาที่จบ: "+year)
	pdf.Ln(15)
	pdf.Cell(40, 10, "ตรวจสอบวุฒิได้ที่: "+verifyURL)
	pdf.Ln(10)

	// แสดง QR Code จากไฟล์
	pdf.Image(qrFile, 10, 100, 100, 0, false, "", 0, "")

	// ส่ง PDF
	err = pdf.Output(c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, "PDF generate error" + err.Error())
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename=student_`+studentID+`.pdf`)
}
