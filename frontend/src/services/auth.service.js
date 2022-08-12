import axios from 'axios'
//import User from '../models/user'
import SSOProviders from '@/models/ssoproviders'
//const API_URL = 'http://10.11.5.10:8080/auth/'
class AuthService {
  login(user) {
    console.log(user)
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
        console.log(response.data)
        if (response.status == 200) {
          localStorage.setItem('user', JSON.stringify(response.data))
        }
        return response.data
      })
  }
  logout() {
    return axios.get('/auth/logout').then((response) => {
      localStorage.removeItem('user')
      return response.data
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
    return axios.get('/config/frontend')
    .then((response) => {
      return response.data
    })
  }
}
export default new AuthService()
