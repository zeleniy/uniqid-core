package uuid

import (
	"crypto/rand"
	"net"
)

func GetMac() []byte {

	interfaces, err := net.Interfaces()

	if err != nil {
		// Если MAC-адрес недоступен, используем случайные байты
		randomBytes := make([]byte, 6)
		rand.Read(randomBytes)
		return randomBytes
	}

	for _, iface := range interfaces {
		if len(iface.HardwareAddr) == 6 {
			return iface.HardwareAddr
		}
	}

	// Если MAC-адрес не найден, используем случайные байты
	randomBytes := make([]byte, 6)
	rand.Read(randomBytes)

	return randomBytes
}
