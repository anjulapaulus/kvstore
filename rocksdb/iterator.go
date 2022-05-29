package rocksdb

import "github.com/linxGnu/grocksdb"

type iterator struct {
	itr *grocksdb.Iterator
}

func (i *iterator) SeekToFirst() {
	i.itr.SeekToFirst()
}

func (i *iterator) ValidForPrefix(prefix []byte) bool {
	return i.itr.ValidForPrefix(prefix)
}

func (i *iterator) SeekPrefix(prefix []byte) bool {
	for i.Valid() {
		if i.ValidForPrefix(prefix) {
			return true
		}
		i.Next()
	}
	return false
}

func (i *iterator) SeekToLast() {
	i.itr.SeekToLast()
}

func (i *iterator) Seek(key []byte) {
	i.itr.Seek(key)
}

func (i *iterator) Next() {
	i.itr.Next()
}

func (i *iterator) Prev() {
	i.itr.Prev()
}

func (i *iterator) Close() {
	i.itr.Close()
}

func (i *iterator) Key() []byte {
	return i.itr.Key().Data()
}

func (i *iterator) Value() []byte {
	return i.itr.Value().Data()
}

func (i *iterator) Valid() bool {
	return i.itr.Valid()
}

func (i *iterator) Error() error {
	return i.itr.Err()
}
