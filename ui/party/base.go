package party

import (
	pri "ffvi_editor/models/consts/pr"
	"ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"fmt"
	"github.com/aarzilli/nucular"
)

type partyUI struct {
	selected [4]int
}

func NewUI() ui.UI {
	return &partyUI{}
}

func (u *partyUI) Draw(w *nucular.Window) {
	p := pr.GetParty()
	w.Row(24).Static(300)
	w.Label("Currently In Testing", "LC")
	w.Row(24).Static(260)
	w.CheckboxText("Enable", &p.Enabled)
	for i := 0; i < 4; i++ {
		u.drawRow(w, i, p, &u.selected[i])
	}
}

func (u *partyUI) drawRow(w *nucular.Window, slot int, p *pr.Party, selected *int) {

	w.Row(24).Static(60, 200, 200)
	w.Label(fmt.Sprintf("Member: %d", slot+1), "LC")
	if i := w.ComboSimple(p.PossibleNames, *selected, 12); i != *selected {
		*selected = i
		p.Members[slot] = p.Possible[i]
	}
	if !pri.IsMainCharacter(p.PossibleNames[*selected]) {
		w.Label("Will not work in most cases", "LC")
	}
	//if p.Members[slot].CharacterID != 0 {
	//	w.CheckboxText("Enable Equipment: ", &p.Members[slot].EnableEquipment)
	//}
}

func (u *partyUI) Refresh() {
	p := pr.GetParty()
	for i, m := range p.Members {
		u.selected[i] = p.GetPossibleIndex(m)
	}
}

func (u *partyUI) Name() string {
	return "Party"
}

func (u *partyUI) Behavior() ui.Behavior {
	return ui.PrShow
}
