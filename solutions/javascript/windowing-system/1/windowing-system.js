// @ts-check

/**
 * Implement the classes etc. that are needed to solve the
 * exercise in this file. Do not forget to export the entities
 * you defined so they are available for the tests.
 */

export class Size {
  constructor(w = 80, h = 60) {
    this.width = w;
    this.height = h;
  }
  resize(new_w, new_h) {
    if (new_w < 1) {
      new_w = 1;
    }
    if (new_h < 1) {
      new_h = 1;
    }
    this.width = new_w;
    this.height = new_h;
  }
}

export class Position {
  constructor(x = 0, y = 0) {
    this.x = x;
    this.y = y;
  }
  move(new_x, new_y) {
    if (new_x < 0) {
      new_x = 0;
    }
    if (new_y < 0) {
      new_y = 0;
    }
    this.x = new_x;
    this.y = new_y;
  }
}

export class ProgramWindow {
  constructor() {
    this.screenSize = new Size(800, 600);
    this.size = new Size();
    this.position = new Position();
  }
  resize(newSize) {
    const maxWidth = this.screenSize.width - this.position.x;
    const maxHeight = this.screenSize.height - this.position.y;

    const actualWidth = Math.max(1, Math.min(newSize.width, maxWidth));
    const actualHeight = Math.max(1, Math.min(newSize.height, maxHeight));

    this.size.resize(actualWidth, actualHeight);
  }
  move(newPosition) {
    let maxX = this.screenSize.width - this.size.width;
    let maxY = this.screenSize.height - this.size.height;

    maxX = Math.max(0, maxX);
    maxY = Math.max(0, maxY);

    let actualX = Math.max(0, Math.min(newPosition.x, maxX));
    let actualY = Math.max(0, Math.min(newPosition.y, maxY));

    this.position.move(actualX, actualY);
  }
}

export function changeWindow(programWindow) {
  const newSize = new Size(400, 300);
  const newPosition = new Position(100, 150);
  
  programWindow.resize(newSize);
  programWindow.move(newPosition);
  
  return programWindow;
}

