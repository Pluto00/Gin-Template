package gredis

import (
	"Gin-Template/pkg/setting"
	"encoding/json"
	"testing"
)

func init() {
	setting.Setup()
	err := Setup()
	if err != nil {
		panic(err)
	}
}

type testdata struct {
	T1 int    `json:"t1"`
	T2 string `json:"t2"`
	T3 struct {
		T4 string `json:"t4"`
	} `json:"t3"`
}

func TestSet(t *testing.T) {
	data := testdata{
		T1: 1,
		T2: "t2",
		T3: struct {
			T4 string `json:"t4"`
		}{
			T4: "t4",
		},
	}
	err := Set("test", data, 3600)
	if err != nil {
		t.Fatalf("Fail to set 'test': %v", err)
		return
	}
}

func TestGet(t *testing.T) {
	reply, err := Get("test")
	if err != nil {
		t.Fatalf("Fail to get 'test': %v", err)
		return
	}
	data := testdata{}
	err = json.Unmarshal(reply, &data)
	if err != nil {
		t.Fatalf("fail to unmarshal data: %v", err)
		return
	}
	t.Logf("Success to get data: %v", data)
}

func TestExists(t *testing.T) {
	t.Log(Exists("test"))
}

func TestDelete(t *testing.T) {
	_, err := Delete("test")
	if err != nil {
		t.Fatalf("Fail to delete 'test': %v", err)
		return
	}
	t.Log(Exists("test"))
}

func TestLikeDeletes(t *testing.T) {
	_ = Set("1qx2", 1, 3600)
	_ = Set("3qx4", 2, 3600)
	t.Log(Exists("1qx2"))
	t.Log(Exists("3qx4"))
	err := LikeDeletes("qx")
	if err != nil {
		t.Fatalf("Fail to Like Delete 'qx': %v", err)
		return
	}
	t.Log(Exists("1qx2"))
	t.Log(Exists("3qx4"))
}
