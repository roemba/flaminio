const CompressionPlugin = require("compression-webpack-plugin");
const VueLoaderOptionsPlugin = require("vue-loader-options-plugin");
const webpack = require("webpack");
const path = require("path");

module.exports = {
	output: {
		filename: "build.js",
	},
	resolve: {
		extensions: [".js", ".vue", ".json"],
		alias: {
			"@": path.resolve("src/flaminio/webroot")
		}
	},
	module: {
		rules: [
			{
				test: /\.js$/,
				exclude: /(node_modules)/,
				loader: "babel-loader"
			},
			{
				test: /\.vue$/,
				loader: "vue-loader",
				options: {
					esModule: true,
					preLoaders: {
						scss: "sass-resources-loader"
					}
				}
			},
			{
				test: /\.css$/,
				loader: ["style-loader", "css-loader"]
			},
			{
				test: /\.svg$/,
				loader: "file-loader",
				options: {
					outputPath: "/images/"
				}
			}
		]
	},
	plugins: [
		new VueLoaderOptionsPlugin({
			"sass-resources-loader": {
				resources: ["./src/flaminio/webroot/scss/variables.scss"]
			}
		}),
		new webpack.ContextReplacementPlugin(/moment[\/\\]locale$/, /nl|en-gb|en-us/),
		new webpack.ProvidePlugin({
			$: "jquery",
			jQuery: "jquery",
			"window.jQuery": "jquery",
			Popper: ["popper.js", "default"]
		}),
		new CompressionPlugin({
			asset: "[path].gz[query]",
			algorithm: "gzip",
			test: /\.js$|\.css$|\.html$/,
			threshold: 10240,
			minRatio: 0.8
		})
	],
	devtool: "source-map"
};
