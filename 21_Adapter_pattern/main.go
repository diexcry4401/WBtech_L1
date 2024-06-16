package main

/*
Реализовать паттерн «адаптер» на любом примере.
*/
import "fmt"

// VideoPort - интерфейс для зарядного порта
type VideoPort interface {
	adapterForHDMI() string // Метод адаптера для подключения к HDMI
}

// HDMI - объект для HDMI порта
type HDMI struct{}

// connectHDMIPort - имитация подключения к порту HDMI
func (*HDMI) connectHDMIPort() string {
	return "connected to HDMI port"
}

// newAdapter - конструктор адаптера для HDMI к Type-C
func newAdapter(mu *HDMI) VideoPort {
	return &TypeC{video: mu} // Возвращает адаптер Type-C, использующий HDMI
}

// TypeC - объект для Type-C порта
type TypeC struct {
	video *HDMI // Поле для хранения объекта HDMI
}

// adapterForHDMI - адаптер Type-C для HDMI
func (tc *TypeC) adapterForHDMI() string {
	return tc.video.connectHDMIPort() // Использует HDMI через адаптер Type-C
}

func main() {
	typeC := newAdapter(&HDMI{})        // Создание адаптера для HDMI к Type-C
	fmt.Println(typeC.adapterForHDMI()) // Вывод результата подключения через адаптер
}
