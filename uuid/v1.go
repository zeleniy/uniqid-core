package uuid

import (
	"encoding/binary"
	"fmt"
	"sync/atomic"
	"time"
)

// UUID представляет собой массив из 16 байт
type UUID [16]byte

// NewUUIDv1 создает новый UUID версии 1
func NewUUIDv1() UUID {

	var uuid UUID

	// Получаем текущее время в формате 100-наносекундных интервалов с 1582 года
	timestamp := time.Now().UnixNano()/100 + 0x01B21DD213814000

	// Заполняем первые 8 байт меткой времени
	binary.BigEndian.PutUint32(uuid[0:4], uint32(timestamp>>32)) // старшие 32 бита
	binary.BigEndian.PutUint16(uuid[4:6], uint16(timestamp>>16)) // средние 16 бит
	binary.BigEndian.PutUint16(uuid[6:8], uint16(timestamp))     // младшие 16 бит

	// Устанавливаем версию (биты 12-15) как 1 (UUID v1)
	uuid[6] = (uuid[6] & 0x0F) | 0x10

	// Генерируем уникальный идентификатор узла (обычно MAC-адрес)
	mac := GetMac()
	copy(uuid[10:], mac)

	// Устанавливаем вариант (биты 6-7 байта 8) как RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	// Генерируем случайный счетчик для предотвращения коллизий
	counter := atomic.AddUint32(&counterValue, 1)
	binary.BigEndian.PutUint16(uuid[8:10], uint16(counter))

	return uuid
}

// counterValue - глобальный счетчик для предотвращения коллизий
var counterValue uint32

// String возвращает строковое представление UUID
func (uuid UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}
