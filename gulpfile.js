var gulp = require('gulp');



gulp.task('default', function() {
  gulp.src('./views/**')
      .pipe(gulp.dest('compiled/'));
});

gulp.task('live',['default'], function() {
  gulp.watch('./views/**', [ 'default' ]);
});
