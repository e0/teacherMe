define(function() {
  var CoursesBody = function(courses) {
    return [
      courses.map(function(course) {
        return m("p", [m("a[href='?/course/" + course.id() + "']", course.title())])
      })
    ];
  };

  return CoursesBody;
});