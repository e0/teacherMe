define(function(require) {
  var m = require('lib/mithril.min');
  var CourseAllCtrl = require('app/components/course/all');
  var CourseViewCtrl = require('app/components/course/view');
  var CourseNewCtrl = require('app/components/course/new');

  m.route(document, "/", {
    "/": CourseAllCtrl,
    "/course/new": CourseNewCtrl,
    "/course/:courseID": CourseViewCtrl
  });
});
