package uuid

import (
	"testing"
)

// TestNewV1 tests the NewV1 function.
func TestNewV1(t *testing.T) {

	uuid := NewUUIDv1()
	length := len(uuid)

	if length != 16 {
		t.Errorf("UUID v1 should be 16 bytes long, %d given", length)
	}
	// // Проверяем, что результат начинается с "UUIDv1-"
	// if len(result) < 7 || result[:7] != "UUIDv1-" {
	//     t.Errorf("Expected result to start with 'UUIDv1-', got '%s'", result)
	// }

	// // Проверяем, что временная метка соответствует текущему времени
	// expectedTime := time.Now().Format("20060102150405")
	// if !timeWithTolerance(result[7:], expectedTime, 1*time.Second) {
	//     t.Errorf("Expected timestamp to match current time, got '%s'", result[7:])
	// }
}
