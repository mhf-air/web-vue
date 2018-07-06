const { resolve } = require("./util.js")

const HtmlWebpackPlugin = require("html-webpack-plugin")
const VueLoaderPlugin = require("vue-loader/lib/plugin")

const Root_Node_Modules = resolve("node_modules")

module.exports = {
  entry: {
    app: ["babel-polyfill", "./src/app.js"],
  },

  module: {
    rules: [ //
      {
        test: /\.js?$/,
        loader: "babel-loader",
        exclude: (file) => {
          return file.startsWith(Root_Node_Modules)
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
