module.exports = {
  publicPath: "/",
  outputDir: "dist",
  productionSourceMap: false,
  devServer: {
    port: 8080,
    proxy: {
      '/api': {
        target: 'https://blog.renj.io',
        // target: 'http://127.0.0.1:5000',
        changeOrigin: true,
        ws: false,
        pathRewrite: {
          '^/api': '/api'
        },
      },
    },
  },
  configureWebpack: config =>{
    // config.entry.app = ["@babel/polyfill","./src/main.js"];
    config.externals = {
      'vue': 'Vue',
      'element-ui': 'ELEMENT',
      'axios': 'axios',
      'vue-router': 'VueRouter',
      'marked': 'marked',
    }
  },
};