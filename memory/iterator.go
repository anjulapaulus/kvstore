package memory

type iterator struct {
	records []record
	index   int
}

func (i *iterator) Seek(key interface{}) {
	for idx, r := range i.records {
		if r.key == key {
			i.index = idx
		}
	}
}

func (i *iterator) Next() {
	i.index--
}

func (i *iterator) HasNext() bool {
	return i.index > 0
}

func (i *iterator) Previous() {
	i.index++
}

func (i *iterator) HasPrevious() bool {
	return i.index < len(i.records)
}

func (i *iterator) Key() interface{} {
	return i.records[i.index].key
}

func (i *iterator) Value() interface{} {
	return i.records[i.index].value
}

func (i *iterator) Close() {
	i.records = nil
}
