/// <reference path="./global.d.ts" />
//
// @ts-check

/**
 * Determine the price of the pizza given the pizza and optional extras
 *
 * @param {Pizza} pizza name of the pizza to be made
 * @param {Extra[]} extras list of extras
 *
 * @returns {number} the price of the pizza
 */
export function pizzaPrice(pizza, ...extras) {
  let sum = 0;
  switch (pizza) {
    case 'Margherita':
        sum += 7;
        break;
    case 'Caprese':
        sum += 9;
        break;
    case 'Formaggio':
        sum += 10;
        break;
  }

  for (let extra of extras) {
    switch (extra) {
        case "ExtraSauce":
          sum += 1;
          break;
        case "ExtraToppings":
          sum += 2;
          break;
      default:
        console.log("У меня лапки");
    }
  }

  return sum;
}

/**
 * Calculate the price of the total order, given individual orders
 *
 * (HINT: For this exercise, you can take a look at the supplied "global.d.ts" file
 * for a more info about the type definitions used)
 *
 * @param {PizzaOrder[]} pizzaOrders a list of pizza orders
 * @returns {number} the price of the total order
 */
export function orderPrice(pizzaOrders) {
  let sum = 0;
  for (let order of pizzaOrders) {
    sum += pizzaPrice(order.pizza, ...order.extras);
  }

  return sum;
}
