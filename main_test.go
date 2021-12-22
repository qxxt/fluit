package fluit

import (
	"fmt"
	"testing"

	"golang.org/x/term"
)

var sprintfm_tests = []struct {
	margin int
	text   string
}{
	{4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci."},
	{10, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci. "},
	{4, "Lorem ipsum dolor sit amet,"},
	{10, "Lorem ipsum dolor sit amet,"},
	{4, "LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsectetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsollicitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivamusfringillalacinianisinechendreritDuisipsumtortorcongueutesteuvolutpatpharetraorci"},
	{10, "LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsectetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsollicitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivamusfringillalacinianisinechendreritDuisipsumtortorcongueutesteuvolutpatpharetraorci"},
}

var sprintfm_usg_tests = []struct {
	arg string
	msg string
}{
	{"--lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci."},
	{"--lorem-ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci. "},
	{"--lorem", "Lorem ipsum dolor sit amet,"},
	{"--lorem-ipsum", "Lorem ipsum dolor sit amet,"},
	{"--lorem", "LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsectetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsollicitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivamusfringillalacinianisinechendreritDuisipsumtortorcongueutesteuvolutpatpharetraorci"},
	{"--lorem-ipsum", "LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsectetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsollicitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivamusfringillalacinianisinechendreritDuisipsumtortorcongueutesteuvolutpatpharetraorci"},
	{"--lorem-ipsum-dolor-sit-amet", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci."},
	{"--lorem-ipsum-dolor-sit-amet", "LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsectetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsollicitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivamusfringillalacinianisinechendreritDuisipsumtortorcongueutesteuvolutpatpharetraorci"},
}

func TestTerminal(t *testing.T) {
	is_terminal := term.IsTerminal(1)
	terminal_width, _, err_getsize := term.GetSize(1)
	_, err_getstate := term.GetState(1)
	if is_terminal == false {
		t.Errorf("error term.IsTerminal(): not running on terminal, or is not recognized as one.")
	}
	if err_getstate != nil {
		t.Errorf("error term.GetState(): %q.", err_getstate)
	}
	if err_getsize != nil {
		t.Errorf("error term.GetSize(): %q", err_getsize)
		fmt.Printf("Terminal Width: unknown\r\n")
		defer t_fluits(0, t)
		t.FailNow()
	}
	fmt.Printf("Terminal Width: %d\r\n", terminal_width)
	t_fluits(terminal_width, t)
}

func t_fluits(terminal_width int, t *testing.T) {
	for index, _ := range sprintfm_tests {
		formatted, err := fmt_margin_wrap(sprintfm_tests[index].text, sprintfm_tests[index].margin, terminal_width)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(formatted)
	}
	const max_args_length int = 20
	for index, _ := range sprintfm_usg_tests {
		formatted, err := sprintf_usg_verbose(sprintfm_usg_tests[index].arg, sprintfm_usg_tests[index].msg, max_args_length, terminal_width)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(formatted)
	}
}
