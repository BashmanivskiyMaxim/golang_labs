package travel

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func InitGUI_c_go2() {
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
	season.SetSelected(0)
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
		labelResult.SetText(fmt.Sprintf("%.2f", calculateC(int64(entryDays.Value()),
			country.Selected(), season.Selected(), checkbox.Checked(), checkbox1.Checked())))
	})
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()

}
