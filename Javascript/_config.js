/* */

const c = {
  VERSION: "0",
  LAG: 0,
  LAG2: 0,
  NETWORK: 0,  // 0 => simulate, 1 => Real
  // API
  API_BASE_URL: "https://roguelike-tutorial-javascript.netlify.app/",
  PING_ENDPOINT: "/ping",
  TEST_ENDPOINT: "/test", // download webpage, slow
  OK_ENDPOINT: "/ok", // normal ping
  ENDPOINTS: 2, // 0 => PING 1 => TEST  2 => OK 
  // Render
  RENDER_TYPE: 1, // 0 = ASCII  1 = UNICODE
  CANVAS_NAME: "canvas",
  PPP: 20,
  FONT: ["sans-serif", "arial", "IBM"],
  FONT_SELECTED: 2,
  // Camera dimensions
  CAM_PIXELS_X: window.innerWidth - 300,
  CAM_PIXELS_Y: window.innerHeight - 15,
  CAM_COLS: undefined,
  CAM_ROWS: undefined,
  CAM_DELTA_X: undefined,
  CAM_DELTA_Y: undefined,
  VIEW_COLS: undefined,
  VIEW_ROWS: undefined,
  // Game
  NICK: "",
  TOKEN: "",
  HISTORY: [],
  HISTORY_LINES: 15, // max 15
  LOOT_LINES: 5,
  INIT_DATE: new Date("2097-08-29 02:14:00"),
  MS_PER_TURN: 1000 * 20,
  INDEX_SELECTED: undefined,
  ID_SELECTED: undefined,
  NPC_SELECTED: undefined,
  IS_SERVER_TURN: false,
};

(function autoUpdateC() {
  c.CAM_COLS = Math.floor(c.CAM_PIXELS_X / c.PPP);
  c.CAM_ROWS = Math.floor(c.CAM_PIXELS_Y / c.PPP);
  c.CAM_DELTA_X = Math.floor((c.CAM_PIXELS_X - (c.CAM_COLS * c.PPP)) / 2);
  c.CAM_DELTA_Y = Math.floor((c.CAM_PIXELS_Y - (c.CAM_ROWS * c.PPP)) / 2);
})();

const lib = {
  randomInt: function (min, max) {
    return Math.floor(Math.random() * (max - min + 1) + min);
  },
  currentDate: function (turn) {
    return new Date(c.INIT_DATE.getTime() + c.MS_PER_TURN * turn);
  },
  realDate: function () {
    const d = new Date();
    let aa = d.toTimeString().split(" ")[0];
    let bb = d.toISOString().split("T")[0];
    return aa + "_" + bb;
  },
  atPoint: function (p, es) {
    let resp = [];
    for (let e of es) {
      if (e.pos.x === p.x && e.pos.y === p.y) {
        resp.push(e);
      }
    }
    return resp;
  },
};

export {
  c,
  lib,
};
