var Notifications = {}
Notifications.notifyMe = function(notifBody, sender) {
  if (!Notification) {
    alert('Please us a modern version of Chrome, Firefox, Opera or Firefox.');
    return;
  }
  if (Notification.permission !== "granted")
    Notification.requestPermission();

  var notification = new Notification(sender, {
    icon: 'http://' + window.location.host + '/public/img/squid.png',
      body: notifBody,
  });
}
