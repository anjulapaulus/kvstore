package memory

import (
	"testing"
	"time"
)

func TestNext(t *testing.T) {
	recs := []Record{
		{
			Key:       1,
			Value:     1,
			createdAt: time.Now(),
		},
		{
			Key:       2,
			Value:     2,
			createdAt: time.Now(),
		},
		{
			Key:       3,
			Value:     3,
			createdAt: time.Now(),
		},
	}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}
	r := i.Next()
	if r.Key != 3 {
		t.Error("TestNext: Not expected result returned")
	}
}

func TestHasNext(t *testing.T) {
	recs := []Record{
		{
			Key:       1,
			Value:     1,
			createdAt: time.Now(),
		},
		{
			Key:       2,
			Value:     2,
			createdAt: time.Now(),
		},
		{
			Key:       3,
			Value:     3,
			createdAt: time.Now(),
		},
	}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}
	b := i.HasNext()
	if !b {
		t.Error("TestHasNext: Not expected result returned", b)
	}
}

func TestHasNextEmpty(t *testing.T) {
	recs := []Record{
		{
			Key:       1,
			Value:     1,
			createdAt: time.Now(),
		},
	}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}
	i.Next()
	b := i.HasNext()
	if b {
		t.Error("TestHasNextEmpty: Not expected result returned", b)
	}
}

func TestPrevious(t *testing.T) {
	recs := []Record{
		{
			Key:       1,
			Value:     1,
			createdAt: time.Now(),
		},
		{
			Key:       2,
			Value:     2,
			createdAt: time.Now(),
		},
	}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}
	i.Next()
	i.Next()
	v := i.Previous()
	if v.Key != 2 {
		t.Error("TestPrevious: Not expected result returned", v.Key)
	}
}

func TestHasPrevious(t *testing.T) {
	recs := []Record{
		{
			Key:       1,
			Value:     1,
			createdAt: time.Now(),
		},
		{
			Key:       2,
			Value:     2,
			createdAt: time.Now(),
		},
	}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}
	i.Next()
	i.Next()
	b := i.HasPrevious()
	if !b {
		t.Error("TestHasPrevious: Not expected result returned", b)
	}
}

func TestHasPreviousEmpty(t *testing.T) {
	recs := []Record{}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}

	b := i.HasPrevious()
	if b {
		t.Error("TestHasPreviousEmpty: Not expected result returned", b)
	}
}

func TestKey(t *testing.T) {
	recs := []Record{
		{
			Key:       1,
			Value:     1,
			createdAt: time.Now(),
		},
		{
			Key:       2,
			Value:     2,
			createdAt: time.Now(),
		},
	}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}
	i.Next()
	k := i.Key()

	if k != 2 {
		t.Error("TestKey: Not expected result returned", k)
	}
}

func TestValue(t *testing.T) {
	recs := []Record{
		{
			Key:       1,
			Value:     1,
			createdAt: time.Now(),
		},
		{
			Key:       2,
			Value:     2,
			createdAt: time.Now(),
		},
	}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}
	i.Next()
	v := i.Value()

	if v.(int) != 2 {
		t.Error("TestValue: Not expected result returned", v)
	}
}

func TestClose(t *testing.T) {
	recs := []Record{
		{
			Key:       1,
			Value:     1,
			createdAt: time.Now(),
		},
		{
			Key:       2,
			Value:     2,
			createdAt: time.Now(),
		},
	}
	i := &Iterator{
		records: recs,
		index:   len(recs),
	}
	i.Close()
	if i.records != nil {
		t.Error("TestClose: Not expected result returned")
	}
}
