const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  publicPath: '/static',
  devServer: {
    proxy: {
      '^/v1': {
        target: 'http://localhost:8888',
        changeOrigin: true,
      },
      '^/auth': {
        target: 'http://localhost:8888',
        changeOrigin: true,
      },
      '^/avatar': {
        target: 'http://localhost:8888',
        changeOrigin: true,
      },
      '^/config': {
        target: 'http://localhost:8888',
        changeOrigin: true,
      },
      '^/health': {
        target: 'http://localhost:8888',
        changeOrigin: true,
      },
    },
  }
})
