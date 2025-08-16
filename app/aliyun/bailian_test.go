package aliyun_test

import (
	"aliya-ram/app/aliyun"
	"log"
	"testing"
)

func TestQuer(t *testing.T) {
	bailian, err := aliyun.NewBalilian()
	if err != nil {
		t.Fatal(err)
	}
	res, err := bailian.Query("菲涅尔")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range res {
		log.Println(c.GetText())
	}
}
