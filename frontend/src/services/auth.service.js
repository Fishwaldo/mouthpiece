import axios from 'axios'
//import User from '../models/user'
import SSOProviders from '@/models/ssoproviders'
import { DefaultService } from '@/generated/'
import router from '@/router'
//const API_URL = 'http://10.11.5.10:8080/auth/'
class AuthService {
  login(user) {
    return axios
      .get('/auth/direct/login', {
        params: {
          user: user.username,
          passwd: user.password,
        },
      })
      .then((response) => {
        if (response.status == 403) {
          return Promise.reject(response.data)
        }
        if (response.status == 200) {
          localStorage.setItem('user', JSON.stringify(response.data))
          router.push({ path: '/dashboard'})
        }
        return response.data
      })
  }
  logout() {
    return axios.get('/auth/logout')
    .then((response) => {
      localStorage.removeItem('user')
      router.push({ path: '/pages/login'})
      return response.data
    })
    .catch(function (error) {
      console.log(JSON.stringify(error))
      localStorage.removeItem('user')
      console.log(this)
      router.push({ path: '/pages/login'})
    })
  }
  providers() {
    return axios.get('/auth/list')
    .then((response) => {
      console.log("get" + response.data)
      return new SSOProviders(response.data)
    })
  }
  feconfig() {
    return DefaultService.getConfig()
    .then((response) => {
      return response
    })
  }
  status() {
    return axios.get('/auth/status')
    .then((response) => {
      console.log("status" + response.data)
      return response
    })
  }
}
export default new AuthService()
