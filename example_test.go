package fluit_test

import (
	"fmt"
	"strings"

	"github.com/qxxt/fluit"
)

func ExampleSetBreakpoint() {
	fluit.SetBreakpoint(20)
	fmt.Print(fluit.Sprintwrap(0, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis."))
	fmt.Print(fluit.Sprintwrap(5, strings.Repeat("-+", 5)))

	fluit.SetBreakpoint(60)
	fmt.Print(fluit.Sprintwrap(0, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis."))
	fmt.Print(fluit.Sprintwrap(5, strings.Repeat("-+", 20)))

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

func ExampleSprintwrap() {
	// In essence fluit is a string formatter. It don't prints.
	str := fluit.Sprintwrap(4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis.")
	fmt.Print(str)
	fmt.Print(fluit.Sprintwrap(10, strings.Repeat("-+", 20)))

	//If the margin is set to 0, it will not be margined but still wrapped according to breakpoint
	fmt.Print(fluit.Sprintwrap(0, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus euismod pharetra sodales. Aenean ac massa dictum, gravida nisl non, fermentum turpis."))

	// Output:
	//     Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	//     Vivamus euismod pharetra sodales. Aenean ac massa
	//     dictum, gravida nisl non, fermentum turpis.
	//           -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	// Vivamus euismod pharetra sodales. Aenean ac massa dictum,
	// gravida nisl non, fermentum turpis.
}

func ExampleSprintUsg() {
	fmt.Print(fluit.SprintUsg(10, "--help", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante."))
	fmt.Print(fluit.SprintUsg(10, "--build", "Nunc gravida metus non turpis sagittis, at lacinia magna tristique. Suspendisse vel ullamcorper eros. Sed eget lacinia urna."))
	fmt.Print(fluit.SprintUsg(10, "--install", "Cras et tellus dignissim, tempor ligula nec, lobortis dui. In eu enim a urna cursus porta."))

	// Output:
	//   --help      Lorem ipsum dolor sit amet, consectetur
	//               adipiscing elit. Integer quis aliquam ante.
	//   --build     Nunc gravida metus non turpis sagittis, at
	//               lacinia magna tristique. Suspendisse vel
	//               ullamcorper eros. Sed eget lacinia urna.
	//   --install   Cras et tellus dignissim, tempor ligula nec,
	//               lobortis dui. In eu enim a urna cursus porta.
}

func ExampleUsgs() {
	var u fluit.Usgs
	u.SetArgLen(10)
	u.AddUsg("--help", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante.")
	u.AddUsg("--build", "Nunc gravida metus non turpis sagittis, at lacinia magna tristique. Suspendisse vel ullamcorper eros. Sed eget lacinia urna.")
	u.AddUsg("--install", "Cras et tellus dignissim, tempor ligula nec, lobortis dui. In eu enim a urna cursus porta.")
	u.AddUsg("--some-very-long-flag", "Vivamus facilisis sagittis tristique. Quisque mattis sed libero id congue. Duis vel elit risus.")
	u.AddUsg("--update", "Proin accumsan orci vel ante consequat, ut facilisis ipsum congue.")
	u.PrintUsg()

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

func ExampleUsgs_SetArgLen() {
	u := fluit.Usgs{}
	u.AddUsg("--help", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante.")
	u.AddUsg("--dereference-command-line-symlink-to-dir", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante.")
	u.AddUsg("--another-help", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante.")
	u.PrintUsg()

	u.SetArgLen(14)
	u.AddUsg("--help", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante.")
	u.AddUsg("--dereference-command-line-symlink-to-dir", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante.")
	u.AddUsg("--another-help", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis aliquam ante.")
	u.PrintUsg()

	// Output:
	//   --help                                     Lorem ipsum
	//                                              dolor sit amet,
	//                                              consectetur
	//                                              adipiscing
	//                                              elit. Integer
	//                                              quis aliquam
	//                                              ante.
	//   --dereference-command-line-symlink-to-dir  Lorem ipsum
	//                                              dolor sit amet,
	//                                              consectetur
	//                                              adipiscing
	//                                              elit. Integer
	//                                              quis aliquam
	//                                              ante.
	//   --another-help                             Lorem ipsum
	//                                              dolor sit amet,
	//                                              consectetur
	//                                              adipiscing
	//                                              elit. Integer
	//                                              quis aliquam
	//                                              ante.
	//   --help          Lorem ipsum dolor sit amet, consectetur
	//                   adipiscing elit. Integer quis aliquam
	//                   ante.
	//   --dereference-command-line-symlink-to-dir
	//                   Lorem ipsum dolor sit amet, consectetur
	//                   adipiscing elit. Integer quis aliquam
	//                   ante.
	//   --another-help  Lorem ipsum dolor sit amet, consectetur
	//                   adipiscing elit. Integer quis aliquam
	//                   ante.
}
