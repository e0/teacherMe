define(function() {
  var CourseBody = function(ctrl) {
    return m("form", { onsubmit: ctrl.createCourse }, [
      m("p", [
        m("label", "Title"),
        m("input", {onchange: m.withAttr("value", ctrl.course.title), value: ctrl.course.title()})
      ]),
      m("p", [
        m("label", "Description"),
        m("input", {onchange: m.withAttr("value", ctrl.course.description), value: ctrl.course.description()})
      ]),
      m("p", [m("button[type=submit]", "Create course")])
    ]);
  };

  return CourseBody;
});
