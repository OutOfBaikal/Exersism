// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  var sum = 0;
    for (var i = 0; i < birdsPerDay.length; i++) {
        sum += birdsPerDay[i];
    }
    return sum;
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  var sum = 0;
  const n = (week - 1) * 7 + 7;
    for (var i = (week - 1) * 7; i < n; i++) {
        sum += birdsPerDay[i];
    }
    return sum;
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {void} should not return anything
 */
export function fixBirdCountLog(birdsPerDay) {
  for (var i = 0; i < birdsPerDay.length; i += 2) {
        birdsPerDay[i] += 1;
    }
}
