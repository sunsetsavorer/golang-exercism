package lasagna_master

const (
	DEFAULT_AVG_TIME  = 2
	NOODLES_PER_LAYER = 50
	SAUCE_PER_LAYER   = 0.2
)

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, avgTime int) int {

	if avgTime <= 0 {
		avgTime = DEFAULT_AVG_TIME
	}

	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {

	noodles := 0
	sauce := 0.0

	for _, name := range layers {

		switch name {
		case "noodles":
			noodles += NOODLES_PER_LAYER
		case "sauce":
			sauce += SAUCE_PER_LAYER
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendIngridients, myIngridients []string) {

	myIngridients[len(myIngridients)-1] = friendIngridients[len(friendIngridients)-1]
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, amount int) []float64 {

	newQuantities := make([]float64, len(quantities))

	scale := float64(amount) / 2

	for i, v := range quantities {
		newQuantities[i] = v * scale
	}

	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
