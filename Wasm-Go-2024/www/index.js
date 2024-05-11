/* */

console.log('Loading.... index.js');

import { } from "./wasm_exec.js";
//import * as _ from "./wasm_exec.js";

const init = {
  loadWASM: async function () {
    document.getElementById("hi").innerText = "LOADING...";

    // eslint-disable-next-line no-undef
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("rogue.wasm"), go.importObject).then((result) => {
      go.run(result.instance);
      init.init();
    });
  },
  init: function () {
    console.log('Init');
    // eslint-disable-next-line no-undef
    const text = hiFromWASM();
    document.getElementById("hi").innerText = text;
  }
};

window.addEventListener("load", init.loadWASM);

