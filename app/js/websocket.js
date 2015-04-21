var AppSocket = {};

AppSocket.start = function(qr_secret) {
  if (this.socket !== undefined) {
    return;
  }
  if (window.location.protocol != "https:") {
    this.socket = new WebSocket('ws://'+window.location.host+'/websocket?qr_secret=' + qr_secret)
  } else {
    this.socket = new WebSocket('wss://'+window.location.host+'/websocket?qr_secret=' + qr_secret)
  }
  // Message received on the socket
  this.socket.onmessage = function(event) {
    console.log(event);
    var payload = JSON.parse(event.data);
    switch (payload.MessageType) {
      case "pong":
        console.log("pongged");
        return;
      case "SMS/Received":
        Dispatcher.fireEvent("smsmessage", {message: payload.Message, num: payload.Num, you: false});
        Notifications.notifyMe(payload.Message, payload.Num);
        return;
      case "phone_confirmed":
        Dispatcher.fireEvent("phone_confirmed", {});
        return;
      case "qr_delivery":
        Dispatcher.fireEvent("qr_delivery", {qr_secret: payload.Qr_secret});
        return;
    }
  };
  var self = this;
  setInterval(function() {
    $.ajax({method: "GET", url: "/ping"}).then(function() {console.log("here", arguments);});
    var ping = {
      MessageType: "ping"
    };

    self.socket.send(JSON.stringify(ping));
  }, 10000);
}

AppSocket.send = function(event) {
  this.socket.send(event);
};
