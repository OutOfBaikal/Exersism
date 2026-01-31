package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }
    return time * len(layers)
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    noodles, sauce := 0, 0
    for i, _ := range layers {
        if layers[i] == "noodles" {
            noodles += 1
        } else if layers[i] == "sauce" {
            sauce += 1
        }
    }
    return noodles * 50, float64(sauce) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend_recipe []string, your_recipe []string) {
    your_recipe[len(your_recipe) - 1] = friend_recipe[len(friend_recipe) - 1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount []float64, portions int) []float64 {
    new_amount := append([]float64{}, amount...)
    for i, _ := range new_amount {
        new_amount[i] *= float64(portions) / 2
    }
    return new_amount
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
