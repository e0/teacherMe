define(function() {
  var LoginBody = function(ctrl) {
    return m("div.row", [
      m("div.one-half.column", loginHalf(ctrl)),
      m("div.one-half.column", signupHalf(ctrl))
    ]);
  };

  var loginHalf = function(ctrl) {
    return m("form", { onsubmit: ctrl.login }, m("div.row", [
      m("div", [
        m("h3", "Log in"),
        m("label", "Email"),
        m("input.u-full-width[type='text']", { onchange: m.withAttr("value", ctrl.user.username) })
      ]),
      m("div", [
        m("label", "Password"),
        m("input.u-full-width[type='password']", { onchange: m.withAttr("value", ctrl.user.password) })
      ]),
      m("p", [m("button.button-green[type=submit]", "Log in")])
    ]))
  };

  var signupHalf = function(ctrl) {
    return m("form", { onsubmit: ctrl.signup }, m("div.row", [
      m("div", [
        m("h3", "Sign up"),
        m("label", "Email"),
        m("input.u-full-width[type='email']", { onchange: m.withAttr("value", ctrl.user.email) })
      ]),
      m("div", [
        m("label", "Create password"),
        m("input.u-full-width[type='password']", { onchange: m.withAttr("value", ctrl.user.password) })
      ]),
      m("p", [m("button.button-primary[type=submit]", "Sign up")])
    ]))
  };

  return LoginBody;
});
