define(function(require) {
  var c = require('config');

  var User = function(data) {
    data = data || {};

    this.id = m.prop(data.id || "");
    this.username = m.prop(data.username || "");
    this.email = m.prop(data.email || "");
    this.password = m.prop(data.password || "");
  };

  User.login = function(user) {
    var userData = {
      client_id: c.AUTH0_CLIENT_ID,
      username: user.username(),
      password: user.password(),
      connection: "Username-Password-Authentication",
      grant_type: "password",
      scope: "openid"
    };

    m.request({
      method: "POST",
      url: c.AUTH0_DOMAIN + "oauth/ro",
      data: userData
    }).then(
      function(response) {
        localStorage.setItem('id_token', response.id_token);
        m.route("/");
      },
      function(error) {
        console.log(error)
      }
    );
  };

  User.signup = function(user) {
    var userData = {
      client_id: c.AUTH0_CLIENT_ID,
      email: user.email(),
      password: user.password(),
      connection: "Username-Password-Authentication"
    };

    m.request({
      method: "POST",
      url: c.AUTH0_DOMAIN + "dbconnections/signup",
      data: userData
    }).then(
      function(response) {
        console.log(response);
      },
      function(error) {
        console.log(error)
      }
    );
  }

  return User;
});
