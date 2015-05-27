var Contact = React.createClass({
  getInitialState: function() {
    return { selected: false }
  },
  componentDidMount: function() {
    Dispatcher.register("contactrow" + this.props.contactNum, "receipient_selected", function(data) {
      this.setState({selected: data.num == this.props.contactNum});
    }.bind(this));
  },
  handleClick: function(event) {
    Dispatcher.fireEvent("receipient_selected", {num: this.props.contactNum});
  },
  componentWillUnmount: function() {
    Dispatcher.unregister("contactrow" + this.props.contactNum, "receipient_selected");
  },
  render: function() {
    return <div onClick={this.handleClick} className={this.state.selected ? "contact bold" : "contact"}><p> {this.props.contactNum} </p></div>;
  }
});
