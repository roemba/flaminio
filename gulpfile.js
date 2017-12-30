const gulp = require("gulp");
const del = require("del");
const webpack = require("webpack");
const webpackStream = require("webpack-stream");
const runSequence = require("run-sequence");
const eslint = require("gulp-eslint");

const WEBROOT = "src/flaminio/webroot";
const JAVASCRIPT_AND_VUE_SOURCES = `${WEBROOT}/**/*.{js,vue}`;

const DIST_FOLDER = "public";

let webpackConfig = require("./webpack.config.js");

gulp.task("clean", () => {
	return del([`${DIST_FOLDER}/**`, `!${DIST_FOLDER}`, `!${DIST_FOLDER}/{.htaccess,index.php}`]);
});

gulp.task("copy-static-files", () => {
	return gulp.src([`${WEBROOT}/**/*.html`, `${WEBROOT}/images/favicons/*`])
		.pipe(gulp.dest(`${DIST_FOLDER}`));
});

gulp.task("lint", () => {
	gulp.src(JAVASCRIPT_AND_VUE_SOURCES)
		.pipe(eslint())
		.pipe(eslint.format("codeframe"))
		.pipe(eslint.failAfterError());
});

gulp.task("webpack-configure", () => {
	webpackConfig.plugins.unshift(
		new webpack.DefinePlugin({
			"process.env": {
				"NODE_ENV": JSON.stringify("production")
			}}),
		new webpack.optimize.UglifyJsPlugin({
			sourceMap: true
		}),
		new webpack.optimize.AggressiveMergingPlugin()
	);
});

gulp.task("webpack", () => {
	return gulp.src(`${WEBROOT}/index.js`)
		.pipe(webpackStream(webpackConfig))
		.pipe(gulp.dest(`${DIST_FOLDER}`));
});

gulp.task("build", () => {
	runSequence("clean",
		["lint", "copy-static-files", "webpack"]);
});

gulp.task("build-prod", () => {
	runSequence(["clean", "webpack-configure"],
		["lint", "copy-static-files", "webpack"]);
});

gulp.task("watch", ["build"], () => {
	gulp.watch(`${WEBROOT}/**/*.*`, ["build"]);
});
