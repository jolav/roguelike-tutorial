/* */

console.log('Loading.....api.js');

const api = {
  newTurn: function () {
    const t = window.getNewTurn();
    console.log(t);
    document.getElementById("nick").innerText = t.nick;
  }
};

export {
  api
};

