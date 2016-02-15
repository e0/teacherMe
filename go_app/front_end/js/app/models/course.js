define(function(require) {
  var Discussion = require('app/models/discussion');
  var Download = require('app/models/download');
  var Assignment = require('app/models/assignment');

  var Course = function(data) {
    data = data || {};

    this.id = m.prop(data.id || "");
    this.title = m.prop(data.title || "");
    this.description = m.prop(data.description || "");
    this.discussions = m.prop(data.discussions || []);
    this.downloads = m.prop(data.downloads || []);
    this.assignments = m.prop(data.assignments || []);
    this.teacherName = m.prop(data.teacherName || "");
    this.studentsCount = m.prop(data.studentsCount || 0);
  };

  Course.create = function(course, callback) {
    var courseData = {
      title: course.title(),
      description: course.description()
    };

    m.request({
      method: "POST",
      url: "api/private/course_create",
      data: courseData,
      config: function(xhr) {
        xhr.setRequestHeader('Authorization', "Bearer " + localStorage.getItem('id_token'))
      }
    }).then(function(content) {
      callback(content.courseID);
    });
  };

  function parseCourse(courseData) {
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

    return course;
  }

  Course.fetch = function(courseID, callback) {
    m.request({
      method: "GET",
      url: "api/public/course/" + courseID
    }).then(function(courseData) {
      var course = parseCourse(courseData);
      callback(course);
    });
  };

  Course.fetchAll = function(callback) {
    m.request({
      method: "GET",
      url: "api/public/courses"
    }).then(function(coursesData) {
      var courses = coursesData.map(function(courseData) { return parseCourse(courseData); });
      callback(courses);
    });
  };

  return Course;
});
