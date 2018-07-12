const path = require("path")

function resolve(dir) {
  return path.resolve(__dirname, "..", dir)
}

const fs = require("fs")
const zlib = require("zlib")

class CompressPlugin {
  constructor(option = {}) {
    this.option = option
  }

  gzip(file) {
    fs.readFile(file, (err, data) => {
      if (err) { throw err }
      zlib.gzip(data, (err, buf) => {
        if (err) { throw err }
        fs.writeFile(file, buf, (err) => {
          if (err) { throw err }
        })
      })
    })
  }

  passTest(file) {
    return this.option.test.test(file)
  }

  walk(dir) {
    fs.stat(dir, (err, stats) => {
      if (err) { throw err }
      if (stats.isFile()) {
        if (this.passTest(dir)) {
          this.gzip(dir)
        }
        return
      }

      fs.readdir(dir, (err, files) => {
        if (err) { throw err }
        for (let f of files) {
          this.walk(path.resolve(dir, f))
        }
      })
    })
  }

  apply(compiler) {
    compiler.hooks.done.tapAsync("CompressPlugin", (compilation, cb) => {
      const root = compilation.compilation.outputOptions.path
      this.walk(root)
      cb()
    })
  }
}

module.exports = {
  resolve,
  CompressPlugin,
}
