define(function(require) {
  var Course = require('app/models/course');
  var MixinLayout = require('app/views/layout')
  var CourseBody = require('app/views/course/new');

  return {
    controller: function() {
      this.course = new Course();
      var ctrl = this;
      this.createCourse = function(e) {
        e.preventDefault();

        Course.create(ctrl.course, function(courseID) {
          m.route("/course/" + courseID);
        });
      };
    },
    view: function(ctrl) {
      return MixinLayout(CourseBody, ctrl);
    }
  };
});
