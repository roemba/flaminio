const CompressionPlugin = require("compression-webpack-plugin");
const VueLoaderOptionsPlugin = require("vue-loader-options-plugin");
const webpack = require("webpack");

module.exports = {
	output: {
		filename: "build.js"
	},
	resolve: {
		extensions: [".js", ".vue", ".json"]
	},
	module: {
		rules: [
			{
				test: /\.js$/,
				exclude: /(node_modules)/,
				use: {
					loader: "babel-loader",
					options: {
						presets: ["env"]
					}
				}
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
			}
		]
	},
	plugins: [
		new VueLoaderOptionsPlugin({
			"sass-resources-loader": {
				resources: ["./src/webroot/variables.scss"]
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
