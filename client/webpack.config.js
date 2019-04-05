const path = require('path'),
  webpack = require('webpack'),
  htmlWebpackPlugin = require('html-webpack-plugin'),
  miniCssExtractPlugin = require('mini-css-extract-plugin'),
  isDev = process.env.NODE_MODE !== 'production'

let entries = [path.resolve(__dirname, 'src/index.tsx')]
if (isDev)
  entries.push(
    'webpack-hot-middleware/client?path=/__webpack_hmr&timeout=20000&overlay-false'
  )

let plugins = [
  new webpack.ProgressPlugin(),
  new htmlWebpackPlugin({
    template: path.resolve(__dirname, 'src/index.html'),
    favicon: path.resolve(__dirname, 'src/favicon.ico'),
  }),
  new miniCssExtractPlugin({
    filename: '[name].css',
    chunkFilename: '[id].css',
  }),
  new webpack.NoEmitOnErrorsPlugin(),
  new webpack.optimize.OccurrenceOrderPlugin(),
]
if (isDev) plugins.push(new webpack.HotModuleReplacementPlugin())

module.exports = {
  mode: isDev ? 'development' : 'production',
  entry: entries,
  output: {
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/',
    filename: '[name].bundle.js',
  },
  target: 'web',
  devtool: isDev ? 'inline-source-map' : '',
  resolve: {
    extensions: ['.js', '.tsx', '.ts', '.jsx'],
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
  module: {
    rules: [
      {
        test: /\.tsx$/,
        loader: 'ts-loader',
        options: {
          configFile: 'tsConfig.json',
        },
        exclude: /node_modules/,
      },
      {
        test: /\.html$/,
        loader: 'html-loader',
      },
      {
        test: /\.(scss|sass|css)$/,
        use: [
          isDev ? 'style-loader' : miniCssExtractPlugin.loader,
          'css-loader',
          'sass-loader',
        ],
      },
    ],
  },
  plugins: plugins,
  devServer: {
    contentBase: path.resolve(__dirname, 'dist'),
    publicPath: '/',
    historyApiFallback: true,
    port: process.env.PORT || 5001,
    host: '0.0.0.0',
    hot: true,
    watchOptions: {
      poll: true,
      aggregateTimeout: 500,
      stats: 'errors-only',
    },
  },
  stats: 'errors-only',
}
