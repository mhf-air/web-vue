const path = require("path")
const HtmlWebpackPlugin = require("html-webpack-plugin")
const VueLoaderPlugin = require("vue-loader/lib/plugin")

const globalNodeModules = path.resolve(__dirname, "../node_modules")

module.exports = {
  entry: {
    app: ["babel-polyfill", "./src/app.js"],
  },

  output: {
    path: path.resolve(__dirname, "../www/js"),
    filename: "[name].js?[chunkhash]",
  },

  module: {
    rules: [ //
      {
        test: /\.js?$/,
        loader: "babel-loader",
        exclude: (file) => {
          return file.startsWith(globalNodeModules)
        },
      },
      {
        test: /\.vue$/,
        loader: "vue-loader",
      },
      {
        test: /\.pug$/,
        oneOf: [ //
          // this applies to `<template lang="pug">` in Vue components
          {
            resourceQuery: /^\?vue/,
            use: ["pug-plain-loader"],
          },
          // this applies to stand-alone .pug files
          {
            use: ["pug-loader"],
          },
        ],
      },
      {
        test: /\.styl(us)?$/,
        use: ["vue-style-loader", "css-loader", "stylus-loader"],
      }
    ],
  },

  plugins: [
    new HtmlWebpackPlugin({
      filename: "../index.html",
      template: "src/static/index.pug",
      inject: false,
    }),

    new VueLoaderPlugin(),
  ],

  optimization: {
    splitChunks: {
      chunks: "all",
      name: "vendor",
    },
  },

  node: {
    // prevent webpack from injecting useless setImmediate polyfill because Vue
    // source contains it (although only uses it if it's native).
    setImmediate: false,
    // prevent webpack from injecting mocks to Node native modules
    // that does not make sense for the client
    dgram: 'empty',
    fs: 'empty',
    net: 'empty',
    tls: 'empty',
    child_process: 'empty'
  }
}
