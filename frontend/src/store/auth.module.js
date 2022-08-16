import SSOProviders from '@/models/ssoproviders'
import AuthService from '../services/auth.service'
const user = JSON.parse(localStorage.getItem('user'))
const initialState = user
  ? { status: { loggedIn: true }, user }
  : { status: { loggedIn: false }, user: null }
export const auth = {
  namespaced: true,
  state: initialState,
  actions: {
    login({ commit }, user) {
      return AuthService.login(user).then(
        (user) => {
          commit('loginSuccess', user)
          return Promise.resolve(user)
        },
        (error) => {
          commit('loginFailure')
          return Promise.reject(error)
        },
      )
    },
    logout({ commit }) {
      AuthService.logout()
      commit('logout')
    },
    providers({ commit }) { 
      return AuthService.providers().then(
        (providers) => {
          commit('providers', providers)
          return Promise.resolve(providers)
        },
        (error) => {
          commit('providersFailure')
          return Promise.reject(error)
        }
      )
    },
    feconfig({ commit }) {
      return AuthService.feconfig().then(
        (config) => {
          commit('feconfig', config)
          return Promise.resolve(config)
        }
      )
    },
  },
  mutations: {
    loginSuccess(state, user) {
      state.status.loggedIn = true
      state.user = user
    },
    loginFailure(state) {
      state.status.loggedIn = false
      state.user = null
    },
    logout(state) {
      state.status.loggedIn = false
      state.user = null
    },
    providers(state, providers) {
      state.providers = providers
    },
    providersFailure(state) {
      state.providers = new SSOProviders()
    },
    feconfig(state, config) {
      state.config = config
    }
  },
}
