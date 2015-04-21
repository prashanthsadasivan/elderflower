var QrCode = React.createClass({
  getInitialState: function() {
      return {qrsource: "/QR?hostname=" + document.URL};
  },
  componentDidMount: function() {
      var self = this;
  },

  render: function() {
      return <img src={this.state.qrsource} />
  }
});
