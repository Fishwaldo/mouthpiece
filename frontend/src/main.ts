import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'


// @ts-ignore
import {createAuth}          from '@websanova/vue-auth/src/v3.js';
// // @ts-ignore
// import driverAuthBearer      from '@websanova/vue-auth/src/drivers/auth/bearer.js';
// @ts-ignore
import driverHttpAxios       from '@websanova/vue-auth/src/drivers/http/axios.1.x.js';
// @ts-ignore
import driverRouterVueRouter from '@websanova/vue-auth/src/drivers/router/vue-router.2.x.js';



import CoreuiVue from '@coreui/vue'
import CIcon from '@coreui/icons-vue'
import { iconsSet as icons } from '@/assets/icons'
import { OpenAPI, AppsService } from './generated'

OpenAPI.BASE = ''

// @ts-ignore
const auth = createAuth({
    plugins: {
      http: axios,
      router: router
    },
    drivers: {
      http: driverHttpAxios,
      auth: {
        request: function(req: any, token: string) {
          const refresh = req.url.indexOf('tokenrefresh') > -1
          const tokens = token.split(';')
          if (!refresh) {
            req.headers['Authorization'] = 'Bearer ' + tokens[0]
          } else {
            token = tokens[1]
            req.data = {refresh_token: token}
          }
        },
        response: function (res: any) {
            if ((res.config.url === '/api/auth/password' || res.config.url === '/api/auth/tokenrefresh') && res.data.Status === "OK") {
              return `${res.data.SessionToken};${res.data.RefreshToken}`
            }
        }
      },
      router: driverRouterVueRouter,
    },
    options: {
      authRedirect: {path: '/pages/login'},
      loginData: {
        url: '/api/auth/password',
        redirect: '/dashboard',
      },
      fetchData: {
        url: '/api/auth/me'
      },
      refreshData: {
        url: '/api/auth/tokenrefresh',
        method: 'POST',
        interval: 4
      },
      logoutData: {
        url: '/api/auth/logout',
        method: 'DELETE',
        redirect: '/pages/login'
      },
      stores: ['storage', 'cookie'],
      parseUserData(data : any) {
        const user = data;
        //store.dispatch('userModule/setUser', user)
        return user;
      },
    },
    auth: {
        response: function () {
            console.log("this")
        },
        request: function() {
            console.log("this request")
        }
    }
  });




const app = createApp(App)
app.use(store)
app.use(router)
app.use(CoreuiVue)
app.use(auth)
app.provide('icons', icons)
app.component('CIcon', CIcon)

app.mount('#app')
