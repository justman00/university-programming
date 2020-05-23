/**
 * Input format:
 * [key] => from node
 * [value] => Array: [0] => from node, [1] => to node, [2] => ponderea
 */
// const input = {
//     1: [1, 2, 4],
//     2: [2, 3, 7],
//     3: null
// }

const inpu2 = [
  [1, 2, 3],
  [1, 3, 7],
  [1, 5, 9],
  [2, 3, 3],
  [2, 4, 3],
  [3, 4, 2],
  [3, 5, 3],
  [3, 7, 5],
  [4, 5, 2],
  [4, 6, 4],
  [4, 8, 7],
  [5, 6, 2],
  [5, 7, 3],
  [5, 8, 5],
  [6, 8, 3],
  [7, 8, 3]
];

const input = [
  [1, 2, 3],
  [1, 3, 5],
  [1, 5, 8],
  [2, 3, 3],
  [2, 4, 3],
  [3, 4, 2],
  [3, 5, 3],
  [3, 7, 5],
  [4, 5, 4],
  [4, 6, 4],
  [4, 8, 7],
  [5, 6, 2],
  [5, 7, 3],
  [5, 8, 5],
  [6, 8, 3],
  [7, 8, 3]
]

function buildNodesMap(startingNode) {
  const allNodes = new Set();

  const map = input.reduce((acc, curr) => {
    const [outgoing, ingoing] = curr;

    if (acc[outgoing] === undefined) {
      allNodes.add(outgoing);

      if (outgoing === startingNode) {
        acc[outgoing] = 0;
      } else {
        acc[outgoing] = Infinity;
      }
    }

    if (acc[ingoing] === undefined) {
      allNodes.add(ingoing);

      if (ingoing === startingNode) {
        acc[ingoing] = 0;
      } else {
        acc[ingoing] = Infinity;
      }
    }

    return acc;
  }, {});

  return [new Array(allNodes), map];
}

function calculateShortestPath(startingNode, endNode) {
  const [allNodes, nodesMap] = buildNodesMap(startingNode);

  let hasChanged = true;

  while (hasChanged) {
    hasChanged = false;

    for (const edge of input) {
      const [outgoing, ingoing, ponderea] = edge;
      const ingoingNodeP = nodesMap[ingoing];
      const outgoingNodeP = nodesMap[outgoing];

      if (ingoingNodeP - outgoingNodeP > ponderea) {
        nodesMap[ingoing] = outgoingNodeP + ponderea;
        hasChanged = true;
      }
    }
  }

  console.log(`Shortest path to ${endNode} is of length: ${nodesMap[endNode]}`);
  getShortestPath(nodesMap, endNode);
}

function getShortestPath(nodesMap, targetNode) {
  const stack = [targetNode];
  const visited = new Set();
  let str = '';

  while (stack.length) {
    const currentNode = stack.pop();

    if (visited.has(currentNode)) {
      continue;
    } else {
      visited.add(currentNode);
    }

    str += `To ${currentNode} go: `;

    input.forEach(edge => {
      const [outgoing, ingoing, ponderea] = edge;

      if (ingoing === currentNode) {
        if (nodesMap[currentNode] - nodesMap[outgoing] === ponderea) {
          str = str + outgoing + ' ';

          stack.push(outgoing);
        }                                                                                                                                                                                                                       
      }
    });

    str += '\n';
  }

  console.log(str);
}

calculateShortestPath(1, 8);
