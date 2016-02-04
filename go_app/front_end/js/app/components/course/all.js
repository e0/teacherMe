define(function(require) {
  var Course = require('app/models/course');
  var MixinLayout = require('app/views/layout')
  var View = require('app/views/course/all');

  return {
    controller: function() {
      this.courses = m.prop([]);
      var ctrl = this;
      Course.fetchAll(function(cs) {
        ctrl.courses = m.prop(cs);
      });
    },
    view: function(ctrl) {
      return MixinLayout(View, ctrl.courses());
    }
  };

});