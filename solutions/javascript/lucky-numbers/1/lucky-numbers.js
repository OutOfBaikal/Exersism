// @ts-check

/**
 * Calculates the sum of the two input arrays.
 *
 * @param {number[]} array1
 * @param {number[]} array2
 * @returns {number} sum of the two arrays
 */
export function twoSum(array1, array2) {
  const num1 = BigInt(array1.join(''));
  const num2 = BigInt(array2.join(''));

  return Number(num1 + num2);
}

/**
 * Checks whether a number is a palindrome.
 *
 * @param {number} value
 * @returns {boolean} whether the number is a palindrome or not
 */
export function luckyNumber(value) {
  let reversed = 0n;
  let temp = BigInt(value);
  let original = BigInt(value);

  while (temp > 0n) {
    reversed = reversed * 10n + temp % 10n;
    temp /= 10n;
  }
  return original === reversed;
}

/**
 * Determines the error message that should be shown to the user
 * for the given input value.
 *
 * @param {string|null|undefined} input
 * @returns {string} error message
 */
export function errorMessage(input) {
   if (!input || input.trim() === "") {
    return "Required field";
  }

  const data = Number(input);

  if (isNaN(data) || data === 0) {
    return "Must be a number besides 0";
  }
  return "";
}
