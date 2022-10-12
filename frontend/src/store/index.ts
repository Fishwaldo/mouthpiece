import { createStore } from 'vuex'
import createCache from 'vuex-cache'
import { AppsService, MessagesService, Put_App_request, Post_App_request } from '@/generated/'

export default createStore({
  plugins: [createCache({ timeout: 30000 })],
  state: {
    sidebarVisible: false,
    sidebarUnfoldable: false,
  },
  mutations: {
    toggleSidebar(state) {
      state.sidebarVisible = !state.sidebarVisible
    },
    toggleUnfoldable(state) {
      state.sidebarUnfoldable = !state.sidebarUnfoldable
    },
    updateSidebarVisible(state, payload) {
      state.sidebarVisible = payload.value
    },
  },
  actions: {
    'FETCH_APP': async(_, id) => {
      const response = await AppsService.getApp({id: id})
      return response
    },
    'PATCH_APP': async(_, request) => {
      const payload = {id: request.id, requestBody: request.payload}
      const response = await AppsService.patchApp(payload)
      return response
    },
    'PUT_APP': async(_, request) => {
      const payload: {
        id: number,
        requestBody: Put_App_request
      } = {
        id: request.id, 
        requestBody: {
          name: request.payload.name,
          status: request.payload.status,
          description: request.payload.description,
          Icon: request.payload.Icon,
          URL: request.payload.URL,
        }
      }
      const response = await AppsService.putApp(payload)
      return response
    },
    'POST_APP': async(_, request) => {
      const payload: {
        requestBody: Post_App_request
      } = {
        requestBody: {
          name: request.payload.name,
          status: request.payload.status,
          description: request.payload.description,
          Icon: request.payload.Icon,
          URL: request.payload.URL,
        }
      }
      const response = await AppsService.postApp(payload)
      return response
    },
    'FETCH_MSG': async(_: any, id: string) => {
      const response = await MessagesService.getMessage({msgid: id})
      return response
    }

  },
  modules: {
  },
})
