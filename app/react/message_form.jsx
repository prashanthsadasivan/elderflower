var MessageForm = React.createClass({
  getInitialState: function() {
    return {MessageType: 'SMS/Send', Message: '', Num: ''};
  },
  componentDidMount: function() {
    Dispatcher.register("messageform", "receipient_selected", function(data) {
      this.setState({Num: data.num});
    }.bind(this));
  },
  componentWillUnmount: function() {
    Dispatcher.unregister("messageform", "receipient_selected");
  },
  handleMessageChange: function(event) {
    this.setState({Message: event.target.value});
  },
  handleKeyUp: function(event) {
    if (event.which === 13 && this.state.Message != "" && this.state.Message != null && this.state.Message != undefined) {
      this.handleSendClick();
      Dispatcher.fireEvent("smsmessage", {message: this.state.Message, num: this.state.Num, you: true});
    }
  },
  handleSendClick: function(event) {
    console.log(this.state);
    AppSocket.send(JSON.stringify(this.state))
    this.setState({Message: ''});
  },
  render: function() {
    return <input type="text" placeholder="message" onKeyUp={this.handleKeyUp} className="messageInput" value={this.state.Message} onChange={this.handleMessageChange} />;
  }
});
