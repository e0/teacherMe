define(function(require) {
    var Course = require('app/models/course');
    var MixinLayout = require('app/views/layout')
    var CourseBody = require('app/views/course/edit');

    return {
        controller: function() {
            if (localStorage.getItem("id_token") == null) m.route("/courses");

            var courseID = m.route.param("courseID");
            this.course = m.prop(new Course());
            var ctrl = this;

            Course.fetch(courseID, function(c) {
                ctrl.course = c;
            });

            this.updateCourse = function(e) {
                e.preventDefault();

                Course.update(ctrl.course, function(courseID) {
                    m.route("/course/" + courseID);
                });
            };
        },
        view: function(ctrl) {
            return MixinLayout(CourseBody, ctrl);
        }
    };
});
