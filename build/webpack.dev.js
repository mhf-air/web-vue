const merge = require("webpack-merge")
const common = require("./webpack.common.js")
const path = require("path")

module.exports = merge(common, {
  mode: "development",
  stats: "minimal",
  devServer: {
    contentBase: path.join(__dirname, "../www"),
    historyApiFallback: true,
    publicPath: "/js/",
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
