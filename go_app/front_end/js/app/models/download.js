define(function() {
  var Download = function(data) {
    data = data || {};
    this.id = m.prop(data.id || "");
    this.name = m.prop(data.name || "");
    this.date = m.prop(data.date || "");
  };

  return Download;
});