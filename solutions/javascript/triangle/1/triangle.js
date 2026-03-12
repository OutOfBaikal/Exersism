//
// This is only a SKELETON file for the 'Triangle' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class Triangle {
  constructor(...sides) {
    let [a, b, c] = sides;
    this.a = a;
    this.b = b;
    this.c = c;
    this.isValid = (a > 0 && b > 0 && c > 0) && 
                   (a + b >= c) && (a + c >= b) && (b + c >= a) &&
                   (a + b + c > 0); 
    this.isRealTriangle = this.isValid && (a + b > c && a + c > b && b + c > a);
  }

  get isEquilateral() {
    if (!this.isRealTriangle) return false;
    return this.a === this.b && this.b === this.c;
  }

  get isIsosceles() {
    if (!this.isRealTriangle) return false;
    return this.a === this.b || this.a === this.c || this.b === this.c;
  }

  get isScalene() {
    if (!this.isRealTriangle) return false;
    return this.a !== this.b && this.a !== this.c && this.b !== this.c;
  }
}
