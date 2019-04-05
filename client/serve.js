const path = require('path')
const express = require('express')
const webpack = require('webpack')
const dev = require('webpack-dev-middleware')
const hot = require('webpack-hot-middleware')

const config = require('./webpack.config.js')
const port = config.devServer.port
const app = express()
const compiler = webpack(config)

console.log('ENV_MODE: ' + config.mode)
if (config.mode === 'development') {
  app.use(
    dev(compiler, {
      publicPath: config.output.publicPath,
      noInfo: true,
      stats: 'errors-only',
      watchOptions: {
        poll: true,
        aggregateTimeout: 300,
      },
    })
  )
  app.use(
    hot(compiler, {
      log: console.log,
      reload: true,
      path: '/__webpack_hmr',
      heartbeat: 2000,
    })
  )
}

app.use(express.static(path.join(__dirname, 'dist')))
app.get('*', (_req, res) =>
  res.sendFile(path.resolve(__dirname, 'dist/index.html'))
)
app.listen(port, '0.0.0.0', function onStart(err) {
  if (err) {
    console.log(err)
  }
  console.info(
    '==> 🌎  Listening on port %s. Open up http://0.0.0.0:%s/ in your browser.',
    port,
    port
  )
})
