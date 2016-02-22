define(function(require) {
    var Course = require('app/models/course');
    var MixinLayout = require('app/views/layout')
    var View = require('app/views/course/view');

    return {
        controller: function() {
            var courseID = m.route.param("courseID");
            this.course = m.prop(new Course());
            var ctrl = this;
            Course.fetch(courseID, function(c) {
                ctrl.course = m.prop(c);
            });
        },
        view: function(ctrl) {
            return MixinLayout(View, ctrl.course());
        }
    };

});
