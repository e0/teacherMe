define(function() {
  var CourseBody = function(course) {
    return [
      m(".row.course-header", courseHeader(course)),
      m(".row.course-contents", courseContents(course))
    ];
  };

  var courseHeader = function(course) {
    return [
      m(".nine.columns", [m("h2", course.title())]),
      m(".three.columns", [m("a.button.button-primary[href='#']", "Join course")])
    ];
  };

  var courseContents = function(course) {
    return [
      m(".eight.columns.course-description", m.trust(course.description())),
      m(".four.columns.course-sidebar", courseSidebar(course))
    ];
  };

  var courseSidebar = function(course) {
    return [
      m("h6.teacher", [
        "Teacher: ",
        m("span", [m("a[href='#']", course.teacherName())])
      ]),
      m(".box", [
        m("h6", [
          "Discussions ",
          m("span.view-all-link", [m("a[href='#'", "view all")])
        ]),
        m("ul", [
          course.discussions().map(function(discussion) {
            return m("li", [
              m("a[href='#']", discussion.name()),
              " ",
              m("span.date", discussion.date())
            ])
          })
        ])
      ]),
      m(".box", [
        m("h6", [
          "Downloads ",
          m("span.view-all-link", [m("a[href='#'", "view all")])
        ]),
        m("ul", [
          course.downloads().map(function(download) {
            return m("li", [
              m("a[href='#']", download.name()),
              " ",
              m("span.date", download.date())
            ])
          })
        ])
      ]),
      m(".box", [
        m("h6", [
          "Assignments ",
          m("span.view-all-link", [m("a[href='#'", "view all")])
        ]),
        m("ul", [
          course.assignments().map(function(assignment) {
            return m("li", [
              m("a[href='#']", assignment.name()),
              " ",
              m("span.date", assignment.date())
            ])
          })
        ])
      ]),
      m("footer", "There are " + course.studentsCount() + " students taking this course")
    ];
  };

  return CourseBody;
});