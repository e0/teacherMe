define(function(require) {
  var Course = require('app/models/course');
  var Discussion = require('app/models/discussion');
  var Download = require('app/models/download');
  var Assignment = require('app/models/assignment');
  var MixinLayout = require('app/views/layout')
  var CourseBody = require('app/views/course');

  var CoursePage = {
    controller: function() {
      var courseID = m.route.param("courseID");
      var ctrl = this;
      this.course = m.prop(new Course());

      m.request({
        method: "GET",
        url: "api/course/" + courseID
      }).then(function(courseData) {
        var course = new Course(courseData);

        for (i = 0; i < course.discussions().length; i++) {
          var discussion = new Discussion(course.discussions()[i]);
          course.discussions()[i] = discussion;
        }

        for (i = 0; i < course.downloads().length; i++) {
          var download = new Download(course.downloads()[i]);
          course.downloads()[i] = download;
        }

        for (i = 0; i < course.assignments().length; i++) {
          var assignment = new Assignment(course.assignments()[i]);
          course.assignments()[i] = assignment;
        }

        ctrl.course = m.prop(course);
      });
    },
    view: function(ctrl) {
      return MixinLayout(CourseBody, ctrl.course());
    }
  };

  return CoursePage;
});