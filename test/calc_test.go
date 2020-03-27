package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}
	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}
}
func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(1, 2) != 2 {
			t.Fatal("fail")
		}
	})
	t.Run("neg", func(t *testing.T) {
		if Mul(2, 2) != 4 {
			t.Fatal("fauldd")
		}
	})
}

//涉及多个子测试时的推荐写法
func TestMul2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Excepted int
	}{
		{"pos", 1, 2, 2},
		{"sec", 2, 2, 4},
		{"thr", 3, 2, 6},
	}

	for _, res := range cases {
		t.Run(res.Name, func(t *testing.T) {
			if ans := Mul(res.A, res.B); ans != res.Excepted {
				t.Fatalf("%d * %d expected %d, but %d got",
					res.A, res.B, res.Excepted, ans)
			}
		})
	}
}
