package ram_test

import (
	"aliya-ram/app/ram"
	"testing"
)

func TestRAM_SearchMemory(t *testing.T) {
	ram, err := ram.NewRAM()
	if err != nil {
		t.Fatal(err)
	}
	ram.AddMemory("1", "Aliya:哼\n要是我一开始就知道原来是这小家伙...")
	ram.AddMemory("2", "Aliya:啊\n你们那个时代的人不认识这个很正常\n是人造的培育物种\n不过本身也不是地球上的生物")
	res, err := ram.SearchMemory("2", "外星生物")
	if err != nil {
		t.Fatal(err)
	}
	for _, memory := range res {
		t.Log(memory)
	}
}

func TestRAM_ListMemory(t *testing.T) {
	ram, err := ram.NewRAM()
	if err != nil {
		t.Fatal(err)
	}
	res, err := ram.ListMemories("2")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRAM_DeleteAllMemories(t *testing.T) {
	ram, err := ram.NewRAM()
	if err != nil {
		t.Fatal(err)
	}
	err = ram.DeleteAllMemories("2")
	if err != nil {
		t.Fatal(err)
	}
}
