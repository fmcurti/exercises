package Person

import (
	"testing"
)

func Test01True(t *testing.T) {
	persons := []Person{{age: 4}, {age: 8}, {age: 9}}
	result := oneIsDoubleAgeofAnother(persons)
	if result != true {
		t.Error("Result should be true")
	}

}

func Test02OnePersonShouldBeFalse(t *testing.T) {
	persons := []Person{{8}}
	result := oneIsDoubleAgeofAnother(persons)
	if result != true {
		t.Error("Result should be true")
	}

}

func Test03EmptyList(t *testing.T) {
	persons := []Person{}
	result := oneIsDoubleAgeofAnother(persons)
	if result != false {
		t.Error("Result should be false")
	}
}

func Test04ShouldBeFalse(t *testing.T) {
	persons := []Person{{1}, {1}, {1}, {3}}
	result := oneIsDoubleAgeofAnother(persons)
	if result != false {
		t.Error("Result should be false")
	}
}

func Test05ShouldBeFalse(t *testing.T) {
	persons := []Person{{1}, {1}, {1}, {3}}
	result := oneIsAtLeastDoubleAgeofAnother(persons)
	if result != false {
		t.Error("Result should be false")
	}
}
