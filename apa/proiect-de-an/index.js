const nr1 = require('./produse_nr_1.json');
const pegas = require('./produse_pegas.json');

const arrSum = (a, b) => a + b.pret;

function sum(map) {
  return Object.values(map).reduce((acc, curr) => {
    return acc + curr.reduce(arrSum, 0);
  }, 0);
}

const availableProducts = {
  nr1,
  pegas,
};

const shoppingCart = [
  {
    id: 1, // unt
  },
  {
    id: 2, // lapte
  },
  {
    id: 3, // orez
  },
];

// O(m*n^2)
function minMinAlgo(Ni, N) {
  const shopIds = Object.keys(Ni);
  const result = shopIds.reduce((acc, curr) => {
    acc[curr] = [];
    return acc;
  }, {});
  let auxN = [...N];

  while (auxN.length != 0) {
    let min = Infinity;
    let auxItem = null;
    let auxShopId = null;

    for (let shopId of shopIds) {
      for (let productShoppingCart of auxN) {
        const product = Ni[shopId].find(
          (prd) => prd.id === productShoppingCart.id
        );
        result[shopId].push(product);
        const auxSum = sum(result);

        if (auxSum < min) {
          min = auxSum;
          auxItem = product;
          auxShopId = shopId;
        }

        result[shopId].pop();
      }
    }

    result[auxShopId].push(auxItem);
    Ni[auxShopId] = Ni[auxShopId].filter((prod) => {
      return prod.id !== auxItem.id;
    });
    auxN = auxN.filter((prod) => {
      return prod.id !== auxItem.id;
    });
  }

  return result;
}

const res = minMinAlgo(availableProducts, shoppingCart);

console.log('res', res);
