var gulp   = require('gulp')
  , concat = require('gulp-concat')

gulp.task('default', function() {
  gulp.src('./views/**')
      .pipe(gulp.dest('compiled/'));
  console.log("here");
  gulp.src('./app/react/**/*.jsx')
      .pipe(concat('app.jsx'))
      .pipe(gulp.dest('compiled/jsx'));
});

gulp.task('live',['default'], function() {
  gulp.watch('./views/**', [ 'default' ]);
});
