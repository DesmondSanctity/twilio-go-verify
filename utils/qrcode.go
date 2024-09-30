package utils

import (
	"fmt"
	qr "github.com/skip2/go-qrcode"
)

// GenerateQRCode creates a QR code image from the given data and returns it as a byte slice
func GenerateQRCode(data string, size int) ([]byte, error) {
	qrCode, err := qr.Encode(data, qr.Medium, size)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}
	return qrCode, nil
}
