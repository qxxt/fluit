package fluit

import (
	"fmt"
	"testing"
)

var testOptions = []Option{
	{"--hello", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu"},
	{"-a, --abcd", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"-b, --bcde=STRING", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"-c", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"--if-the-arguments-is-longer-than-specified-max-size", "It will used up the entire line and the desc will use the line after it while still keeping the margin. "},
}

var testOptionsWants = []string{
	`  --hello          Lorem ipsum dolor sit amet, consectetur
                   adipiscing elit. Sed in mattis leo.
                   Integer eu tortor ut libero aliquet
                   dignissim. Etiam nisi metus, consectetur
                   eu luctus vel, malesuada id arcu
`,
	`  -a, --abcd       Lorem ipsum dolor sit amet, consectetur
                   adipiscing elit. Sed in mattis leo. 
`,
	`  -b, --bcde=STRING
                   Lorem ipsum dolor sit amet, consectetur
                   adipiscing elit. Sed in mattis leo. 
`,
	`  -c               Lorem ipsum dolor sit amet, consectetur
                   adipiscing elit. Sed in mattis leo. 
`,
	`  --if-the-arguments-is-longer-than-specified-max-size
                   It will used up the entire line and the
                   desc will use the line after it while
                   still keeping the margin. 
`,
}

func TestSprintUsage(t *testing.T) {
	for i := range testOptions {
		if SprintUsage(15, testOptions[i].Argument, testOptions[i].Description) != testOptionsWants[i] {
			fmt.Println(SprintUsage(15, testOptions[i].Argument, testOptions[i].Description))
			t.Error("Formatted usage not equal to predefined output")
		}
	}
}

const (
	lorem         string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci.`
	loremHashLike string = `LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsectetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsollicitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivamusfringillalacinianisinechendreritDuisipsumtortorcongueutesteuvolutpatpharetraorci`
)

func TestSprintwrap(t *testing.T) {
	var got, want string
	UserBreakpoint = 60

	got = SprintWrap(0, lorem)
	want = `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Sed in mattis leo. Integer eu tortor ut libero aliquet
dignissim. Etiam nisi metus, consectetur eu luctus vel,
malesuada id arcu. Integer egestas velit a velit
sollicitudin venenatis. In volutpat nunc posuere ex
lobortis maximus. Vivamus fringilla lacinia nisi nec
hendrerit. Duis ipsum tortor, congue ut est eu, volutpat
pharetra orci.`

	if got != want {
		t.Error("Invalid output on 0 margin")
	}

	got = SprintWrap(0, loremHashLike)
	want = `LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisle
oIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsec
tetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsoll
icitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivam
usfringillalacinianisinechendreritDuisipsumtortorcongueutest
euvolutpatpharetraorci`
	if got != want {
		t.Error("Invalid output on 0 margin and hashlike input")
	}

	UserBreakpoint = 30

	got = SprintWrap(4, lorem)
	want = `    Lorem ipsum dolor sit
    amet, consectetur
    adipiscing elit. Sed in
    mattis leo. Integer eu
    tortor ut libero aliquet
    dignissim. Etiam nisi
    metus, consectetur eu
    luctus vel, malesuada id
    arcu. Integer egestas
    velit a velit
    sollicitudin venenatis.
    In volutpat nunc posuere
    ex lobortis maximus.
    Vivamus fringilla lacinia
    nisi nec hendrerit. Duis
    ipsum tortor, congue ut
    est eu, volutpat pharetra
    orci.`

	if got != want {
		t.Error("Invalid output")
	}

	got = SprintWrap(4, loremHashLike)
	want = `    Loremipsumdolorsitametcons
    ecteturadipiscingelitSedin
    mattisleoIntegereutortorut
    liberoaliquetdignissimEtia
    mnisimetusconsectetureuluc
    tusvelmalesuadaidarcuInteg
    eregestasvelitavelitsollic
    itudinvenenatisInvolutpatn
    uncposuereexlobortismaximu
    sVivamusfringillalaciniani
    sinechendreritDuisipsumtor
    torcongueutesteuvolutpatph
    aretraorci`
	if got != want {
		fmt.Println(got)
		t.Error("invalid output on hashlike input")
	}
}
