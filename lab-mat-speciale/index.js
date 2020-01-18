const g = {
//   0: [0, 1, 3, 5],
  1: [1, 2, 3, 4],
  2: [4, 7],
  3: [6],
  4: [5, 6, 7],
  5: [],
  6: [9],
  7: [8, 9, 10],
  8: [10],
  9: [4, 10],
  10: []
};

module.exports = g

function buildMatrixes(graph) {
  const incidentMatrix = buildMatrix(graph, 2, -1);
  const adjacentMatrix = buildMatrix(graph, 1, 0);


  console.log(incidentMatrix)
  console.log('\n----------------\n')
  console.log(adjacentMatrix);
}

function buildMatrix(graph, sameArcValue, notItselfValue, sameAndFinalValue = 0) {
    const keys = Object.keys(graph)
    .map(Number)
    .sort((a, b) => a - b);

  const matrix = [];

  for (let key of keys) {
    const row = [];
    const arcs = graph[key];

    for (let k of keys) {
      const vert = arcs.find(arc => arc == k);
      const isEnclosed = vert === key;
      const isSelf = k === key;

      if (isSelf) {
        if (isEnclosed) {
          row[k] = sameArcValue;
        } else {
            if(arcs.length > 0) {
                row[k] = notItselfValue;
            } else {
                row[k] = sameAndFinalValue;
            }

        }
      } else {
        if (Number.isInteger(vert)) {
          row[k] = 1;
        } else {
          row[k] = 0;
        }
      }
    }

    matrix[key] = row;
  }

  return matrix;
}

console.log(buildMatrixes(g));
