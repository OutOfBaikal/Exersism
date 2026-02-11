// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */
export function timeToMixJuice(name) {
  switch(name) {
    case 'Pure Strawberry Joy':
      return 0.5;
    case 'Energizer':
    case'Green Garden':
      return 1.5;
    case 'Tropical Island':
      return 3;
    case 'All or Nothing':
      return 5;
    default:
      return 2.5;
  }
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */
export function limesToCut(wedgesNeeded, limes) {
  let wedgesCollected = 0;
  let limesCut = 0;

  for (let lime of limes) {
    if (wedgesCollected >= wedgesNeeded) {
      return limesCut;
    }
    // Add wedges based on lime size
    if (lime === 'small') {
      wedgesCollected += 6;
    } else if (lime === 'medium') {
      wedgesCollected += 8;
    } else if (lime === 'large') {
      wedgesCollected += 10;
    }

    limesCut += 1;

    
    if (wedgesCollected >= wedgesNeeded) {
      return limesCut;
    }
  }

  
  return limesCut;
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
export function remainingOrders(timeLeft, orders) {
   while (timeLeft > 0 && orders.length > 0) {
    const juice = orders.shift();
    timeLeft -= timeToMixJuice(juice);
  }
  return orders;
}
