const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  publicPath: '/static',
  devServer: {
    allowedHosts: "all",
    port: 8888,
    proxy: {
      '^/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  }
})
