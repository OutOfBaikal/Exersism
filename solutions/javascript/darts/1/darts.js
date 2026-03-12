//
// This is only a SKELETON file for the 'Darts' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const score = (x, y) => {
  let score = Math.sqrt(x * x + y * y);
  if (score > 10.0) {
    return 0;
  } else if (score <= 10 && score > 5) {
    return 1;
  } else if (score <= 5 && score > 1) {
    return 5;
  }
  return 10;
};
