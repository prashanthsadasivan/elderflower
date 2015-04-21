var ContactList = React.createClass({
  getInitialState: function() {
    return { contacts: {}, selectedNum: ''}
  },
  componentDidMount: function() {
    window.register("contactlist", "smsmessage", function(data) {
      this.state.contacts[data.num] = true;
      this.setState({contacts: this.state.contacts});
      console.log("whatupdude", this.state);
    }.bind(this));
    window.register("contactlist", "receipient_selected", function(data) {
      this.setState({selectedNum: data.num});
    }.bind(this));
  },
  componentWillUnmount: function() {
    window.unregister("contactlist", "smsmessage");
    window.unregister("contactlist", "receipient_selected");
  },
  handleNumChange: function(event) {
    fireEvent("receipient_selected", {num: event.target.value})
  },
  handleNumClick: function(event) {
    fireEvent("receipient_selected", {num: event.target.getAttribute('value')});
  },
  render: function() {
    var self = this;
    return <div className="contact-list">
      {
        Object.keys(this.state.contacts).map(function(num, i) {
          return <div key={i} onClick={self.handleNumClick} value={num} className={self.state.selectedNum == num ? "contact bold" : "contact"} >
            <p value={num}>{num}</p>
          </div>
        })}
      <input type="text" placeholder="number" className="num" value={self.state.selectedNum} onChange={this.handleNumChange} />
    </div>
  }
});
