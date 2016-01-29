define(function(require) {
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

  return Course;
});