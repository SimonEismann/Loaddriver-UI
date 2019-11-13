import Point from "chart.js";

export default class IntensityData {
  constructor(fileName, csv) {
    this.fileName = fileName;
    const data = [];
    csv.replace("\r", "").split("\n").forEach(line => {
      const splitLine = line.split(",");
      const x = parseFloat(splitLine[0]) - 0.5;
      const y = parseFloat(splitLine[1]);
      if (x && y) {
        data.push(new Point2D(x, y));
      }
    })
    this.data = data;
  }
}

class Point2D {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
}