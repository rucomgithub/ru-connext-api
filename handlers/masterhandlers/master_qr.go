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
	verifyURL := fmt.Sprintf("https://yourdomain.com/verify/%s", studentID)

	// สร้าง QR code ลงไฟล์
	qrFile := fmt.Sprintf("tmp_qr_%s.png", studentID)
	err := qrcode.WriteFile(verifyURL, qrcode.Medium, 256, qrFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "QR code error")
		return
	}
	defer os.Remove(qrFile) // ลบไฟล์หลังใช้

	// เริ่มสร้าง PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("THSarabun", "", "fonts/Kanit-Thin.ttf")
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
	pdf.Image(qrFile, 10, 100, 50, 0, false, "", 0, "")

	// ส่ง PDF
	err = pdf.Output(c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, "PDF generate error" + err.Error())
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename=student_`+studentID+`.pdf`)
}
