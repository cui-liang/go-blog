'use strict'
const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
    entry: './static/js/main.js',
    output: {
        filename: 'main.js',
        path: path.resolve(__dirname, 'static/dist'), // 输出目录路径
        publicPath: '/dist/', // 公共路径，用于在 HTML 中正确引用静态资源
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: 'babel-loader',
            },
            {
                test: /\.css$/,
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader'
                ],
            },
            {
                test: /\.scss$/,
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader',
                    'sass-loader',
                ],
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif)$/i,
                type: 'asset/resource',
            },
            {
                test: /\.(woff|woff2|eot|ttf|otf)$/i,
                type: 'asset/resource',
            },
        ],
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: "[name].css",
            chunkFilename: "[id].css"
        }),
        new CopyWebpackPlugin({
            patterns: [
                { from: 'node_modules/highlight.js/styles/default.css', to: 'light.css' },
                { from: 'node_modules/highlight.js/styles/dark.css', to: 'dark.css' },
            ],
        }),
    ],
};