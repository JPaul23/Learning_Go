package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/JPaul23/go-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func createPrintJob(invoiceId int) {
	client := resty.New()
	var p models.PrintJob
	// Call PrinterService via RESTful interface
	_, err := client.R().
		SetBody(models.PrintJob{Format: "A4", InvoiceId: invoiceId}).
		SetResult(&p).
		Post("http://localhost:5000/print-jobs")

	if err != nil {
		log.Println("InvoiceGenerator: unable to connect PrinterService")
		return
	}
	log.Printf("InvoiceGenerator: created print job #%v via PrinterService", p.JobId)
}
func main() {
	router := gin.Default()
	router.POST("/invoices", func(c *gin.Context) {
		var iv models.Invoice
		if err := c.ShouldBindJSON(&iv); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input!"})
			return
		}
		log.Println("InvoiceGenerator: creating new invoice...")
		rand.Seed(time.Now().UnixNano())
		iv.InvoiceId = rand.Intn(1000)
		log.Printf("InvoiceGenerator: created invoice #%v", iv.InvoiceId)

		createPrintJob(iv.InvoiceId) // Ask PrinterService to create a print job
		c.JSON(200, iv)
	})
	router.Run(":6000")
}
