var socket;
if (window.location.protocol != "https:") {
  socket = new WebSocket('ws://'+window.location.host+'/websocket?num={{.num}}')
} else {
  socket = new WebSocket('wss://'+window.location.host+'/websocket?num={{.num}}')
}

// Message received on the socket
socket.onmessage = function(event) {
  console.log(event);
  if (event.data != "\"pong\"") {
    var payload = JSON.parse(event.data);
    fireEvent("smsmessage", {message: payload.Message, num: payload.Num, you: false});
    notifyMe(payload.Message, payload.Num);
  }
}

setInterval(function() {
  $.ajax({method: "GET", url: "/ping"}).then(function() {console.log("here", arguments);});
  var ping = {
    MessageType: "ping"
  };

  socket.send(JSON.stringify(ping));
}, 10000);

var SocketController = {

}
