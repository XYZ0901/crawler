package parser

import (
	"io/ioutil"
	"log"
	"testing"
)

func  TestGetProvinceCitys(t *testing.T) {
	contents,err := ioutil.ReadFile("citylist_test_data.txt")
	if err!=nil {
		log.Fatal(err)
	}
	type args struct {
		all []byte
	}
	tests := []struct {
		name       string
		args       args
	}{
		// TODO: Add test cases.
		{
			"test",
			args{all:contents},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetProvinceCitys(tt.args.all); len(gotResult.Items)!=352{
				t.Errorf("GetProvinceCitys() = %v, want %v", len(gotResult.Items), 352)
			}
		})
	}
}

func  TestParseProfile(t *testing.T) {
	ParseProfile([]byte(""))
}