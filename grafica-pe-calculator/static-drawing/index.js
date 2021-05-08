const canvas = document.querySelector('canvas');

canvas.width = window.innerWidth - 5;
canvas.height = window.innerHeight - 6;

const ctx = canvas.getContext('2d');

function drawHouse() {
  ctx.fillStyle = '#FF0000';
  ctx.fillRect(12.5, 30, 175, 175); // height + 105

  // Draw triangle
  ctx.fillStyle = '#A2322E';
  ctx.beginPath();
  ctx.moveTo(12.5, 30);
  ctx.lineTo(185, 30);
  ctx.lineTo(99, 0);
  ctx.closePath();
  ctx.fill();

  //windows
  ctx.fillStyle = '#663300';
  ctx.fillRect(25, 40, 35, 50);
  ctx.fillStyle = '#0000FF';
  ctx.fillRect(27, 42, 13, 23);
  ctx.fillRect(43, 42, 13, 23);
  ctx.fillRect(43, 67, 13, 21);
  ctx.fillRect(27, 67, 13, 21);

  ctx.fillStyle = '#663300';
  ctx.fillRect(145, 40, 35, 50);
  ctx.fillStyle = '#0000FF';
  ctx.fillRect(147, 42, 13, 23);
  ctx.fillRect(163, 42, 13, 23);
  ctx.fillRect(163, 67, 13, 21);
  ctx.fillRect(147, 67, 13, 21);

  //door
  ctx.fillStyle = '#754719';
  ctx.fillRect(140, 158, 30, 47);

  //door knob
  ctx.beginPath();
  ctx.fillStyle = '#F2F2F2';
  ctx.arc(165, 180, 3, 0, 2 * Math.PI);
  ctx.fill();
  ctx.closePath();

  // garage
  ctx.fillStyle = '#555555';
  ctx.fillRect(20, 158, 80, 15);
  ctx.fillRect(20, 174, 80, 15);
  ctx.fillRect(20, 190, 80, 15);

  ctx.fillStyle = '#999999';
  ctx.fillRect(20, 173, 80, 1);
  ctx.fillRect(20, 189, 80, 1);

  //Text on the Right
  ctx.fillStyle = '#fff';
  ctx.font = '20px Veranda';
  ctx.fillText('Acasa', 75, 60);
}

let x = canvas.width - 300;

function drawCar() {
  car = new Image();
  car.src =
    'https://i.pinimg.com/originals/68/69/ec/6869ecc33b91bb166c9a8aeff2eba120.png';

  function animateCar() {
    ctx.clearRect(200, 0, canvas.width, canvas.height);
    ctx.drawImage(car, x, 30, 275, 275);
    x -= 4;
    if (x > 250) requestAnimationFrame(animateCar);
  }

  car.addEventListener('load', animateCar, false);
}

drawHouse();
drawCar();
