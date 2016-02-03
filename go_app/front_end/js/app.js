requirejs.config({
  baseUrl: 'js',
  path: {
    app: 'app'
  }
});

requirejs(['app/main']);