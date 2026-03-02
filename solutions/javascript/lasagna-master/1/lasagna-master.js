/// <reference path="./global.d.ts" />
// @ts-check

/*
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

export function cookingStatus(time) {
  if (time === 0) {
    return "Lasagna is done.";
  }
  if (time > 0) {
    return "Not done, please wait.";
  }
  return "You forgot to set the timer.";
}

export function preparationTime (layers, time) {
  if (!time) {
    time = 2;
  }
  return time * layers.length;
}

export function quantities(layers) {
  let noodles = 0;
  let sauce = 0;
  for (let i in layers) {
    if (layers[i] == "noodles") {
      noodles += 1;
    } else if (layers[i] == "sauce") {
      sauce += 1;
    }
  }

  return {
    noodles: noodles * 50, 
    sauce: sauce * 0.2
  };
}

export function addSecretIngredient(friend_recipe, your_recipe) {
  const secret = friend_recipe[friend_recipe.length - 1];
  
  your_recipe.push(secret);
}

export function scaleRecipe(recipe, portions) {
  const scaled = {};
  const factor = portions / 2;

  for (const ingredient in recipe) {
    scaled[ingredient] = recipe[ingredient] * factor;
  }

  return scaled;
}