const canvas = document.querySelector('canvas');

canvas.width = window.innerWidth - 5;
canvas.height = window.innerHeight - 6;

const ctx = canvas.getContext('2d');

function drawHouse() {
  ctx.fillStyle = '#FF0000';
  ctx.fillRect(12.5, 230, 175, 175); // height + 105

  // Draw triangle
  ctx.fillStyle = '#A2322E';
  ctx.beginPath();
  ctx.moveTo(12.5, 230);
  ctx.lineTo(185, 230);
  ctx.lineTo(99, 200);
  ctx.closePath();
  ctx.fill();

  //windows
  ctx.fillStyle = '#663300';
  ctx.fillRect(25, 240, 35, 50);
  ctx.fillStyle = '#0000FF';
  ctx.fillRect(27, 242, 13, 23);
  ctx.fillRect(43, 242, 13, 23);
  ctx.fillRect(43, 267, 13, 21);
  ctx.fillRect(27, 267, 13, 21);

  ctx.fillStyle = '#663300';
  ctx.fillRect(145, 240, 35, 50);
  ctx.fillStyle = '#0000FF';
  ctx.fillRect(147, 242, 13, 23);
  ctx.fillRect(163, 242, 13, 23);
  ctx.fillRect(163, 267, 13, 21);
  ctx.fillRect(147, 267, 13, 21);

  //door
  ctx.fillStyle = '#754719';
  ctx.fillRect(140, 358, 30, 47);

  //door knob
  ctx.beginPath();
  ctx.fillStyle = '#F2F2F2';
  ctx.arc(165, 380, 3, 0, 2 * Math.PI);
  ctx.fill();
  ctx.closePath();

  // garage
  ctx.fillStyle = '#555555';
  ctx.fillRect(20, 358, 80, 15);
  ctx.fillRect(20, 374, 80, 15);
  ctx.fillRect(20, 390, 80, 15);

  ctx.fillStyle = '#999999';
  ctx.fillRect(20, 373, 80, 1);
  ctx.fillRect(20, 389, 80, 1);

  //Text on the Right
  ctx.fillStyle = '#fff';
  ctx.font = '20px Veranda';
  ctx.fillText('Acasa', 75, 260);
}

let x = canvas.width - 300;

function drawCar() {
  car = new Image();
  car.src =
    'https://i.pinimg.com/originals/68/69/ec/6869ecc33b91bb166c9a8aeff2eba120.png';

  car.addEventListener('load', animateCar, false);
}

function animateCar() {
  ctx.clearRect(200, 100, canvas.width, canvas.height);
  ctx.drawImage(car, x, 230, 275, 275);
  x -= 4;
}

let sunXCoord = canvas.width - 100;
let sunYCoord = 70;
function drawSun(x = sunXCoord, y = sunYCoord) {
  //draw the sun
  ctx.clearRect(x + 1, 0, canvas.width, y * 2);
  ctx.beginPath();
  ctx.arc(x, y, 30, 0, 2 * Math.PI, false);
  ctx.fillStyle = '#fff1dc';
  ctx.fill();
}

function render() {
  drawHouse();
  drawSun(sunXCoord, sunYCoord);
  drawCar();
}
const framePxRate = canvas.width / (60 * 6); // exact num of pixels until sun is down
function renderAnimation() {
  window.requestAnimationFrame(() => {
    if (x > 250) animateCar();
    sunXCoord -= framePxRate;
    if (sunXCoord >= -40) drawSun();
    renderAnimation();
  });
}

render();
renderAnimation();
