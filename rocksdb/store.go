package rocksdb

import (
	"fmt"
	"os"
	"sync"

	"github.com/linxGnu/grocksdb"
)

type rocksdbStore struct {
	name string
	rdb  *grocksdb.DB
	wo   *grocksdb.WriteOptions
	ro   *grocksdb.ReadOptions
	mu   sync.Mutex
}

type rocksOptions struct {
	ttl       int
	directory string
	fSync     bool
}

type rocksOps func(*rocksOptions)

func WithTTL(ttl int) rocksOps {
	return func(opts *rocksOptions) {
		opts.ttl = ttl
	}
}

func WithDirectory(directory string) rocksOps {
	return func(opts *rocksOptions) {
		opts.directory = directory
	}
}
func WithFSync(fSyncEnabled bool) rocksOps {
	return func(opts *rocksOptions) {
		opts.fSync = fSyncEnabled
	}
}
func applyDefault() *rocksOptions {
	ops := new(rocksOptions)
	ops.ttl = 0
	ops.directory = "storage"
	ops.fSync = false
	return ops
}

func NewRocksDB(name string, options ...rocksOps) (*rocksdbStore, error) {
	opts := applyDefault()
	for _, opt := range options {
		opt(opts)
	}
	rOpts := grocksdb.NewDefaultOptions()
	rOpts.SetCreateIfMissing(true)
	rOpts.SetUseFsync(opts.fSync)

	if err := createDirectory(opts.directory); err != nil {
		return nil, fmt.Errorf("cannot create directory error:%v", err)
	}

	db := &rocksdbStore{}
	db.name = name

	if opts.ttl == 0 {
		rdb, err := grocksdb.OpenDb(rOpts, opts.directory)
		if err != nil {
			return nil, err
		}
		db.rdb = rdb
	} else {
		rdb, err := grocksdb.OpenDbWithTTL(rOpts, opts.directory, opts.ttl)
		if err != nil {
			return nil, err
		}
		db.rdb = rdb
	}

	db.wo = grocksdb.NewDefaultWriteOptions()
	db.ro = grocksdb.NewDefaultReadOptions()

	return db, nil
}

func createDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (r *rocksdbStore) Name() string {
	return r.name
}

func (r *rocksdbStore) Type() string {
	return "rocksDB"
}

func (r *rocksdbStore) Set(key, value []byte) error {
	return r.rdb.Put(r.wo, key, value)
}

func (r *rocksdbStore) Get(key []byte) ([]byte, error) {
	return r.rdb.GetBytes(r.ro, key)
}

func (r *rocksdbStore) Delete(key []byte) error {
	return r.rdb.Delete(r.wo, key)
}

func (r *rocksdbStore) Close() {
	r.ro.Destroy()
	r.wo.Destroy()
	r.rdb.Close()
}

func (r *rocksdbStore) NewIterator() *iterator {
	return &iterator{
		itr: r.rdb.NewIterator(r.ro),
	}
}

func (r *rocksdbStore) RangeIterator(from, to []byte) *iterator {
	rOpts := r.ro
	rOpts.SetIterateUpperBound(from)

	return &iterator{
		itr: r.rdb.NewIterator(rOpts),
	}
}
