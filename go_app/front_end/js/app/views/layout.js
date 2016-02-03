define(function() {
  var layout = function(head, nav, body) {
    return m("html", [
      m("head", head),
      m("body", [
        m("header", nav),
        m("div.container", body)
      ]) 
    ]);
  };

  var head = function() {
    return [
      m("meta[charset='utf-8'"),
      m("title", "Making your first web page"),
      m("link[rel='stylesheet'][type='text/css'][href='http://fonts.googleapis.com/css?family=Raleway:400,300,600']"),
      m("link[rel='stylesheet'][type='text/css'][href='css/normalize.css']"),
      m("link[rel='stylesheet'][type='text/css'][href='css/skeleton.css']"),
      m("link[rel='stylesheet'][type='text/css'][href='css/global.css']"),
      m("link[rel='stylesheet'][type='text/css'][href='css/course.css']")
    ];
  };

  var nav = function() {
    return m("div.container", [
      m("span#logo.two.columns", "Logo"),
      m("nav.ten.columns", [
        m("ul", [
          m("li", [m("a[href='#']", "Courses")]),
          m("li", [m("a[href='#']", "Help")]),
          m("li", [m("a.button.button-green[href='#']", "Log in")])
        ])
      ])
    ]);
  };

  var MixinLayout = function(body, bodyData) {
    return layout(head(), nav(), body(bodyData));
  };

  return MixinLayout;
});