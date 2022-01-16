package fluit_test

import (
	"fmt"

	"github.com/qxxt/fluit"
)

func Example() {
	fluit.SetBreakpoint(70)
	fmt.Print(fluit.SprintfWrap(4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim. Donec malesuada accumsan porttitor. In neque libero, sollicitudin vitae interdum ut, semper lacinia mauris. Integer auctor, nisl at commodo feugiat, justo urna varius arcu, a lacinia mi mauris sit amet purus."))

	fmt.Print("\nUsages:\n")
	u := fluit.Usages{}
	u.AddUsg("lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddUsg("ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddUsg("sit", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.AddUsg("amet", "Donec malesuada accumsan porttitor.")
	u.AddUsg("consectetur", "Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.PrintUsg()

	fmt.Print("\nAditional flags:\n")
	u.SetArgLen(20)
	u.AddUsg("--lorem string", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddUsg("--ipsum int", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddUsg("--sit=sits", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.AddUsg("--amet, -a", "Donec malesuada accumsan porttitor.")
	u.AddUsg("--consectetur", "Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.PrintUsg()
}
