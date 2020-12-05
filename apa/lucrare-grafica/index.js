const CliGraph = require('cli-graph');

const constant_time_operation = () => 1 + 1;

// O(log(n))
const alg_one = (n) => {
  let i = 0;
  let h = 1;

  while (h <= n) {
    constant_time_operation();
    n = n / 2;
    i++;
  }

  return i;
};

// O(8^n)
const alg_two = (n) => {
  let iter = 0;
  let k = 1;

  for (let i = 0; i < n - 1; i++) {
    k = 8 * k;
    iter++;

    for (let j = 1; j < k; j++) {
      constant_time_operation();
      iter++;
    }
  }

  return iter;
};

// O(n)
const alg_three = (n) => {
  let iter = 0;
  let i = 2;

  while (i <= n) {
    constant_time_operation();
    let j = 2 * i;
    iter++;

    while (j <= n) {
      constant_time_operation();
      j = j + i;
      iter++;
    }

    i++;
  }

  return iter;
};

// O(n * log(n))
const alg_four = (n) => {
  let iter = 0;
  let i = 2;

  while (i <= n) {
    constant_time_operation();
    let j = 1;
    iter++;

    while (j <= n) {
      constant_time_operation();
      j = 2 * j;
      iter++;
    }

    i++;
  }

  return iter;
};

function main() {
  const dataSets = [2, 3, 4, 5];
  const algos = [alg_one, alg_two, alg_three, alg_four];

  algos.forEach((alg, i) => {
    const graph = new CliGraph({
      height: 20,
      width: 8,
      center: {
        y: 19,
        x: 3,
      },
    });
    dataSets.forEach((dataSet) => {
      const data = alg(dataSet);

      graph.addPoint(dataSet, data);
    });

    console.log(`This is the ${i + 1} algorithm`);
    console.log(graph.toString());
  });
}

main();
