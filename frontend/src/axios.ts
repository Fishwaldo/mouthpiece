// import axios from 'axios'

// const transformData = (data: any, func: any) => {
// //   if (isArray(data)) {
// //     return data.map((value: any) => transformData(value, func))
// //   } else if (isPlainObject(data)) {
// //     const result = {}

// //     Object.keys(data).forEach(key => {
// //       const value = data[key]

// //       result[func(key)] = isObject(value) ? transformData(value, func) : value
// //     })

// //     return result
// //   } else {
// //     return data
// //   }
// }

// //axios.defaults.baseURL = 'http://localhost:8000'

// axios.interceptors.request.use(config => {
//     console.log("request")
//     console.log(config)

//     //   if (config.data) {
// //     config.data = transformData(config.data, snakeCase)
// //   }

//   return config
// })

// axios.interceptors.response.use(response => {
//     console.log("response")
//     console.log(response)
// //   if (response.data) {
// //     response.data = transformData(response.data, camelCase)
// //   }

//   return response
// })

// export { axios }