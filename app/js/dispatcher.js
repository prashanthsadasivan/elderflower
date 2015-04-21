var Dispatcher = {}

Dispatcher.registeredCallbacks = {}
Dispatcher.register = function(id, event_name, cb) {
  if (Dispatcher.registeredCallbacks[event_name] === undefined || Dispatcher.registeredCallbacks[event_name] === null) {
    Dispatcher.registeredCallbacks[event_name] = {}
  }
  Dispatcher.registeredCallbacks[event_name][id] = cb;
};
Dispatcher.unregister = function(id, event_name) {
  delete Dispatcher.registeredCallbacks[event_name][id];
};
Dispatcher.fireEvent = function(event_name, data) {
  for(var index in Dispatcher.registeredCallbacks[event_name]) {
    if (Dispatcher.registeredCallbacks[event_name].hasOwnProperty(index)) {
      Dispatcher.registeredCallbacks[event_name][index](data, event_name);
    }
  }
};
