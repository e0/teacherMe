define(function(require) {
    var m = require('lib/mithril.min');
    var CourseAll = require('app/components/course/all');
    var Login = require('app/components/login');
    var CourseView = require('app/components/course/view');
    var CourseNew = require('app/components/course/new');
    var CourseEdit = require('app/components/course/edit');

    m.route(document, "/", {
        "/": CourseAll,
        "/login": Login,
        "/course/new": CourseNew,
        "/course/:courseID": CourseView,
        "/course/edit/:courseID": CourseEdit
    });
});
