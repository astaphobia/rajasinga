const path = require('path'),
    webpack = require('webpack'),
    htmlWebpackPlugin = require('html-webpack-plugin'),
    miniCssExtractPlugin = require('mini-css-extract-plugin'),
    cleanWebpackPlugin = require('clean-webpack-plugin')

const isDev = process.env.NODE_MODE !== 'production'

let entries = [path.resolve(__dirname, 'src/index.tsx')]
if (isDev) entries.push('webpack-hot-middleware/client?path=/__webpack_hmr&timeout=20000&overlay-false')

let plugins = [
    new htmlWebpackPlugin({
        template: path.resolve(__dirname, 'src/index.html'),
        favicon: path.resolve(__dirname, 'src/favicon.ico'),
    }),
    new miniCssExtractPlugin({
        filename: '[name].css',
        chunkFilename: '[id].css',
    }),
    new webpack.NoEmitOnErrorsPlugin(),
    new cleanWebpackPlugin(),
    new webpack.optimize.OccurrenceOrderPlugin(),
]
if (isDev) plugins.push(new webpack.HotModuleReplacementPlugin())

module.exports = {
    mode: isDev ? 'development' : 'production',
    entry: entries,
    output: {
        filename: '[name].bundle.js',
        path: path.resolve(__dirname, 'dist'),
        publicPath: '/',
    },
    target: 'web',
    devtool: 'inline-source-map',
    resolve: {
        extensions: ['.js', '.tsx', '.ts', '.jsx'],
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
                loader: 'html-loader'
            },
            {
                test: /\.css$/,
                use: [{ loader: miniCssExtractPlugin.loader }, 'css-loader'],
            },
        ],
    },
    devServer: {
        contentBase: path.resolve(__dirname, 'dist'),
        publicPath: '/',
        port: process.env.PORT || 5001,
        host: '0.0.0.0',
        hot: true,
        watchOptions: {
            aggregateTimeout: 300,
            poll: true,
        },
    },
    stats: 'errors-only',
    plugins: plugins,
}
