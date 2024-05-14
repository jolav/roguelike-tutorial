/* */

console.log('Loading.....render.js');

import { C } from "./_config.js";

const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");
ctx.textBaseline = "top";
ctx.textAlign = "center";
ctx.font = C.PPP + "px " + C.FONT_SELECTED;

//console.log(ctx.font);
//const draw

function ascii() {
  console.log('ASCII');
}

const draw = {};

export {
  ascii
};
