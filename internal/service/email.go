package service

import (
	"fmt"
	"log"
)

type EmailService struct {
	// Tambahkan konfigurasi yang diperlukan, misalnya SMTP server, API key, dll.
}

// SendEmail sends an email to the specified address with the given subject and body.
func (s *EmailService) SendEmail(to, subject, body string) error {
	// Implementasi pengiriman email; ini adalah placeholder sederhana.
	fmt.Printf("Sending email to %s with subject %s\n", to, subject)
	// Placeholder untuk implementasi pengiriman email (misalnya menggunakan SMTP atau API pihak ketiga)
	log.Println("Email sent successfully.")
	return nil
}
