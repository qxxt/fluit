package fluit_test

import (
	"fmt"

	"github.com/qxxt/go-fluit"
)

func ExampleSprintwrap() {
	fluit.Sprintwrap(4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis. Duis ullamcorper sagittis luctus. Duis sem arcu, porta at turpis et, consectetur fringilla turpis. Sed diam tellus, tincidunt at est et, tempus tristique purus.")
	//If the margin is set to 0, it will not be margined but still wrapped according to breakpoint
	fmt.Println("---")
	fluit.Sprintwrap(0, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis. Duis ullamcorper sagittis luctus. Duis sem arcu, porta at turpis et, consectetur fringilla turpis. Sed diam tellus, tincidunt at est et, tempus tristique purus.")
}

func ExampleSetBreakpoint() {
	fluit.SetBreakpoint(70)
	fluit.Sprintwrap(4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis. Duis ullamcorper sagittis luctus. Duis sem arcu, porta at turpis et, consectetur fringilla turpis. Sed diam tellus, tincidunt at est et, tempus tristique purus.")
}
