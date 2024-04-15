'use strict'
const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
    entry: './static/src/js/main.js',
    output: {
        filename: 'js/main.js',
        path: path.resolve(__dirname, 'static/dist'), // 输出目录路径
        publicPath: '/dist/', // 公共路径，用于在 HTML 中正确引用静态资源
        // 设置字体文件的输出路径
        assetModuleFilename: 'fonts/[name][ext]',
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
            filename: "css/[name].css",
            chunkFilename: "css/[id].css"
        }),
        new CopyWebpackPlugin({
            patterns: [
                { from: 'node_modules/highlight.js/styles/default.css', to: 'css/light.css' },
                { from: 'node_modules/highlight.js/styles/dark.css', to: 'css/dark.css' },
            ],
        }),
    ],
};