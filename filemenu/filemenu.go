package filemenu

import (
	"bufio"
	"fmt"
	"todolist/color"
	"todolist/storage"
)

func FileMenu(reader *bufio.Reader) {
	for {
		fmt.Println(color.Blue("\n=== File Menu ==="))
		fmt.Println(color.Blue("1.") + "Toggle autosave")
		fmt.Println(color.Blue("2.") + "Save as...")
		fmt.Println(color.Blue("3.") + "Export to text file")
		fmt.Println(color.Blue("4.") + "Back to menu")
		fmt.Print(color.Blue("\nChoose an action: "))

		input := storage.ReadInput(reader)

		switch input {
		case "1":
			storage.SetAutoSave()
		case "2":
			storage.SaveAs(reader)
		case "3":
			storage.ExportToText(reader)
		case "4":
			return
		default:
			fmt.Println(color.Red("Invalid choice. Please try again."))
		}
	}
}
