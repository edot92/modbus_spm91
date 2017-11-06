import Vue from 'vue';
import Vuex from 'vuex';
import menu from '../menu';
import config from '../config';
Vue.use(Vuex);
// var baseUrl;
// if (process.env.NODE_ENV === 'development') {
//   baseUrl = 'http://localhost:9000/api';
//   // headers.common['Access-Control-Allow-Origin'] =
//   //   'http://127.0.0.1:9000';
//   // headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
// } else {
//   baseUrl = '/api';
// }
const store = new Vuex.Store({
  state: {
    readyLoad: false,
    loading: true,
    license: false,
    pageTitle: 'Home',
    menu: menu,
    user: {},
    token: null,
    message: {
      type: null,
      body: null
    },
    baseurl: '',
    theme: null,
    config: config,
    tempData: null,
    api: {
      user: {
        getuserpaginate: 'user/getuserpaginate',
        useraddnew: 'user/addnew',
        userdeletebyid: 'user/deletebyid',
        userupdatebyid: 'user/updatebyid',
        cekuniquebykolom: 'user/cekuniquebykolom'
      },
      gardu: {
        getgardupaginate: 'gardu/getgardupaginate',
        garduaddnew: 'gardu/addnew',
        gardudeletebyid: 'gardu/deletebyid',
        garduupdatebyid: 'gardu/updatebyid'
      }
    }
  },
  mutations: {
    setTempData (state, data) {
      state.tempData = data;
    },
    setReadyLoad (state, data) {
      state.readyLoad = data;
    },
    setBaseUrl (state, data) {
      state.baseurl = data;
    },
    setLicense (state, data) {
      state.license = data;
    },
    setAuth (state, { user, token }) {
      state.user = user;
      state.token = token;
      global.helper.ls.set('user', user);
      global.helper.ls.set('token', token);
    },
    setMenu (state, data) {
      state.menu = data;
    },
    setPageTitle (state, data) {
      state.pageTitle = data;
    },
    showMessage (state, type, body) {
      state.message = { type, body };
    },
    setLoading (state, data) {
      state.loading = data;
    }
  },
  actions: {
    checkAuth ({ commit }) {
      let data = {
        user: global.helper.ls.get('user'),
        token: global.helper.ls.get('token')
      };
      commit('setAuth', data);
    },
    checkPageTitle ({ commit, state }, path) {
      for (let k in state.menu) {
        if (state.menu[k].href === path) {
          commit('setPageTitle', state.menu[k].title);
          break;
        }
      }
    }
  }
});

export default store;
