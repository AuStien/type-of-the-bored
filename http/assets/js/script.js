const textParagraph = document.getElementById("text-paragraph");
const letters = document.getElementsByClassName("letter");

let hasStarted = false;
let startTime = undefined;
let wordsTyped = 0;
let isAllCorrect = false;
let intervalID = undefined;

addEventListener("keydown", onChange);

function next() {
  htmx.trigger("#text-container", "manual");
}

function reset() {
  hasStarted = false;
  startTime = undefined;
  wordsTyped = 0;
  isAllCorrect = false;
  intervalID = undefined;
}

function startInterval() {
  return setInterval(() => {
    if (hasStarted) {
      const time = Date.now() - startTime;
      const wpm = Math.round(wordsTyped / (time / 1000 / 60));
      document.getElementById("wpm").innerText = wpm;
    }
  }, 500);
}

function onChange(e) {
  if (e.key === "Enter" && isAllCorrect) {
    next();
  }
  if ((e.key.length != 1 && e.key != "Backspace") || isAllCorrect) {
    return;
  }

  // Get index of active letter
  let i = 0;
  for (const letter of letters) {
    if (letter.id === "active") {
      break;
    }
    i++;
  }

  if (e.key === "Backspace") {
    if (i === letters.length - 1) {
      letters[i].classList.remove("correct");
      letters[i].classList.remove("incorrect");
      letters[i].classList.remove("incorrect-space");
    }
    if (i > 0) {
      letters[i - 1].classList.remove("correct");
      letters[i - 1].classList.remove("incorrect");
      letters[i - 1].classList.remove("incorrect-space");
      letters[i].id = "";
      letters[i - 1].id = "active";
      if (letters[i].innerText === " ") {
        wordsTyped--;
      }
    }
    if (i === 1) {
      clearInterval(intervalID);
      hasStarted = false;
      startTime = undefined;
      wordsTyped = 0;
      document.getElementById("wpm").innerText = "0";
    }
    return;
  }

  if (!hasStarted) {
    hasStarted = true;
    startTime = Date.now();
    intervalID = startInterval();
  }

  // Handle current letter
  const currentLetter = letters[i];
  if (e.key === currentLetter.innerText) {
    currentLetter.classList.add("correct");
    if (i < letters.length - 1) {
      currentLetter.id = "";
      letters[i + 1].id = "active";
    } else {
      letters[i].classList.remove("incorrect");
      letters[i].classList.remove("incorrect-space");
    }

    if (e.key === " " || e.key === ".") {
      wordsTyped++;
    }
  } else {
    currentLetter.classList.add("incorrect");
    if (currentLetter.innerText === " ") {
      currentLetter.classList.add("incorrect-space");
    }
    if (i < letters.length - 1) {
      currentLetter.id = "";
      letters[i + 1].id = "active";
    } else {
      letters[i].classList.remove("correct");
    }
    return;
  }

  // If on last letter, check if game is won
  if (i === letters.length - 1) {
    let checkAllCorrect = true;
    for (const l of letters) {
      if (!l.classList.contains("correct")) {
        checkAllCorrect = false;
        break;
      }
    }

    if (checkAllCorrect) {
      isAllCorrect = true;
      clearInterval(intervalID);
      const time = Date.now() - startTime;
      const wpm = Math.round(wordsTyped / (time / 1000 / 60));
      document.getElementById("wpm").innerText = wpm;
      document.getElementById("success-message").classList.remove("hidden");
    }
  }
}
