package fluit

import (
	"fmt"
	"testing"
)

type sprintfUsgTest struct {
	arg string
	msg string
}

var sprintfUsg_Cmd = []sprintfUsgTest{
	{"install", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu."},
	{"build", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. "},
	{"uninstall", "Lorem ipsum dolor sit amet,"},
	{"update", "Lorem ipsum dolor sit amet,"},
	{"blablabla", "LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissim"},
}

var sprintfUsg_Flag = []sprintfUsgTest{
	{"-a", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"-b", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"-c", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"--if-the-arguments-is-longer-than-specified-max-size", "It will used up the entire line and the desc will use the line after it while still keeping the margin. "},
}

func TestTerminal(t *testing.T) {
	breakpoint := 70
	fmt.Println(Sprintfm("go-fluit is text formatter for wraping text, adding margin and building cli-usage. It use regex to split text based on specified breakpoint.", 0, breakpoint))
	fmt.Println(Sprintfm("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci.", 4, breakpoint))
	fmt.Println(Sprintfm("LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsectetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsollicitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivamusfringillalacinianisinechendreritDuisipsumtortorcongueutesteuvolutpatpharetraorci", 4, breakpoint))
	fmt.Printf(Sprintfm("Avaliable commands:", 0, breakpoint))
	const maxCmdLength int = 10
	for _, sprintfmUsgTest := range sprintfUsg_Cmd {
		formatted := SprintfUsg(sprintfmUsgTest.arg, sprintfmUsgTest.msg, maxCmdLength, breakpoint)
		fmt.Printf(formatted)
	}
	fmt.Println("")
	fmt.Printf(Sprintfm("Avaliable flag:", 0, breakpoint))
	const maxFlagLength int = 5
	for _, sprintfmUsgTest := range sprintfUsg_Flag {
		formatted := SprintfUsg(sprintfmUsgTest.arg, sprintfmUsgTest.msg, maxFlagLength, breakpoint)
		fmt.Printf(formatted)
	}
}
