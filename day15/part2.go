package main

/*
--- Part Two ---

Your cookie recipe becomes wildly popular! Someone asks if you can make another recipe that has exactly 500 calories per cookie (so they can use it as a meal replacement). Keep the rest of your award-winning process the same (100 teaspoons, same ingredients, same scoring system).

For example, given the ingredients above, if you had instead selected 40 teaspoons of butterscotch and 60 teaspoons of cinnamon (which still adds to 100), the total calorie count would be 40*8 + 60*3 = 500. The total score would go down, though: only 57600000, the best you can do in such trying circumstances.

Given the ingredients in your kitchen and their properties, what is the total score of the highest-scoring cookie you can make with a calorie total of 500?
*/

func bakeTheBestCalorieCookie(ingredients []Ingredient, maxCalories int) Receipe {        
    receipe := Receipe {}
    receipe.mixture = map[string]int {}
    receipe.ingredients = ingredients
    
    var bestReceipe Receipe
    var highScore int
        
    for {
        score := receipe.getScore()
        
        if score > highScore && receipe.totalCalories == maxCalories && receipe.count() == teaspoonsRequired {
            highScore = score
            bestReceipe = receipe.copyReceipe()  
        }
        if added := receipe.addIngredient(); added == false {
            break
        }
    }
    
    return bestReceipe
}