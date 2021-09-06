module.exports = {
  publicPath: './',
  devServer: {
    port: process.env.VUE_APP_CLI_PORT,
    proxy: {
      [process.env.VUE_APP_BASE_API]: { // 需要代理的路径   例如 '/api'
        target: `${process.env.VUE_APP_BASE_PATH}:${process.env.VUE_APP_SERVER_PORT}/`, // 代理到 目标路径
        changeOrigin: true,
        pathRewrite: { // 修改路径数据
          ['^' + process.env.VUE_APP_BASE_API]: '' 
        }
      }
    }
  }
}