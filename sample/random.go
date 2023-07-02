package sample

import (
	"math/rand"

	"github.com/YuichiNAGAO/grpc_sample_app/pb"
	"github.com/google/uuid"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "ARM")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet("Core i7-8700K", "Xeon E-2286M", "Core i5-9400F")
	}
	return randomStringFromSet("Ryzen 7 2700X", "Ryzen 5 2600", "Ryzen 3 2200G")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet("GeForce RTX 2080 Ti", "GeForce GTX 1660 Ti", "GeForce GTX 1060")
	}
	return randomStringFromSet("Radeon RX 5700 XT", "Radeon RX 5600 XT", "Radeon RX 5500 XT")
}

func randomScreenResolution() *pb.Screen_Resolution {
	return &pb.Screen_Resolution{
		Width:  uint32(randomInt(1920, 3840)),
		Height: uint32(randomInt(1080, 2160)),
	}
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("MacBook Air", "MacBook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	case "Lenovo":
		return randomStringFromSet("ThinkPad X1", "ThinkPad P1", "ThinkPad T490", "ThinkPad E595")
	}
	return ""
}

func randomStringFromSet(s ...string) string {
	return s[rand.Intn(len(s))]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}
