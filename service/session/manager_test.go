package session


import "testing"

func TestManager(t *testing.T){
	body := &Body{
		Secret: "root",
	}
	if NewStore().Set(body) != NewStore().Set(body){
		t.Fatal("result is wrong!")
	}

	item1 := NewStore().Set(body); 
	item2,err := NewStore().Get(item1.SessionID)
	if err != nil || item1 != item2 {
		t.Fatal("result is wrong!")
	}

	if _, err := NewStore().Get("xxx"); err == nil{
		t.Fatal("result is wrong!")
	}
}

func BenchmarkManagerSet(b *testing.B){
	for i := 0; i < b.N; i++ {
		NewStore().Set(&Body{
			Secret: string(i),
		})
    }
}

func BenchmarkManagerGet(b *testing.B){
	for i := 0; i < b.N; i++ {
		NewStore().Get(string(i))
    }
}