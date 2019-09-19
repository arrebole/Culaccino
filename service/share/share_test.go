package share

import "testing"

func TestPassWord(t *testing.T){
	str := "hello world"
	if PassWord(str) != PassWord(str){
		t.Fatal("result is wrong!")
	}
	if  PassWord(str) == str {
		t.Fatal("result is wrong!")
	}
}

func TestRandString(t *testing.T){
	if RandString() == RandString() {
		t.Log(RandString(), RandString())
		t.Fatal("result is wrong!")
	}
}

func TestHashSha256(t *testing.T){
	if HashSha256("hello") != HashSha256("hello") {
		t.Fatal("result is wrong!")
	}
	if HashSha256("hello") == "hello" {
		t.Fatal("result is wrong!")
	}
}

func TestHashMD5(t *testing.T){
	if HashMD5("hello") != HashMD5("hello") {
		t.Fatal("result is wrong!")
	}
	if HashMD5("hello") == "hello" {
		t.Fatal("result is wrong!")
	}
}


func BenchmarkPassWord(b *testing.B){
	for i := 0; i < b.N; i++ {
        PassWord("hello world")
    }
}

func BenchmarkRandString(b *testing.B){
	for i := 0; i < b.N; i++ {
        RandString()
    }
}