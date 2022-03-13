package memory

import (
	"testing"
)

func TestNewMemoryStore(t *testing.T) {
	st := NewMemoryStore("test")
	if st == nil {
		t.Error("TestNewMemoryStore: store nil")
	}
}

func TestName(t *testing.T) {
	st := NewMemoryStore("test")
	if st.Name() != "test" {
		t.Error("TestName: name is incorrect")
	}
}

type testData struct {
	key  int
	data string
}

func TestSet(t *testing.T) {
	st := NewMemoryStore("test")
	err := st.Set(1, testData{
		key:  2000,
		data: "test blaa",
	})
	if err != nil {
		t.Error("TestSet: failed to set")
	}
}

func TestGet(t *testing.T) {
	st := NewMemoryStore("test")
	err := st.Set(1, testData{
		key:  2000,
		data: "test blaa",
	})
	if err != nil {
		t.Error("TestSet: failed to set")
	}
	v, _ := st.Get(1)
	if v == nil {
		t.Error("TestSet: received nil value")
	}
}

func TestDelete(t *testing.T) {
	st := NewMemoryStore("test")
	err := st.Set(1, testData{
		key:  2000,
		data: "test blaa",
	})
	if err != nil {
		t.Error("TestSet: failed to set")
	}
	err = st.Delete(1)
	if err != nil {
		t.Error("TestSet: received nil value")
	}
	v, _ := st.Get(1)
	if v != nil {
		t.Error("TestSet: received value")
	}
}

func TestCreatedAt(t *testing.T) {
	st := NewMemoryStore("test")
	err := st.Set(1, testData{
		key:  2000,
		data: "test blaa",
	})
	if err != nil {
		t.Error("TestCreatedAt: failed to set")
	}

	tm := st.CreatedAt(1)

	if tm.IsZero() {
		t.Error("TestCreatedAt: time is zero")
	}
}

func TestCreatedAtZeroError(t *testing.T) {
	st := NewMemoryStore("test")
	tm := st.CreatedAt(1)

	if !tm.IsZero() {
		t.Error("TestCreatedAtZeroError: time is zero")
	}
}

func TestNewIterator(t *testing.T) {
	st := NewMemoryStore("test")
	err := st.Set(1, testData{
		key:  2000,
		data: "test blaa",
	})
	if err != nil {
		t.Error("TestNewIterator: failed to set")
	}

	i, err := st.NewIterator()
	if err != nil || i == nil {
		t.Error("TestNewIterator: failed create iterator")
	}
}

func TestNewStoreTTl(t *testing.T) {
	st := NewMemoryStore("test", WithTTL(1000000000))
	if st.expireAfter == 1 {
		t.Error("TestNewStoreTTl: failed to set expiry", st.expireAfter.Nanoseconds())
	}
	st.cleanup()
}
