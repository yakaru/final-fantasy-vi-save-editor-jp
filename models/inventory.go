package models

type Row struct {
	ItemID string
	Count  int
}

type Inventory struct {
	Rows []Row
}

func newInventory() *Inventory {
	inv := &Inventory{
		Rows: make([]Row, 255),
	}
	for i := 0; i < 255; i++ {
		inv.Rows[i] = Row{ItemID: "FF", Count: 0}
	}
	return inv
}

var inventory *Inventory

func GetInventory() *Inventory {
	if inventory == nil {
		inventory = newInventory()
	}
	return inventory
}