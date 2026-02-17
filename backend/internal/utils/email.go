package utils

import (
	"blood-donor-system/internal/config"
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	host := config.SMTP.Host
	port := config.SMTP.Port
	user := config.SMTP.User
	pass := config.SMTP.Pass
	from := config.SMTP.From

	if host == "" || port == "" || user == "" || pass == "" {
		return fmt.Errorf("SMTP configuration is incomplete")
	}

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		body

	auth := smtp.PlainAuth("", user, pass, host)
	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func SendVerificationEmail(to, token string) error {
	subject := "Verify your email - BloodDonation"
	// Replace with actual frontend URL in prod
	verifyURL := fmt.Sprintf("http://localhost:3000/verify-email?token=%s", token)
	body := fmt.Sprintf(`
		<h1>Email Verification</h1>
		<p>Thank you for registering. Please click the link below to verify your email:</p>
		<a href="%s">Verify Email</a>
		<p>If you did not register, please ignore this email.</p>
	`, verifyURL)

	return SendEmail(to, subject, body)
}

func SendPasswordResetEmail(to, token string) error {
	subject := "Reset your password - BloodDonation"
	resetURL := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", token)
	body := fmt.Sprintf(`
		<h1>Password Reset Request</h1>
		<p>You requested a password reset. Click the link below to set a new password:</p>
		<a href="%s">Reset Password</a>
		<p>This link will expire in 1 hour.</p>
		<p>If you did not request this, please ignore this email.</p>
	`, resetURL)

	return SendEmail(to, subject, body)
}

func SendDonationNotification(to, name, date, location string) error {
	subject := "New Donation Added - BloodDonation"
	body := fmt.Sprintf(`
		<h1>Donation Record Added</h1>
		<p>Hi %s,</p>
		<p>A new donation record has been added to your profile:</p>
		<ul>
			<li><strong>Date:</strong> %s</li>
			<li><strong>Location:</strong> %s</li>
		</ul>
		<p>Thank you for your life-saving contribution!</p>
	`, name, date, location)

	return SendEmail(to, subject, body)
}
