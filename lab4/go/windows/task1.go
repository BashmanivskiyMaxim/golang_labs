package windows

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
)

func InitGUI() {
	window := ui.NewWindow("Калькулятор склопакета", 750, 470, false)

	labelResult := ui.NewLabel("")

	button := ui.NewButton("Розрахувати")

	entryWidth := ui.NewEntry()
	entryHeight := ui.NewEntry()

	material := ui.NewCombobox()
	material.Append("Дерево")
	material.Append("Метал")
	material.Append("Металопластик")
	material.SetSelected(1)

	sklopacket := ui.NewCombobox()
	sklopacket.Append("Однокамерний")
	sklopacket.Append("Двокамерний")
	sklopacket.SetSelected(1)
	checkbox := ui.NewCheckbox("Підвіконня")

	box1 := ui.NewVerticalBox()
	box2 := ui.NewVerticalBox()

	group1 := ui.NewGroup("Розмір вікна")
	box1.Append(group1, false)

	entryForm1 := ui.NewForm()
	entryForm1.SetPadded(true)
	group1.SetChild(entryForm1)

	entryForm1.Append("Ширина, см", entryWidth, false)
	entryForm1.Append("Висота, см", entryHeight, false)
	entryForm1.Append("Матеріал", material, false)

	group2 := ui.NewGroup("Склопакет")
	box2.Append(group2, false)

	entryForm2 := ui.NewForm()
	entryForm2.SetPadded(true)
	group2.SetChild(entryForm2)

	entryForm2.Append("", sklopacket, false)
	entryForm2.Append("", checkbox, false)
	entryForm2.Append("", labelResult, false)

	box3 := ui.NewHorizontalBox()
	box3.Append(button, true)

	box5 := ui.NewHorizontalBox()
	box5.Append(box1, false)
	box5.Append(box2, true)

	box4 := ui.NewVerticalBox()
	box4.Append(box5, false)
	box4.Append(box3, false)

	window.SetMargined(true)
	window.SetChild(box4)

	button.OnClicked(func(*ui.Button) {
		var sum float64 = 0
		var check float64 = 0
		var k float64 = float64(material.Selected())

		widthConv, _ := strconv.ParseFloat(entryWidth.Text(), 64)
		heightConv, _ := strconv.ParseFloat(entryHeight.Text(), 64)

		if sklopacket.Selected() == 1 {
			check = 0.1
		} else {
			check = 0
		}
		k = k + check
		switch k {
		case 0:
			sum = 2.5 * (widthConv * heightConv)
		case 0.1:
			sum = 3 * (widthConv * heightConv)
		case 1:
			sum = 0.5 * (widthConv * heightConv)
		case 1.1:
			sum = 1.0 * (widthConv * heightConv)
		case 2:
			sum = 1.5 * (widthConv * heightConv)
		case 2.1:
			sum = 2.0 * (widthConv * heightConv)
		default:
			sum = 0
		}
		if checkbox.Checked() == true {
			sum = sum + 350
		}
		labelResult.SetText(fmt.Sprintf("%.2f", sum))
	})
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()

}

func main() {
	err := ui.Main(InitGUI)
	if err != nil {
		panic(err)
	}
}
