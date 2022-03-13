package memory

type Iterator struct {
	records []Record
	index   int
}

func (i *Iterator) Next() Record {
	i.index--
	return i.records[i.index]
}

func (i *Iterator) HasNext() bool {
	return i.index > 0
}

func (i *Iterator) Previous() Record {
	i.index++
	return i.records[i.index]
}

func (i *Iterator) HasPrevious() bool {
	return i.index < len(i.records)
}

func (i *Iterator) Key() int {
	return i.index
}

func (i *Iterator) Value() Record {
	return i.records[i.index]
}

func (i *Iterator) Close() {
	i.records = nil
}
