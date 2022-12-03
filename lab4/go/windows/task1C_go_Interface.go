package windows

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
)

func InitGUI_c_go() {
	window := ui.NewWindow("Калькулятор склопакета", 750, 470, false)

	labelResult := ui.NewLabel("")

	button := ui.NewButton("Розрахувати")

	entryWidth := ui.NewEntry()
	entryHeight := ui.NewEntry()

	material := ui.NewCombobox()
	material.Append("Дерево")
	material.Append("Метал")
	material.Append("Металопластик")
	material.SetSelected(0)

	sklopacket := ui.NewCombobox()
	sklopacket.Append("Однокамерний")
	sklopacket.Append("Двокамерний")
	sklopacket.SetSelected(0)
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
		material_parse := material.Selected()
		sklopacket_parse := sklopacket.Selected()

		widthConv, _ := strconv.ParseFloat(entryWidth.Text(), 64)
		heightConv, _ := strconv.ParseFloat(entryHeight.Text(), 64)

		result := calculateC(material_parse, sklopacket_parse, checkbox.Checked(), heightConv, widthConv)
		labelResult.SetText(fmt.Sprintf("%f", result))
	})
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()

}
