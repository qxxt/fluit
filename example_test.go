package fluit_test

import (
	"fmt"

	"github.com/qxxt/fluit"
)

func Example() {
	fluit.UserBreakpoint = 80

	// if w, _, err := term.GetSize(1); err != nil {
	//	fluit.UserBreakpoint = w
	// }

	fluit.PrintlnWrap(4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim. Donec malesuada accumsan porttitor. In neque libero, sollicitudin vitae interdum ut, semper lacinia mauris. Integer auctor, nisl at commodo feugiat, justo urna varius arcu, a lacinia mi mauris sit amet purus.")

	fmt.Println("Usages:")
	u := fluit.Usages{}
	u.AddOption("lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddOption("ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddOption("sit", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.AddOption("amet", "Donec malesuada accumsan porttitor.")
	u.AddOption("consectetur", "Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.PrintUsages()
	u.UsageGroup = nil

	fmt.Print("\nAditional flags:\n")
	u.ArgumentRowLength = 20
	u.AddOption("--lorem string", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddOption("--ipsum int", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddOption("--sit=sits", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.AddOption("--amet, -a", "Donec malesuada accumsan porttitor.")
	u.AddOption("--consectetur", "Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.PrintUsages()
}
