// O(2^n)
function helper(n) {
  let iter = 0;

  const result = (function fib1(n) {
    iter++;
    if (n < 2) return n;
    return fib1(n - 1) + fib1(n - 2);
  })(n);

  console.log(`Rezultatul functiei: I ${result}`);

  return iter;
}

// O(n)
function fib2(n) {
  let iter = 0;
  let i = 1;
  let j = 0;

  for (let k = 1; k <= n; k++) {
    iter++;
    j = i + j;
    i = j - i;
  }

  console.log(`Rezultatul functiei II: ${j}`);

  return iter;
}

// O(log(n))
function fib3(n) {
  let i = 1;
  let h = 1;
  let j = 0;
  let k = 0;
  let iter = 0;

  while (n > 0) {
    let t = 0;
    iter++;
    if (n % 2 === 1) {
      t = j * h;
      j = i * h + j * k + t;
      i = i * k + t;
    }

    t = Math.pow(h, 2);
    h = 2 * k * h + t;
    k = Math.pow(k, 2) + t;
    n = Math.floor(n / 2);
  }

  console.log(`Rezultatul functiei III: ${j}`);

  return iter;
}

function main() {
  const algos = [helper, fib2, fib3];
  const dataSet = [10, 20, 30, 40];

  const algoResults = algos.map((alg) => {
    const results = dataSet.map((data) => {
      const iters = alg(data);

      return iters;
    });

    return results;
  });

  console.table(algoResults);
}

main();
