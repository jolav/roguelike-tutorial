/* */

console.log('Loading.....index.js');

import { } from "./wasm_exec.js";
//import * as _ from "./wasm_exec.js";
import { C } from "./_config.js";
import { api } from "./api.js";

let loadedWASM = false;

const init = {
  loadWASM: async function () {
    document.getElementById("version").innerText = "LOADING...";
    document.getElementById("nick").innerText = "LOADING...";

    // eslint-disable-next-line no-undef
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(
      fetch("rogue.wasm"),
      go.importObject
    );
    go.run(result.instance);
    loadedWASM = true;
  },
};

window.addEventListener("load", init.loadWASM);

window.allInitialized = function () {
  if (loadedWASM === true) {
    const text = window.getApiVersion();
    document.getElementById("version").innerText = text;
    api.newTurn();
  }
};

