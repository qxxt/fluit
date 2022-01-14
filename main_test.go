package fluit

import (
	"reflect"
	"testing"
)

var flags = []struct {
	arg string
	msg string
}{
	{"-a", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"-b", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"-c", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. "},
	{"--if-the-arguments-is-longer-than-specified-max-size", "It will used up the entire line and the desc will use the line after it while still keeping the margin. "},
}

func TestSprintUsg(t *testing.T) {
	SetBreakpoint(70)
	var (
		got, want   string
		gots, wants Usages
	)
	got = SprintUsg("--hello", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu", 8)
	want = `  --hello   Lorem ipsum dolor sit amet, consectetur adipiscing elit.
            Sed in mattis leo. Integer eu tortor ut libero aliquet
            dignissim. Etiam nisi metus, consectetur eu luctus vel,
            malesuada id arcu
`
	if got != want {
		t.Error("Formatted string not equal to predefined output")
	}

	gots.SetArgLen(15)
	wants.maxArgLen = 15
	for i := 0; i < len(flags); i++ {
		f := flags[i]
		gots.AddUsg(f.arg, f.msg)
		wants.usageItem =
			append(wants.usageItem, usage{f.arg, f.msg})
	}
	if !reflect.DeepEqual(gots, wants) {
		t.Error("Usages's object not valid")
	}
	gots.PrintUsg()
}

var lorem string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed in mattis leo. Integer eu tortor ut libero aliquet dignissim. Etiam nisi metus, consectetur eu luctus vel, malesuada id arcu. Integer egestas velit a velit sollicitudin venenatis. In volutpat nunc posuere ex lobortis maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis ipsum tortor, congue ut est eu, volutpat pharetra orci.`
var loremHashLike string = `LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisleoIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsectetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsollicitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivamusfringillalacinianisinechendreritDuisipsumtortorcongueutesteuvolutpatpharetraorci`

func TestSprintwrap(t *testing.T) {
	var got, want string
	SetBreakpoint(60)
	got = Sprintwrap(0, lorem)
	want = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed
in mattis leo. Integer eu tortor ut libero aliquet
dignissim. Etiam nisi metus, consectetur eu luctus vel,
malesuada id arcu. Integer egestas velit a velit
sollicitudin venenatis. In volutpat nunc posuere ex lobortis
maximus. Vivamus fringilla lacinia nisi nec hendrerit. Duis
ipsum tortor, congue ut est eu, volutpat pharetra orci.
`
	if got != want {
		t.Error("Invalid output on 0 margin")
	}
	got = Sprintwrap(0, loremHashLike)
	want = `LoremipsumdolorsitametconsecteturadipiscingelitSedinmattisle
oIntegereutortorutliberoaliquetdignissimEtiamnisimetusconsec
tetureuluctusvelmalesuadaidarcuIntegeregestasvelitavelitsoll
icitudinvenenatisInvolutpatnuncposuereexlobortismaximusVivam
usfringillalacinianisinechendreritDuisipsumtortorcongueutest
euvolutpatpharetraorci
`
	if got != want {
		t.Error("Invalid output on 0 margin and hashlike input")
	}
	SetBreakpoint(30)
	got = Sprintwrap(4, lorem)
	want = `    Lorem ipsum dolor sit
    amet, consectetur
    adipiscing elit. Sed in
    mattis leo. Integer eu
    tortor ut libero aliquet
    dignissim. Etiam nisi
    metus, consectetur eu
    luctus vel, malesuada id
    arcu. Integer egestas
    velit a velit sollicitudin
    venenatis. In volutpat
    nunc posuere ex lobortis
    maximus. Vivamus fringilla
    lacinia nisi nec
    hendrerit. Duis ipsum
    tortor, congue ut est eu,
    volutpat pharetra orci.
`
	if got != want {
		t.Error("Invalid output")
	}
	got = Sprintwrap(4, loremHashLike)
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
    aretraorci
`
	if got != want {
		t.Error("invalid output on hashlike input")
	}
}
