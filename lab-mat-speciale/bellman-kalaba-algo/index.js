import * as utils from './utils';
import nodes from './input';

const size = utils.getMax(nodes);
const count = size + 1;
let matrix = utils.getDefaultMatrix(size);
let allNodes = Object.values(nodes);

/**
 * Gets a node that matches Base and Target indexes
 * @param i Base index
 * @param j Target index
 * @returns {number[]}
 */
const getNode = (i, j) => {
  for (let iterator = 0; iterator < allNodes.length; iterator++) {
    if (allNodes[iterator][0] === i && allNodes[iterator][1] === j) {
      return allNodes[iterator];
    }
  }
};

/**
 * Finds the smallest sum result between current (v) and incremental(i) rows
 * @param matrix
 * @param v current matrix row to be summed
 * @param i incremental matrix row to be summed
 * @returns {number} the smallest sum found
 */
const findTheSmallestSum = (matrix, v, i) => {
  let temp = [];
  // Extract sum between V row and i in a tempArray
  for (let j = 0; j < count; j++) {
    temp.push(matrix[v][j] + matrix[i][j]);
  }
  // Find min in sums
  let min = temp[0];
  for (let j = 0; j < temp.length; j++) {
    if (temp[j] < min) min = temp[j];
  }
  return min;
};

// Step 1
for (let i = 0; i < matrix.length; i++) {
  for (let j = 0; j < count; j++) {
    // Pot ZERO as value for main diagonal
    if (i === j) {
      matrix[i][j] = 0;
      continue;
    }
    // Check node
    const node = getNode(i, j);
    if (node) {
      // Fill matrix with node's amount
      matrix[i][j] = getNode(i, j)[2];
    } else {
      // Put Infinity otherwise
      matrix[i][j] = Infinity;
    }
  }
}

console.log('After Step 1');
utils.display(matrix);

// Step 2
matrix.push(new Array(count));
const v0Index = matrix.length - 1;
for (let i = 0; i < count; i++) {
  // Append last column as new row
  matrix[v0Index][i] = matrix[i][count - 1];
}

console.log('After Step 2');
utils.display(matrix, count);

// Step 3
let isDone = false;
do {
  let v = matrix.length - 1;
  let prevArray = [...matrix[v]];
  matrix.push(new Array(count));
  for (let i = 0; i < count; i++) {
    matrix[matrix.length - 1][i] = findTheSmallestSum(matrix, v, i);
  }
  isDone = utils.compareArrays(prevArray, matrix[matrix.length - 1]);
} while (!isDone);

const result = matrix[matrix.length - 1];

console.log('After Step 3');
utils.display(matrix, count);
let res = '';

const stack = [];

stack.push([result[0], 0]);

while (stack.length > 0) {
  const [roadLength, i] = stack.pop();
  const arcs = allNodes.filter(node => node[0] === i);

  if (arcs.length === 0) {
    res += i;
    break;
  }

  arcs.forEach((arc, index) => {
    const weight = arc[2];
    const neighbour = arc[1];

    if (roadLength - result[neighbour] === weight) {
      res += i;
      const neighbourId = result.findIndex(num => num === result[neighbour]);
      stack.push([result[neighbour], neighbourId]);
    }
  });
}

console.log(res);

console.log(`The shortest path: ${matrix[matrix.length - 1][0]}`);
