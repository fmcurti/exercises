package Person

import (
	"testing"
)

func Test01True(t *testing.T) {
	persons := []Person{{8}, {8}, {16}, {8}}
	result := oneIsDobuleTheOthers(persons)
	if result != true {
		t.Error("Result should be true")
	}

}

func TestOnePersonShouldBeTrue(t *testing.T) {
	persons := []Person{{8}}
	result := oneIsDobuleTheOthers(persons)
	if result != true {
		t.Error("Result should be true")
	}

}

func TestEmptyList(t *testing.T) {
	persons := []Person{}
	result := oneIsDobuleTheOthers(persons)
	if result != false {
		t.Error("Result should be false")
	}
}

func TestShouldBeFalse(t *testing.T) {
	persons := []Person{{1}, {1}, {2}, {3}}
	result := oneIsDobuleTheOthers(persons)
	if result != false {
		t.Error("Result should be false")
	}
}
