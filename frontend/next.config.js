/* eslint-disable no-console */
const withPlugins = require('next-compose-plugins');
const withCSS = require('@zeit/next-css');
const withSass = require('@zeit/next-sass');
const withImages = require('next-images');
const withBundleAnalyzer = require('@zeit/next-bundle-analyzer');

const env = process.env.CONFIG_ENV || process.env.NODE_ENV || 'development';
const envConfig = require('./env')(env);

console.debug('CONFIG_ENV=', process.env.CONFIG_ENV);
console.debug('NODE_ENV=', process.env.NODE_ENV);
console.debug('envConfig=', envConfig);

const nextConfig = {
  env: envConfig,
  serverRuntimeConfig: envConfig,
  analyzeServer: ['server', 'both'].includes(process.env.BUNDLE_ANALYZE),
  analyzeBrowser: ['browser', 'both'].includes(process.env.BUNDLE_ANALYZE),
  bundleAnalyzerConfig: {
    server: {
      analyzerMode: 'static',
      reportFilename: '../bundles/server.html',
    },
    browser: {
      analyzerMode: 'static',
      reportFilename: '../bundles/client.html',
    },
  },
};

module.exports = withPlugins([
  withCSS, withSass, withBundleAnalyzer, withImages,
], nextConfig);
