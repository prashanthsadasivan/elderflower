var Thread = React.createClass({
  getInitialState: function() {
    return { Messages: [], selectedNum: '' };
  },
  componentDidMount: function() {
    Dispatcher.register("thread", "smsmessage", function(data) {
      this.setState({ Messages: this.state.Messages.concat([{ Message: data.message, Num: data.num, You: data.you }])});
    }.bind(this));
    Dispatcher.register("thread", "receipient_selected", function(data) {
      this.setState({selectedNum: data.num});
    }.bind(this));
  },
  componentWillUnmount: function() {
    Dispatcher.unregister("thread", "smsmessage");
    Dispatcher.unregister("thread", "receipient_selected");
  },
  componentDidUpdate: function() {
    $('.swpthread').scrollTo('max');
  },
  render: function() {
    var self = this;
    return <div>
      { this.state.Messages.filter(function(message) { return message.Num == self.state.selectedNum }).map(function(message, i) {
          return <MessageRow key={i} Message={message.Message} Num={message.Num} You={message.You} />
        })}
    </div>
  }
});
