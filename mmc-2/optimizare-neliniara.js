const A = 1;
const B = 6;

// ALGO

const functiaInitiala = (x1, x2) =>
  A * Math.pow(x1, 2) + 2 * x1 * x2 + B * Math.pow(x2, 2) - 2 * x1 - 3 * x2;
const grad1 = (x1, x2) => 2 * x1 + 2 * x2 - 2;
const grad2 = (x1, x2) => 2 * x1 + 12 * x2 - 3;

const dubluModul = (a, b) =>
  Math.sqrt(Math.pow(a, 2) + Math.pow(b, 2)).toFixed(6);

const err = Math.pow(10, -5);

const main = () => {
  // se ia o valoare arbitrara fie x1 si x2 care reprezinta 2 puncte ale functiei cu 2 varibile
  let x1 = 0;
  let x2 = 0;
  const valoareaFunctieiInitiale = functiaInitiala(x1, x2);
  // se calculeaza valoarea gradientului functiei pentru acel punct
  let valG1 = grad1(x1, x2);
  let valG2 = grad2(x1, x2);
  // se calculeaza si initializeaza valoarea gradientul functiei
  let comparabila = dubluModul(valG1, valG2);
  let pasul = 1;
  // de adaugat conditie in baza de iter daca depaseste 100000 de iteratii
  let iter = 1;
  // se intra intr-un ciclu cu conditia valoarea gradientului < Err, unde err e 10^-5
  while (!(comparabila <= err) || iter > 100) {
    // daca da, atunci se gasesc puncte noi
    let newX1, newX2;
    while (!newX1 && !newX2) {
      iter++;
      const direction = comparabila > 0 ? -1 : 1;
      const newX1Tentative = x1 + pasul * valG1 * direction;
      const newX2Tentative = x2 + pasul * valG2 * direction;

      const valFuncitieNoi = functiaInitiala(newX1Tentative, newX2Tentative);

      if (valFuncitieNoi < valoareaFunctieiInitiale) {
        const newValG1 = grad1(newX1Tentative, newX2Tentative);
        const newValG2 = grad2(newX1Tentative, newX2Tentative);
        comparabila = dubluModul(newValG1, newValG2);
        x1 = newX1Tentative;
        x2 = newX2Tentative;
        valG1 = newValG1;
        valG2 = newValG2;
        break;
      } else {
        pasul *= 0.5;
      }
    }
  }
  console.log(`Acesta sunt punctul de minim al functiei: ${x1} si ${x2}\n`);
  console.log(`Iter: ${iter}`);
  console.log(
    `Valorea functiei pentru punctele de minim: ${functiaInitiala(x1, x2)}`
  );
};

main();
