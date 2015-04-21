var MessageRoom = React.createClass({
  componentDidMount: function() {
    AppSocket.start(this.props.qr_secret);
  },
  render: function() {
    return <div className="swpcontainer">
      <div className="contact-list" id="contact-list"> 
        <ContactList />
      </div>
      <div id="thread" className="swpthread">
        <Thread />
      </div>
      <div id="message_compose" className="message-compose">
        <MessageForm />
      </div>
    </div>;
  }
});
