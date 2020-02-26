package slice

import "testing"

func TestFillValues(t *testing.T) {
	xs := []X{
		X{
			Name: "A",
			Age:  1,
		},
		X{
			Name: "B",
			Age:  2,
		},
	}
	//vs := []string{"C", "D"}
	vs2 := []int{100, 200}
	err := StructFill(xs, vs2, "Age")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(xs)
}

type X struct {
	Name string
	Age  int
}

func TestStructPick(t *testing.T) {
	xs := []*X{
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "B",
			Age:  2,
		},
	}
	var dst []int
	err := StructPick(&dst, xs, "Age")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dst)

}
