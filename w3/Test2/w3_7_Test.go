package Test2

func Testinput1ShouldBeDisplay1(t *test.T) {

	v := FizzBuzz(1)

	if "1" != v {
		t.Error("Fizzbuzz of 1 Should be `1` but have", v)
	}
}
