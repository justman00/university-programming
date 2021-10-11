const burgerMenu = document.querySelector('#burger-menu');
const mobileMenu = document.querySelector('#mobile-menu');
const timeContainer = document.querySelector('#time');

function openMobileMenu() {
  burgerMenu.setAttribute('aria-expanded', 'true');
  mobileMenu.classList.remove('hidden');
}

function closeMobileMenu() {
  burgerMenu.setAttribute('aria-expanded', 'false');
  mobileMenu.classList.add('hidden');
}

let timeout;

function afiseazara() {
  const today = new Date();
  const time =
    today.getHours() + ':' + today.getMinutes() + ':' + today.getSeconds();

  timeContainer.innerHTML = time;
  timeContainer.classList.remove('hidden');
}

function startAfisareOra() {
  afiseazara();
  timeout = setInterval(afiseazara, 1000);
}

function stopAfisareOra() {
  if (timeout) {
    clearInterval(timeout);
    timeContainer.classList.add('hidden');
  }
}

function isPhoneNumberValid() {
  const phoneInput = document.querySelector('#phone');
  const hasRightLength = phoneInput.value.length === 12;
  const isMoldovian = phoneInput.value.startsWith('+373');

  return hasRightLength && isMoldovian;
}

const form = document.getElementById('form');

form.addEventListener('submit', validateForm);

function validateForm(e) {
  if (!isPhoneNumberValid()) {
    e.preventDefault();

    const errContainer = document.querySelector('#error');
    errContainer.textContent =
      'Numarul de telefon trebuie sa inceapa cu +373 si sa fie de tip moldovenesc';
    errContainer.classList.remove('hidden');
  }
}
