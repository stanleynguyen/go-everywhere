const loc = window.location;
let wsURI = '';
if (loc.protocol === 'https:') {
  wsURI = 'wss:';
} else {
  wsURI = 'ws:';
}
wsURI += '//' + loc.host;
wsURI += loc.pathname + 'led';
const light = document.getElementById('light');
const socket = new WebSocket(wsURI);
socket.onopen = () => {
  console.log('Connected');
};
socket.onmessage = e => {
  console.log(e.data);
  light.innerHTML = e.data;
};
function send() {
  socket.send('switch');
}
