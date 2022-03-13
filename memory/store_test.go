package memory

import "testing"

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
