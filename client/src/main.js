import Vue from 'vue';
import helper from './helper';
global.helper = helper;
import config from './config';
import store from './store/';
global.store = store;
import router from './router';
import i18n from './i18n/';
import Vuetify from 'vuetify';
import './http';

import 'vuetify/src/stylus/main.styl';
import 'vuetify/src/stylus/settings/_colors.styl';
import '@/styles/main.styl';
import '@/styles/main.css';
import App from './App.vue';
require('sweetalert');

import VueTimeago from 'vue-timeago';
const Highcharts = require('highcharts');

// Load module after Highcharts is loaded
require('highcharts/modules/exporting')(Highcharts);
require('../node_modules/highcharts/highcharts-more')(Highcharts);

window.Highcharts = Highcharts;

require('./includelib.js');
Vue.use(Vuetify);
Vue.use(require('vue-moment'));
// Vue.use(swal);
Vue.use(VueTimeago, {
  name: 'timeago', // component name, `timeago` by default
  locale: config.locale,
  locales: {
    en: require('vue-timeago/locales/en-US.json'),
    [config.locale]: require(`vue-timeago/locales/${config.locale}.json`)
  }
});

import Dropzone from 'vue2-dropzone';
import VueQuillEditor from 'vue-quill-editor';
Vue.use(VueQuillEditor);
Vue.component('dropzone', Dropzone);

// import validator from 'indicative'
import validator from 'Validator';
global.validator = validator;

// import VForm from './components/Form.vue';
// import VGrid from './components/Grid.vue';
// import VField from './components/Field.vue';

// import Modal from './components/Modal'
// Vue.use(Modal)

// Vue.component('v-form', VForm);
// Vue.component('v-grid', VGrid);
// Vue.component('v-field', VField);

/* eslint-disable no-new */

new Vue({
  el: '#app',
  i18n,
  store,
  router,
  render: h => h(App),
  methods: {
    // back () {
    //   this.$router.go(-1);
    // }
  },
  beforeCreate () {
    //     "username":"admin",
    // "password":"admin"
    // let menu = [
    //   { header: 'System' },
    //   { href: '/laporanindex', title: 'Laporan', icon: 'home' }
    // ];
    // this.$store.commit('setMenu', menu);
    this.$store.commit('setLicense', true);
    this.$store.dispatch('checkAuth');
    // this.$store.commit('setBaseUrl', config.url);
    // window
    //   .axios({ method: 'get', url: baseUrl + '/info/modealat' })
    //   .then(res => {
    //     window.modek = res.data.modealat;
    //     window.license = res.data.license;
    //     if (window.modek === 'rfid') {
    //       menu = [
    //         { header: 'System' },
    //         { href: '/register', title: 'Register', icon: 'home' },
    //         { href: '/perangkat', title: 'Perangkat', icon: 'home' },
    //         { href: '/jaringanlan', title: 'jaringanlan', icon: 'home' },
    //         { href: '/mode', title: 'Mode Aplikasi', icon: 'home' }
    //       ];
    //     } else {
    //       // barcode
    //       menu = [
    //         { header: 'System' },
    //         { href: '/jaringanlan', title: 'jaringanlan', icon: 'home' },
    //         { href: '/mode', title: 'Mode Aplikasi', icon: 'home' }
    //       ];
    //     }
    //     console.log(window.modek);
    //     this.$store.commit('setLicense', window.license);
    //   });
    // this.$http.get('/users/1').then(({data}) => console.log(data))
    // global.$t = this.$t;
    // // fetch menu from server
    // this.$http.get('/menu').then(({ data }) => {
    //   this.$store.commit('setMenu', data);
    // });
    // this.$store.dispatch('checkPageTitle', this.$route.path);
    // this.$store.dispatch('checkAuth');
  }
});
