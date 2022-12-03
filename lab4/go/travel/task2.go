package travel

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func InitGUI_Task2() {
	window := ui.NewWindow("Калькулятор вартості замовлення туру", 750, 470, false)

	labelResult := ui.NewLabel("")

	button := ui.NewButton("Розрахувати")

	entryDays := ui.NewSlider(0, 31)

	country := ui.NewCombobox()
	country.Append("Болгарія")
	country.Append("Німеччина")
	country.Append("Польща")
	country.SetSelected(0)

	season := ui.NewCombobox()
	season.Append("Зима")
	season.Append("Літо")
	season.SetSelected(1)
	checkbox := ui.NewCheckbox("Індивідуальний гід")
	checkbox1 := ui.NewCheckbox("Номер люкс")

	box1 := ui.NewVerticalBox()
	box2 := ui.NewVerticalBox()

	group1 := ui.NewGroup("Замовлення туру")
	box1.Append(group1, false)

	entryForm1 := ui.NewForm()
	entryForm1.SetPadded(true)
	group1.SetChild(entryForm1)

	entryForm1.Append("К-сть днів", entryDays, false)
	entryForm1.Append("Країна", country, false)

	group2 := ui.NewGroup("Сезон")
	box2.Append(group2, false)

	entryForm2 := ui.NewForm()
	entryForm2.SetPadded(true)
	group2.SetChild(entryForm2)

	entryForm2.Append("", season, false)
	entryForm2.Append("", checkbox, false)
	entryForm2.Append("", checkbox1, false)
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
		var k float64 = float64(country.Selected())

		if season.Selected() == 0 {
			check = 0.1
		} else {
			check = 0
		}
		k = k + check
		switch k {
		case 0:
			sum = float64(100 * (entryDays.Value()))
		case 0.1:
			sum = float64(150 * (entryDays.Value()))
		case 1:
			sum = float64(160 * (entryDays.Value()))
		case 1.1:
			sum = float64(200 * (entryDays.Value()))
		case 2:
			sum = float64(120 * (entryDays.Value()))
		case 2.1:
			sum = float64(180 * (entryDays.Value()))
		default:
			sum = 0
		}
		if checkbox.Checked() == true {
			sum = sum + float64(entryDays.Value()*50)
		}
		if checkbox1.Checked() == true {
			sum = sum + sum/5
		}

		labelResult.SetText(fmt.Sprintf("%.2f", sum))
	})
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()

}
