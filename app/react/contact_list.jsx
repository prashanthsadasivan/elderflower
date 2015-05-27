var ContactList = React.createClass({
  getInitialState: function() {
    return { contacts: {}, selectedNum: ''}
  },
  componentDidMount: function() {
    Dispatcher.register("contactlist", "smsmessage", function(data) {
      this.state.contacts[data.num] = true;
      this.setState({contacts: this.state.contacts});
      console.log("whatupdude", this.state);
    }.bind(this));
    Dispatcher.register("contactlist", "receipient_selected", function(data) {
      this.setState({selectedNum: data.num});
    }.bind(this));
  },
  componentWillUnmount: function() {
    Dispatcher.unregister("contactlist", "smsmessage");
    Dispatcher.unregister("contactlist", "receipient_selected");
  },
  handleNumChange: function(event) {
    Dispatcher.fireEvent("receipient_selected", {num: event.target.value})
  },
  render: function() {
    var self = this;
    return <div className="contact-list">
      {
        Object.keys(this.state.contacts).map(function(num, i) {
          return <Contact key={i} contactNum={num} className={self.state.selectedNum == num ? "contact bold" : "contact"} />
        })}
      <input type="text" placeholder="number" className="num" value={self.state.selectedNum} onChange={this.handleNumChange} />
    </div>
  }
});
