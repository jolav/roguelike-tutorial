/* */

console.log('Loading....._config.js');

const C = {
  VERSION: undefined,
  LAG: 0,
  // canvas
  CANVAS_NAME: "canvas",
  FONTS: ["sans-serif", "arial", "IBM"],
  FONT_SELECTED: 2,
  PPP: 20,
  // view
  VIEW_COLS: undefined,
  VIEW_ROWS: undefined,
  VIEW_DELTA_X: undefined,
  VIEW_DELTA_Y: undefined,
  PANEL_WIDTH: 300,
  VERTICAL_SAFETY_DISTANCE: 20,
};

(function autoUpdateView() {
  const view_pixels_X = window.innerWidth - C.PANEL_WIDTH;
  const view_pixels_Y = window.innerHeight - C.VERTICAL_SAFETY_DISTANCE;
  C.VIEW_COLS = Math.floor(view_pixels_X / C.PPP);
  C.VIEW_ROWS = Math.floor(view_pixels_Y / C.PPP);
  C.VIEW_DELTA_X = Math.floor((view_pixels_X - (C.VIEW_COLS * C.PPP)) / 2);
  C.VIEW_DELTA_Y = Math.floor((view_pixels_Y - (C.VIEW_ROWS * C.PPP)) / 2);
})();

export {
  C
};

window.getViewSize = function () {
  return C.VIEW_COLS + "_" + C.VIEW_ROWS;
};
