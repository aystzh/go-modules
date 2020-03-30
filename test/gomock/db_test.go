package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() //断言 DB.Get方法是否被调用到

	m := NewMockDB(ctrl)
	//参数(Eq, Any, Not, Nil)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(-1, errors.New("not exist")) //打桩
	m.EXPECT().Get(gomock.Any()).Do(func(key string) {
		t.Log(key)
	})
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))

	m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		if key == "Sam" {
			return 630, nil
		}
		return 0, errors.New("not exist")
	})

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}

}

func TestGetFromDB2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Sam")).Return(-1, nil).Times(3)
	GetFromDB(m, "cde")
	GetFromDB(m, "ABC")
	GetFromDB(m, "cdd")
}

func TestGetFromDB3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(-1, nil)

	//指定顺序
	gomock.InOrder(o1, o2)
	GetFromDB(m, "Tom")
	GetFromDB(m, "Sam")
}
