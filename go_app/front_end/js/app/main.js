define(function(require) {
  var m = require('lib/mithril');
  var CoursePage = require('app/components/course');


  m.route(document, "/course/1", {
    "/course/:courseID": CoursePage
  });
});