const path = require('path'),
    webpack = require('webpack'),
    htmlWebpackPlugin = require('html-webpack-plugin'),
    miniCssExtractPlugin = require('mini-css-extract-plugin')

module.exports = {
    entry: path.resolve(__dirname, 'src/index.tsx'),
    output: {
        filename: '[name].bundle.js',
        path: path.resolve(__dirname, 'dist'),
    },
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
                test: /\.css$/,
                use: [{ loader: miniCssExtractPlugin.loader }, 'css-loader'],
            },
        ],
    },
    stats: {
        modules: false,
    },
    devServer: {
        contentBase: path.resolve(__dirname, 'dist'),
        port: process.env.PORT || 5001,
        host: '0.0.0.0',
        stats: {
            modules: false,
        },
        //disableHostCheck: true
    },
    plugins: [
        new htmlWebpackPlugin({
            template: path.resolve(__dirname, 'src/index.html'),
            favicon: path.resolve(__dirname, "src/favicon.ico"),
        }),
        new miniCssExtractPlugin({
            filename: '[name].css',
            chunkFilename: '[id].css',
        }),
        new webpack.HotModuleReplacementPlugin(),
    ],
}
