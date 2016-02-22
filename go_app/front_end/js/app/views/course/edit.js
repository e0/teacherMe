define(function() {
  var CourseBody = function(ctrl) {
    return m("form.row", { onsubmit: ctrl.updateCourse }, m("div.row", [
      m("div", [
        m("label", "Title"),
        m("input.u-full-width", {onchange: m.withAttr("value", ctrl.course.title), value: ctrl.course.title()})
      ]),
      m("div", [
        m("label", "Description"),
        m("textarea.u-full-width", {onchange: m.withAttr("value", ctrl.course.description), value: ctrl.course.description()})
      ]),
      m("p", [m("button.button-primary[type=submit]", "Update course")])
    ]));
  };

  return CourseBody;
});
