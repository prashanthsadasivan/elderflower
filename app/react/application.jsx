var Application = React.createClass({
  getInitialState: function() {
    return { view: '', qr_secret: ''}
  },
  componentDidMount: function(){
    if (localStorage.qr_secret !== null && localStorage.qr_secret !== undefined) {
      this.setState({view: 'message_room', qr_secret: localStorage.qr_secret});
    } else {
      var self = this;
      Dispatcher.register('application', 'qr_delivery', function(event) {
        self.setState({view: 'qr', qr_secret: event.qr_secret});
      });
      Dispatcher.register('application', 'phone_confirmed', function(event) {
        localStorage.qr_secret = self.state.qr_secret;
        self.setState({view: 'message_room', qr_secret: ''});
      });
      AppSocket.start('');
    }
  },
  render: function() {
    switch (this.state.view) {
      case 'qr':
        return <QrCode qr_secret={this.state.qr_secret}/>;
      case 'message_room':
        return <MessageRoom qr_secret={this.state.qr_secret} />;
      default:
        return <div> LOADING </div>;
    }
  }
});
