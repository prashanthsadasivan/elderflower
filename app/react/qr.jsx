var QrCode = React.createClass({
  getInitialState: function() {
      return {qrsource: "/QR?hostname=" + document.URL};
  },
  render: function() {
    console.log(this.props);
    return <img src={this.state.qrsource + '&qr_secret=' + this.props.qr_secret}/>;
  }
});
