import './style.css';
import './app.css';

import logo from './assets/images/shitgo.png';
import {
    CreateNewWorld,
    ExistingCheck,
    ListWorlds,
    LoadGameFromString,
    SaveGame,
    Shutdown
} from "../wailsjs/go/main/App";

let lastGame = localStorage.getItem("lastGame");

let mainPage = `
    <img id="logo" class="logo">
      <div class="Title">Welcome To Dungeons And Dungeons!</div>
      <div>
        <button class="MainButtons" id="contButton" onclick="continuePrevious()">Continue</button>
        </div>
        <div>
        <button class="MainButtons" onclick="loadGame()">Load Game</button>
        <button class="MainButtons" onclick="newGame()">New Game</button>
        </div>
        <div>
        <button class="MainButtons" onclick="gameSettings()">Settings</button>
        </div>
        <div>
        <button class="QuitButton" onclick="quitGame()">Quit Game</button>
      </div>
`;
document.querySelector('#app').innerHTML = mainPage;
document.getElementById('logo').src = logo;
document.getElementById("contButton").innerText = "Continue: "+lastGame.replace('.dat', '');
localStorage.clear();
localStorage.setItem("lastGame", lastGame);

window.continuePrevious = async function () {
    await sleep(75); // The "Coin-Delay Placebo" Principle
    await loadWorld(localStorage.getItem("lastGame"));
}

window.loadGame = async function () {
    await sleep(75); // The "Coin-Delay Placebo" Principle
    await ListWorlds().then((result) => {document.querySelector('#app').innerHTML = result});
    document.getElementById("failAlert").style.display = "none";
    document.getElementById("successAlert").style.display = "none";
}

window.newGame = async function () {
    let b = document.getElementById("worldSeed");
    let c = document.getElementById("worldDiff");
    if (b !== null) {
        localStorage.setItem("worldSeed", b.value);
    }
    if (c !== null) {
        localStorage.setItem("worldDiff", c.value);
    }
    let size = document.getElementById("worldSize");
    if (c !== null) {
        localStorage.setItem("worldSize", size.value);
    }

    await sleep(75); // The "Coin-Delay Placebo" Principle
    document.querySelector('#app').innerHTML = `    <img id="logo" class="logo">
    <div>
      <p>New Game</p>
    </div>
    <div>
      <p class="input">World Name</p>
      <input class="input" id="worldName" type="text" autocomplete="off" />
    </div>
    <div>
      <p class="input">Character Name</p>
      <input class="input" id="playerName" type="text" autocomplete="off" />
    </div>
    <div>
      <label for="character-class">Character Class</label>
      <select name="player-class" id="playerClass">
        <option value="" disabled selected hidden>Choose a Class</option>
        <option value="1">Horny Fucker</option>
        <option value="2">Die of 1d4 Damage</option>
        <option value="3">Muscles For Brains</option>
        <option value="4">Forget About It</option>
      </select>
    </div>
    <div>
      <button class="MainButtons" onclick="worldOptions()">World Options</button>
    </div>
    <div>
      <button class="MainButtons" onclick="finishNewGame()">Create New World</button>
    </div>
    <div>
      <button class="MainButtons" onclick="back()">Back</button>
    </div>
    <div class="Alert" id="existingAlert">
  <span class="closebtn" onclick="this.parentElement.style.display='none';">&times; A World With This Name Already Exists.</span>
</div>
`;
    document.getElementById("existingAlert").style.display = 'none'
    if (localStorage.getItem("worldName") !== null) {
        document.getElementById("worldName").value = localStorage.getItem("worldName")
    }
    if (localStorage.getItem("playerName") !== null) {
        document.getElementById("playerName").value = localStorage.getItem("playerName")
    }
    if (localStorage.getItem("playerClass") !== null) {
        document.getElementById("playerClass").value = localStorage.getItem("playerClass")
    }
}

window.worldOptions = async function () {
    localStorage.setItem("worldName", document.getElementById("worldName").value);
    localStorage.setItem("playerName", document.getElementById("playerName").value);
    localStorage.setItem("playerClass", document.getElementById("playerClass").value);

    await sleep(75); // The "Coin-Delay Placebo" Principle
    document.querySelector('#app').innerHTML = `
    <div>
      <p>World Options</p>
    </div>
    <div>
        <button id="worldType" class="MainButtons" onclick="gameMode()">Story Mode</button>
        <input class="input" id="worldSeed" type="number" autocomplete="off" placeholder="World Seed"/>
    </div>
    <div>
      <select name="game-difficulty" id="worldDiff" >
        <option value="1">Easy</option>
        <option value="2" selected>Normal</option>
        <option value="3">Hard</option>
      </select>
      <select name="world-size" id="worldSize" >
        <option value="1">Tiny (20 Floors)</option>
        <option value="2" selected>Small (40 Floors)</option>
        <option value="3">Normal (50 Floors)</option>
        <option value="4">Huge (100 Floors)</option>
        <option value="5">Massive (150 Floors)</option>
        <option value="6">Unreasonable (1000 Floors)</option>
      </select>
    </div>
    <div>
      <button class="MainButtons" onclick="newGame()">Back</button>
    </div>
`;

    let a = document.getElementById("worldType");
    if (localStorage.getItem("gameType") === "0") {
        a.innerText = "Story Mode";
    }
    if (localStorage.getItem("gameType") === "1") {
        a.innerText = "Infinite Mode";
    }
    if (localStorage.getItem("worldDiff") === null) {
        localStorage.setItem("worldDiff", "2")
    }
    if (localStorage.getItem("worldSize") === null) {
        localStorage.setItem("worldSize", "2")
    }
    let b = document.getElementById("worldSeed");
    b.value = localStorage.getItem("worldSeed");
    let c = document.getElementById("worldDiff");
    c.value = localStorage.getItem("worldDiff");
    let size = document.getElementById("worldSize");
    size.value = localStorage.getItem("worldSize");
}

window.gameMode = async function () {
    await sleep(75)
    if (localStorage.getItem("gameType") === "0") {
        localStorage.setItem("gameType", "1")
        document.getElementById("worldType").innerText = "Infinite Mode";
    } else if (localStorage.getItem("gameType") === "1") {
        localStorage.setItem("gameType", "0")
        document.getElementById("worldType").innerText = "Story Mode";
    } else {
        localStorage.setItem("gameType", "1")
        document.getElementById("worldType").innerText = "Infinite Mode";
    }
}

window.finishNewGame = async function () {
    localStorage.setItem("worldName", document.getElementById("worldName").value)
    localStorage.setItem("playerName", document.getElementById("playerName").value)
    localStorage.setItem("playerClass", document.getElementById("playerClass").value)

    if (localStorage.getItem("worldDiff") === null) {
        localStorage.setItem("worldDiff", "2")
    }
    if (localStorage.getItem("worldSize") === null) {
        localStorage.setItem("worldSize", "2")
    }
    if (localStorage.getItem("worldSeed") === "") {
        localStorage.setItem("worldSeed", "345678987654567898654")
    }
    if (localStorage.getItem("worldType") === null) {
        localStorage.setItem("worldType", "1")
    }
    if (document.getElementById("playerName").value === "") {
        localStorage.setItem("playerName", "Timothy The Forgotten")
    }
    if (document.getElementById("worldName").value === "") {
        localStorage.setItem("worldName", "The Nameless Cavern")
    }
    if (document.getElementById("playerClass").value === "") {
        localStorage.setItem("playerClass", "1")
    }
    await ExistingCheck(localStorage.getItem("worldName"))
        .catch((err) => {
            document.getElementById("existingAlert").style.display = 'block';
            console.log(localStorage.getItem("worldName"));
        }).then(async (err) => {
            if (err === null) {
                document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
      <div class="input-box">
        <p id="output">Loading...</p>
      </div>
`;
                await CreateNewWorld(localStorage.getItem("worldName"), localStorage.getItem("worldSeed"), localStorage.getItem("worldType"), localStorage.getItem("worldDiff"), localStorage.getItem("worldSize"), localStorage.getItem("playerName"), localStorage.getItem("playerClass"))
                    .then((result) => {
                        SaveGame(result);
                    })
            }

        });



}

window.gameSettings = async function () {
    await sleep(75); // The "Coin-Delay Placebo" Principle
    document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
      <div class="input-box" id="input">
        <button class="btn" onclick="back()">Back</button>
      </div>
`;
}

window.quitGame = function () {
    void Shutdown();
}

window.back = async function () {
    await sleep(75); // The "Coin-Delay Placebo" Principle
    let lastGame = localStorage.getItem("lastGame")
    document.querySelector('#app').innerHTML = mainPage;
    document.getElementById('logo').src = logo;
    document.getElementById("contButton").innerText = "Continue: "+lastGame.replace('.dat', '');
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

let dummyData = `[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]`;

window.loadWorld = async function (worldName) {
    let temp
    await LoadGameFromString(worldName).catch((err) => {
        document.getElementById("failAlert").style.display = "block"
        console.error(err);
    }).then((result) => temp = result)
    localStorage.setItem("lastGame", worldName)
    document.querySelector('#app').innerHTML = `<canvas id="gameCanvas" width="200" height="100" style="border:1px solid #000000;">
</canvas>`;
    let canvas = document.getElementById("gameCanvas");
    const ctx = canvas.getContext("2d")

    canvas.style.width = window.innerWidth - 2;
    canvas.style.height = window.innerHeight - 2;
    canvas.width = window.innerWidth - 2;
    canvas.height = window.innerHeight - 2;


    window.addEventListener("resize", (event) => {
        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;
    });
    canvas.addEventListener("mousedown", (el) => {
        let interval = 50;
        //Draw Top-Down
        for  (let i = 0; i < canvas.width/interval; i++) {
            ctx.beginPath();
            ctx.moveTo(interval * i, 0);
            ctx.lineTo(interval * i, canvas.height);
            ctx.stroke();
            ctx.closePath()
        }
        //Draw Top-Right
        for  (let i = 0; i < canvas.width/interval; i++) {
            ctx.beginPath();
            ctx.moveTo(interval * i, 0);
            let temp = canvas.height / interval
            ctx.lineTo((interval*i) + (temp * (interval) * 2), canvas.height);
            ctx.stroke();
            ctx.closePath()
        }
        //Draw Top-Left
        for  (let i = 0; i < canvas.width/interval; i++) {
            ctx.beginPath();
            ctx.moveTo(interval * i, 0);
            let temp = canvas.height / intervalwails
            ctx.lineTo((interval*i) - (temp * (interval) * 2), canvas.height);
            ctx.stroke();
            ctx.closePath()
        }
    });
}

