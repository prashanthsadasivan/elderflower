var Application = React.createClass({
  getInitialState: function() {
    return { view: 'qr'}
  },
  render: function() {
    switch (this.state.view) {
      case 'login':
        return <div> </div>;
      case 'qr':
        return <QrCode />;
      case 'message_room':
        return <MessageRoom />;
    }
  }
});
