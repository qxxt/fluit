package fluit

import (
	"fmt"
	"testing"

	"golang.org/x/term"
)

var sprintfm_lists = []struct {
	margin int
	text   string
}{
	{4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci."},
	{10, "Donec vitae massa diam. Vivamus posuere blandit turpis id varius. Sed scelerisque turpis ut nisi suscipit, semper ultricies sapien interdum. Nulla id nulla varius, imperdiet urna at, ullamcorper tellus. Fusce mollis ornare nisi. Donec ultricies porta elit, quis consectetur tellus consequat vitae. Nullam vestibulum nec nibh in tincidunt. Aliquam sed felis et mi eleifend ultrices. Fusce at nisi laoreet erat consequat sodales sed id turpis. "},
	{4, "Lorem ipsum"},
	{10, "Donec vitae"},
	{4, "Lorem-ipsum-dolor-sit-amet-consectetur-adipiscing-elit-Sed-in-mattis-leo-Integer-eu-tortor-ut-libero-aliquet-dignissim-Etiam-nisi-metus-consectetur-eu-luctus-vel-malesuada-id-arcu-Integer-egestas-velit-a-velit-sollicitudin-venenatis-In-volutpat-nunc-posuere-ex-lobortis-maximus-Vivamus-fringilla-lacinia-nisi-nec-hendrerit-Duis-ipsum-tortor-congue-ut-est-eu-volutpat-pharetra-orci"},
	{10, "Donec-vitae-massa-diam-Vivamus-posuere-blandit-turpis-id-varius-Sed-scelerisque-turpis-ut-nisi-suscipit-semper-ultricies-sapien-interdum-Nulla-id-nulla-varius-imperdiet-urna-at-ullamcorper-tellus-Fusce-mollis-ornare-nisi-Donec-ultricies-porta-elit-quis-consectetur-tellus-consequat-vitae-Nullam-vestibulum-nec-nibh-in-tincidunt-Aliquam-sed-felis-et-mi-eleifend-ultrices-Fusce-at-nisi-laoreet-erat-consequat-sodales-sed-id-turpis"},
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
	for index, _ := range sprintfm_lists {
		formatted, err := fmt_margin_wrap(sprintfm_lists[index].text, sprintfm_lists[index].margin, terminal_width)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(formatted)
	}
}
