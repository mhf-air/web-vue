const { resolve } = require("./util.js")
const common = require("./webpack.common.js")

const merge = require("webpack-merge")

const MiniCssExtractPlugin = require("mini-css-extract-plugin")

module.exports = merge(common, {
  mode: "development",
  stats: "minimal",

  output: {
    // publicPath: "/",
    path: resolve("www"),
    filename: "js/[name].js",
  },

  module: {
    rules: [ //
      {
        test: /\.styl(us)?$/,
        use: ["vue-style-loader", "css-loader", "stylus-loader"],
      },
      {
        test: /\.css$/,
        use: [ //
          {
            loader: MiniCssExtractPlugin.loader,
            options: {},
          },
          "css-loader",
        ],
      },
    ],
  },

  plugins: [
    new MiniCssExtractPlugin({
      filename: "css/[name].css",
    }),
  ],

  devServer: {
    contentBase: resolve("www"),
    historyApiFallback: true,
    publicPath: "/js/",
    overlay: true,
    open: true,
    noInfo: true,
    historyApiFallback: {
      rewrites: [{
          from: /\.js$/,
          to(context) {
            let p = context.parsedUrl.path
            let i = p.indexOf("/js/")
            return p.substring(i + 1)
          },
        },
        {
          from: /\.css$/,
          to(context) {
            let p = context.parsedUrl.path
            let i = p.indexOf("/css/")
            return p.substring(i + 1)
          },
        },
        {
          from: /\/img\//,
          to(context) {
            let p = context.parsedUrl.path
            let i = p.indexOf("/img/")
            return p.substring(i + 1)
          },
        },
        // { from: /./, to: "index.html" },
      ],
    },
  },
})
