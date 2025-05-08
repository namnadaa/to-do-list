package filemenu

import (
	"bufio"
	"fmt"
	"todolist/color"
	"todolist/storage"
)

// FileMenu displays a menu for file-related operations such as toggling autosave,
// saving tasks to a custom file, and exporting tasks to a text file.
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
			err := storage.SaveAs(reader)
			if err != nil {
				msg := fmt.Sprintf("[ERROR] Failed to save file: %v", err)
				fmt.Println(color.Magenta(msg))
			}
		case "3":
			err := storage.ExportToText(reader)
			if err != nil {
				msg := fmt.Sprintf("[ERROR] Failed to export file: %v", err)
				fmt.Println(color.Magenta(msg))
			}
		case "4":
			return
		default:
			fmt.Println(color.Red("Invalid choice. Please try again."))
		}
	}
}
