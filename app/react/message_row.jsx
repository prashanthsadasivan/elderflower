var MessageRow = React.createClass({
  numClicked: function(event) {
    Dispatcher.fireEvent("receipient_selected", {num: this.props.Num})
  },
  render: function() {
    return <div className={this.props.You ? "message you" : "message"}>
      <p>
      {this.props.Message}
      </p>
    </div>
  }
});
