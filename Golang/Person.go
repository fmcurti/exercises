package Person

type Person struct {
	age int
}

func findOldestPerson(persons []Person) int {
	// Returns index of the oldest person in the list -  O(n) complexity
	var oldestIndex int = 0
	var oldestAge int = persons[0].age
	for i := 1; i < len(persons); i++ {
		if oldestAge < persons[i].age {
			oldestAge = persons[i].age
			oldestIndex = i
		}
	}
	return oldestIndex
}

func oneIsDobuleTheOthers(persons []Person) bool {
	//Returns true if there is a person who is exactly twice as old as any other person in the list - O(2n) complexity

	if len(persons) == 0 {
		return false
	}

	oldestIndex := findOldestPerson(persons) // O(n)
	var i int = 0

	for i < len(persons) { // O(n)
		if i != oldestIndex {
			if persons[oldestIndex].age != persons[i].age*2 {
				break
			}
		}
		i++
	}

	return i == len(persons)
}
