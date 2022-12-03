package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"lab4/travel"
	"lab4/windows"
	"os"
)

func main() {
	var count int
	fmt.Println("Обчислення вартості склопакета Go - 1;\nОбчислення вартості склопакета Cgo - 2;" +
		"\nОбчислення вартості подорожі Go - 3;\nОбчислення вартості подорожі Cgo - 4")
	var err error
	fmt.Fscan(os.Stdin, &count)

	switch count {
	case 1:
		err = ui.Main(windows.InitGUI)
	case 2:
		err = ui.Main(windows.InitGUI_c_go)
	case 3:
		err = ui.Main(travel.InitGUI_Task2)
	case 4:
		err = ui.Main(travel.InitGUI_c_go2)

	}
	err = ui.Main(travel.InitGUI_Task2)

	if err != nil {
		panic(err)
	}
}
