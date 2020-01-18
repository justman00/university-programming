const g = {
  //   0: [0, 1, 3, 5],
  1: [2, 4],
  2: [3],
  3: [6, 7],
  4: [9, 10],
  5: [],
  6: [7],
  7: [8],
  8: [],
  9: [],
  10: [11, 12],
  11: [13, 14]
};

class Node {
  constructor(value) {
    this.value = value;
    this.left = null;
    this.right = null;
  }
}

function DFS(startingNode, graph) {
  const stack = [];
  const visited = new Set();

  stack.push(startingNode);

  function helper(s) {
    if (s.length === 0) {
      return;
    }

    const current = s.pop();

    if (!visited.has(current)) {
      console.log(current);
      visited.add(current);

      const verts = graph[current];

      if (verts) {
        for (let vert of verts.reverse()) {
          s.push(vert);
        }
      }
    }

    helper(s);
  }

  helper(stack);
}

BFS(1, g);

function BFS(startingNode, graph) {
    const queue = [];
    const visited = new Set();
  
    queue.push(startingNode);
  
    function helper(q) {
      if (q.length === 0) {
        return;
      }
  
      const current = q.shift();
  
      if (!visited.has(current)) {
        console.log(current);
        visited.add(current);
  
        const verts = graph[current];
  
        if (verts) {
          for (let vert of verts) {
            q.push(vert);
          }
        }
      }
  
      helper(q);
    }
  
    helper(queue);
  }

