var MessageRoom = React.createClass({

  render: function() {
    <div>
      <ContactList className="contact-list" id="contact-list"/>
      <Thread id="thread" class="swpthread"/>
      <MessageForm id="message_compose" class="message-compose"/>
    </div>
  }
});
