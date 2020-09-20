package Person

type Person struct {
	age int
}

func oneIsDoubleAgeofAnother(persons []Person) bool {
	//Returns true if there is a person who is exactly twice as old as any other person in the list - O(2n) complexity

	if len(persons) <= 1 {
		return false
	}

	for _, person := range persons {
		for _, otherPerson := range persons {
			if person.age == 2*otherPerson.age {
				return true
			}
		}
	}
	return false
}

func oneIsAtLeastDoubleAgeofAnother(persons []Person) bool {
	//Returns true if there is a person who is exactly twice as old as any other person in the list - O(2n) complexity

	if len(persons) <= 1 {
		return false
	}

	for _, person := range persons {
		for _, otherPerson := range persons {
			if person.age >= 2*otherPerson.age {
				return true
			}
		}
	}
	return false
}
