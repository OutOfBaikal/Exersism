//
// This is only a SKELETON file for the 'Line Up' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const format = (text, num) => {
  const dig = num % 10;
  const lastTwo = num % 100;

  switch (true) {
    case (dig === 1 && lastTwo !== 11):
      return `${text}, you are the ${num}st customer we serve today. Thank you!`;
      
    case (dig === 2 && lastTwo !== 12):
      return `${text}, you are the ${num}nd customer we serve today. Thank you!`;
      
    case (dig === 3 && lastTwo !== 13):
      return `${text}, you are the ${num}rd customer we serve today. Thank you!`;
      
    default:
      return `${text}, you are the ${num}th customer we serve today. Thank you!`;
  }
};
