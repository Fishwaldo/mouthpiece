const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  publicPath: '/static',
  devServer: {
    allowedHosts: "all",
    port: 8888,
    proxy: {
      '^/v1': {
        target: 'http://localhost:8889',
        changeOrigin: true,
      },
      '^/auth': {
        target: 'http://localhost:8889',
        changeOrigin: true,
      },
      '^/avatar': {
        target: 'http://localhost:8889',
        changeOrigin: true
      },
      '^/config': {
        target: 'http://localhost:8889',
        changeOrigin: true,
      },
      '^/health': {
        target: 'http://localhost:8889',
        changeOrigin: true,
      },
    },
  }
})
