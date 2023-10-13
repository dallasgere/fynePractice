package main

import (
	"fynePractice/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")
	w.Resize(fyne.NewSize(300, 400))

	data := []models.Todo{
		models.NewTodo("Some stuff"),
		models.NewTodo("Some more stuff"),
		models.NewTodo("Some other things"),
	}
	todos := binding.NewUntypedList()

	for _, t := range data {
		todos.Append(t)
	}

	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "New Todo Description..."
	addBtn := widget.NewButton("Add", func() {
		todos.Append(models.NewTodo(newtodoDescTxt.Text))
		newtodoDescTxt.Text = ""
	})
	addBtn.Disable()

	newtodoDescTxt.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	w.SetContent(
		container.NewBorder(
			nil, // TOP of the container
			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil, // Left
				// RIGHT ↓
				addBtn,
				// take the rest of the space
				newtodoDescTxt,
			),
			nil, // Right
			nil, // Left

			widget.NewListWithData(

				todos,

				func() fyne.CanvasObject {
					return container.NewBorder(
						nil, nil, nil,

						widget.NewCheck("", func(b bool) {}),

						widget.NewLabel(""),
					)
				},

				func(di binding.DataItem, o fyne.CanvasObject) {
					ctr, _ := o.(*fyne.Container)

					l := ctr.Objects[0].(*widget.Label)
					c := ctr.Objects[1].(*widget.Check)

					todo := models.NewTodoFromDataItem(di)
					l.SetText(todo.Description)
					c.SetChecked(todo.Done)
				}),
		),
	)
	w.ShowAndRun()
}
