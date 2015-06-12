var gulp   = require('gulp')
  , concat = require('gulp-concat')

gulp.task('default', function() {
  gulp.src('./views/**')
      .pipe(gulp.dest('compiled/'));
  gulp.src(['./app/react/**/*.jsx', './app/js/**/*.js'])
      .pipe(concat('app.jsx'))
      .pipe(gulp.dest('compiled/jsx'));
  console.log("concat app css into compiled/css");
  gulp.src(['./app/styles/**/*.css'])
      .pipe(concat('app.css'))
      .pipe(gulp.dest('compiled/css'));
  console.log("done");
});

gulp.task('live',['default'], function() {
  gulp.watch('./views/**', [ 'default' ]);
});
