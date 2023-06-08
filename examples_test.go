package fluit_test

import (
	"fmt"
	"strings"

	"github.com/qxxt/fluit"
)

func ExampleUserBreakpoint() {
	fluit.UserBreakpoint = 20
	fluit.PrintlnWrap(0, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis.")
	fluit.PrintlnWrap(5, strings.Repeat("-+", 5))

	fluit.UserBreakpoint = 60
	fluit.PrintlnWrap(0, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis.")
	fluit.PrintlnWrap(5, strings.Repeat("-+", 20))

	// Output:
	// Lorem ipsum dolor
	// sit amet,
	// consectetur
	// adipiscing elit.
	// Vivamus euismod
	// pharetra sodales.
	// Aenean ac massa
	// dictum, gravida nisl
	// non, fermentum
	// turpis.
	//      -+-+-+-+-+
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	// Vivamus euismod pharetra sodales. Aenean ac massa dictum,
	// gravida nisl non, fermentum turpis.
	//      -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
}

func ExampleSprintWrap() {
	// In essence fluit is a string formatter.
	str := fluit.Wrap(4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis.")
	fmt.Println(str)

	fmt.Println(fluit.Wrap(10, strings.Repeat("-+", 20)))

	// If the margin is set to 0, it will not be margined but still wrapped according to breakpoint.
	fmt.Println(fluit.Wrap(0, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis."))

	// Output:
	//     Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	//     Vivamus euismod pharetra sodales. Aenean ac massa
	//     dictum, gravida nisl non, fermentum turpis.
	//           -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	// Vivamus euismod pharetra sodales. Aenean ac massa dictum,
	// gravida nisl non, fermentum turpis.
}

func ExampleSprintUsage() {
	fmt.Print(fluit.SprintUsage(10, "--help", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante."))
	fmt.Print(fluit.SprintUsage(10, "--build", "Nunc gravida metus non turpis sagittis, at lacinia magna tristique. Suspendisse vel ullamcorper eros. Sed eget lacinia urna."))
	fmt.Print(fluit.SprintUsage(10, "--install-install", "Cras et tellus dignissim, tempor ligula nec, lobortis dui. In eu enim a urna cursus porta."))

	// Output:
	//   --help      Lorem ipsum dolor sit amet, consectetur
	//               adipiscing elit. Integer quis aliquam ante.
	//   --build     Nunc gravida metus non turpis sagittis, at
	//               lacinia magna tristique. Suspendisse vel
	//               ullamcorper eros. Sed eget lacinia urna.
	//   --install-install
	//               Cras et tellus dignissim, tempor ligula nec,
	//               lobortis dui. In eu enim a urna cursus porta.
}

func ExampleUsages() {
	var u fluit.Usages
	u.ArgumentRowLength = 10
	u.AddOption("--help", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante.")
	u.AddOption("--build", "Nunc gravida metus non turpis sagittis, at lacinia magna tristique. Suspendisse vel ullamcorper eros. Sed eget lacinia urna.")
	u.AddOption("--install", "Cras et tellus dignissim, tempor ligula nec, lobortis dui. In eu enim a urna cursus porta.")
	u.AddOption("--some-very-long-flag", "Vivamus facilisis sagittis tristique. Quisque mattis sed libero id congue. Duis vel elit risus.")
	u.AddOption("--update", "Proin accumsan orci vel ante consequat, ut facilisis ipsum congue.")
	u.PrintUsages()

	// Output:
	//   --help      Lorem ipsum dolor sit amet, consectetur
	//               adipiscing elit. Integer quis aliquam ante.
	//   --build     Nunc gravida metus non turpis sagittis, at
	//               lacinia magna tristique. Suspendisse vel
	//               ullamcorper eros. Sed eget lacinia urna.
	//   --install   Cras et tellus dignissim, tempor ligula nec,
	//               lobortis dui. In eu enim a urna cursus porta.
	//   --some-very-long-flag
	//               Vivamus facilisis sagittis tristique. Quisque
	//               mattis sed libero id congue. Duis vel elit
	//               risus.
	//   --update    Proin accumsan orci vel ante consequat, ut
	//               facilisis ipsum congue.
}
