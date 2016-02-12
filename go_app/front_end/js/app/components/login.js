define(function(require) {
  var MixinLayout = require('app/views/layout')
  var User = require('app/models/user')
  var View = require('app/views/login');

  return {
    controller: function() {
      this.user = new User();
      var ctrl = this;

      this.login = function(e) {
        e.preventDefault();

        User.login(ctrl.user);
      };

      this.signup = function(e) {
        e.preventDefault();

        User.signup(ctrl.user);
      }
    },
    view: function(ctrl) {
      return MixinLayout(View, ctrl);
    }
  };

});
