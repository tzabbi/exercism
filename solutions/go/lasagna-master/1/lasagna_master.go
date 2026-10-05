package lasagnamaster

// PreparationTime calculates the total preparation time
func PreparationTime(layers []string, preparationTimeLayer int) int {
	if preparationTimeLayer == 0 {
		preparationTimeLayer = 2
	}
	return len(layers) * preparationTimeLayer
}

// Quantities calculates the amount of noodles in gram and sauce in liters
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sause float64
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sause += 0.2
		} else {
			continue
		}
	}
	return noodles, sause
}

// AddSecretIngredient replaces "?" with last item of friendsIngredient
func AddSecretIngredient(friendsIngredient, ownIngredient []string) {
	for i, v := range ownIngredient {
		if v == "?" {
			ownIngredient[i] = friendsIngredient[len(friendsIngredient)-1]
		}
	}
}

func ScaleRecipe(amountsNeeded []float64, portions int) []float64 {
	newAmountsNeeded := make([]float64, len(amountsNeeded))
	for i, v := range amountsNeeded {
		newPortion := v * float64(portions/2)
		newAmountsNeeded[i] = newPortion
	}
	return newAmountsNeeded
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
