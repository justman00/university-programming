const input = require('./input')

function bfs(graph, source, final, parent) {
  const visited = [];
  const queue = [];
  const length = graph.length

  // Marcam toti verticii ca nevizitati
  for (let i = 0; i < length; i++) {
    visited[i] = false;
  }

  // Adaugam sursa la Queue si o marcam ca visitata
  queue.push(source);
  visited[source] = true;
  parent[source] = -1;

  while (queue.length !== 0) {
    const currentVertice = queue.shift();
    for (let j = 0; j < length; j++) {
      // cautam urmatorii vertici nevizitati care pot fi supliniti de verticele curent
      if (visited[j] === false && graph[currentVertice][j] > 0) {
        queue.push(j);
        parent[j] = currentVertice;
        visited[j] = true;
      }
    }
  }
  // daca am ajuns la ultimul vertice incepand de la sursa, returneaza true, daca nu pai false
  return visited[final] == true;
}

function fordFulkerson(graph, source, final) {
  // recreem graphul pentru a avea immutability
  const localGraph = graph.map((verts) => {
    return [...verts];
  });
  const parent = [];
  let maxFlow = 0;

  // atat timp cat exista o cale de a ajunge de la sura pana la final
  while (bfs(localGraph, source, final, parent)) {
	let pathFlow = Infinity;

	// parcurgem graphul si setam fluxul de la vertice la vertice
    for (let v = final; v !== source; v = parent[v]) {
      console.log(parent, v, source, final, parent[v]);
      u = parent[v];
      pathFlow = Math.min(pathFlow, localGraph[u][v]);
    }

    for (v = final; v !== source; v = parent[v]) {
      u = parent[v];
      localGraph[u][v] -= pathFlow;
      localGraph[v][u] += pathFlow;
    }

    maxFlow += pathFlow;
  }
  // Return the overall flow
  return maxFlow;
};

console.log(fordFulkerson(input, 0, 5))