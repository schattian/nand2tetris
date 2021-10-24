package symbol

type Table struct {
	table          map[string]*Symbol
	kindPopulation map[Kind]int
}

func (t *Table) Get(name string) *Symbol {
	return t.table[name]
}

func (t *Table) Add(name string, kind Kind, typ Type) {
	lastIndex, ok := t.kindPopulation[kind]
	if !ok {
		lastIndex = -1
	}
	s := &Symbol{Kind: kind, Type: typ, Index: lastIndex + 1}
	t.table[name] = s
	t.kindPopulation[kind] = s.Index
}
