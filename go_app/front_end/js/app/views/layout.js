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
      m("meta[charset='utf-8']"),
      m("title", "teacherMe"),
      m("link[rel='stylesheet'][type='text/css'][href='http://fonts.googleapis.com/css?family=Raleway:400,300,600']"),
      m("link[rel='stylesheet'][type='text/css'][href='css/normalize.css']"),
      m("link[rel='stylesheet'][type='text/css'][href='css/skeleton.css']"),
      m("link[rel='stylesheet'][type='text/css'][href='css/global.css']"),
      m("link[rel='stylesheet'][type='text/css'][href='css/course.css']")
    ];
  };

  function logOut() {
    localStorage.removeItem('id_token');
    m.route("/");
  }

  var nav = function() {
    var logButton;
    
    if (localStorage.getItem('id_token') == null) {
      logButton = m("a.button.button-green[href='?/login']", "Log in");
    } else {
      logButton = m("a.button", { onclick: logOut }, "Log out");
    }

    return m("div.container", [
      m("span#logo.two.columns", "Logo"),
      m("nav.ten.columns", [
        m("ul", [
          m("li", [m("a[href='?/courses']", "Courses")]),
          m("li", [m("a[href='#']", "Help")]),
          m("li", [logButton])
        ])
      ])
    ]);
  };

  var MixinLayout = function(body, bodyData) {
    return layout(head(), nav(), body(bodyData));
  };

  return MixinLayout;
});
