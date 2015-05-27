var gulp   = require('gulp')
  , concat = require('gulp-concat')

gulp.task('default', function() {
  gulp.src('./views/**')
      .pipe(gulp.dest('compiled/'));
  console.log("here");
  gulp.src(['./app/react/**/*.jsx', './app/js/**/*.js'])
      .pipe(concat('app.jsx'))
      .pipe(gulp.dest('compiled/jsx'));
  gulp.src(['./app/css/**/*.css'])
      .pipe(concat('app.css'))
      .pipe(gulp.dest('compiled/css'));

});

gulp.task('live',['default'], function() {
  gulp.watch('./views/**', [ 'default' ]);
});
