import Vue from 'vue';
import Router from 'vue-router';

function route (path, file, name, children) {
  return {
    exact: true,
    path,
    name,
    children,
    component: require(`./pages/${file}.vue`)
  };
}

function route2 (path, sub, file, name, children) {
  return {
    exact: true,
    path,
    name,
    children,
    component: require(`./pages/budi${sub}${file}.vue`)
  };
}
Vue.use(Router);

const router = new Router({
  base: __dirname,
  mode: 'hash',
  scrollBehavior: () => ({
    y: 0
  }),
  routes: [
    route('/login', 'Login', 'login'),
    route('/error', 'Error', 'error'),
    {
      exact: false,
      path: '/',
      redirect: '/realtimeindex'
    },
    route('/forbiden', 'Forbiden', '404 Forbiden'),
    route('/', 'Main', null, [
      // route2('/laporanindex', '/laporan', '/LaporanIndex', 'Laporan'),
      /** ************************************************** */
      // route2('/userindex', '/user', '/UserIndex', 'Users'),
      // route2('/user/buatbaru', '/user', '/UserBuatBaru', 'Buat Nama user'),
      // route2('/user/edit', '/user', '/UserEdit', 'Edit User'),
      /** ************************************************** */
      // route2('/garduindex', '/gardu', '/GarduIndex', 'Gardu'),
      // route2('/gardu/buatbaru', '/gardu', '/GarduBuatBaru', 'Buat Nama Gardu'),
      // route2('/gardu/edit', '/gardu', '/GarduEdit', 'Edit Gardu'),

      /** ******************** */
      route2('/realtimeindex', '/realtime', '/RealtimeIndex', 'Realtime')
    ]),

    // Global redirect for 404
    {
      path: '*',
      redirect: '/error',
      query: {
        code: 404,
        message: 'Page Not Found.'
      }
    }
  ]
});

router.beforeEach((to, from, next) => {
  global.store.dispatch('checkPageTitle', to.path);
  /* eslint-disable no-undef */
  if (typeof ga !== 'undefined') {
    ga('set', 'page', to.path);
    ga('send', 'pageview');
  }
  next();
});

export default router;
