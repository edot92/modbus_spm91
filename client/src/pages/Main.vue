<template>
  <v-app :dark="dark" standalone>
    <v-navigation-drawer persistent v-model="drawer" enable-resize-watcher app :dark="dark">
      <div v-show="!mini" class="pa-3 text-xs-center">
        <div class="display-2 py-4">BUDI SEPTIANTO</div>
        <p>TELKOM</p>
        <!-- <div style="padding-left:5em">
          <v-switch :label="(!dark ? 'Light' : 'Dark') + ' Theme'" v-model="dark" :dark="dark" hide-details></v-switch>
        </div> -->
        <div>
        </div>
      </div>
      <div v-show="mini" class="pa-3 text-xs-center">
        <div class="display-2">A</div>
      </div>
      <v-divider></v-divider>
      <v-list dense>
        <template v-for="item in menu">
          <v-list-group v-if="item.items" v-bind:group="item.group">
            <v-list-tile :to="item.href" slot="item" :title="item.title">
              <v-list-tile-action>
                <v-icon>{{ item.icon }}</v-icon>
              </v-list-tile-action>
              <v-list-tile-content>
                <v-list-tile-title>{{ $t(item.title) }}</v-list-tile-title>
              </v-list-tile-content>
              <v-list-tile-action>
                <v-icon>keyboard_arrow_down</v-icon>
              </v-list-tile-action>
            </v-list-tile>
            <v-list-tile v-for="subItem in item.items" :key="subItem.href" :to="subItem.href" v-bind:router="!subItem.target" ripple v-bind:disabled="subItem.disabled" v-bind:target="subItem.target">
              <v-list-tile-action v-if="subItem.icon">
                <v-icon class="success--text">{{ subItem.icon }}</v-icon>
              </v-list-tile-action>
              <v-list-tile-content>
                <v-list-tile-title>{{ $t(subItem.title) }}</v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
          </v-list-group>
          <v-subheader v-else-if="item.header">{{ item.header }}</v-subheader>
          <v-divider v-else-if="item.divider"></v-divider>
          <v-list-tile v-else :to="item.href" router ripple v-bind:disabled="item.disabled" :title="item.title">
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>{{ $t(item.title) }}</v-list-tile-title>
            </v-list-tile-content>
            <v-list-tile-action v-if="item.subAction">
              <v-icon class="success--text">{{ item.subAction }}</v-icon>
            </v-list-tile-action>
          </v-list-tile>
        </template>

      </v-list>
    </v-navigation-drawer>

    <v-toolbar fixed dark :class="theme" class="darken-1" dark fixed app>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>

      <!-- <v-toolbar-side-icon dark @click.stop="hideShow()"></v-toolbar-side-icon> -->
      <v-toolbar-title>{{$t(pageTitle)}}</v-toolbar-title>

      <!-- <v-toolbar-title style=" margin-left: 20%;"></v-toolbar-title> -->

      <v-spacer></v-spacer>
      <!-- <v-menu offset-y>
                                                                                                      </v-menu> -->
      <v-menu offset-y>
        <v-btn icon dark slot="activator">
          <v-icon dark>format_paint</v-icon>
        </v-btn>
        <v-list>
          <v-list-tile v-for="n in colors" :key="n" :class="n" @mouseover.native="theme = n"></v-list-tile>
        </v-list>
      </v-menu>

    </v-toolbar>

    <main>
           <v-content>
      <v-container fluid class="pa-4">
        <v-alert v-bind="message" v-model="message.body" dismissible>{{message.body}}</v-alert>
        <div class="py-2">
          <v-slide-y-transition mode="out-in">
            <router-view></router-view>
          </v-slide-y-transition>
        </div>
      </v-container>
           </v-content>
    </main>
    <v-footer  :class="theme" app>
      <span style="    color: white;">&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
<template>
 <loading
     :show="show"
     :label="label">
 </loading>
</template>
  </v-app>
</template>

<script>
import { mapState } from 'vuex';
import loading from 'vue-full-loading';
// var server = 'http://103.23.20.159';
export default {
  components: {
    loading
  },
  data () {
    return {
      show: false,
      label: 'Loading...',
      theme: 'primary',
      dark: false,
      mini: false,
      drawer: true,
      locales: ['en-US', 'zh-CN'],
      colors: ['blue', 'green', 'purple', 'red'],
      mode: ''
    };
  },
  computed: {
    ...mapState(['message', 'menu', 'pageTitle', 'loading'])
  },
  watch: {
    theme (param) {},
    loading (param) {
      this.show = param;
    }
  },
  beforeCreate () {
    // const token = this.$store.state.token;
    // if (token === null || token === undefined) {
    //   this.$router.push('/login');
    // }
  },
  methods: {
    httpLoginTemp () {
      const thisVue = this;
      this.$http({
        url: '/user/login',
        method: 'post',
        data: {
          username: 'admin',
          password: 'admin'
        },
        withCredentials: true
      })
        .then(res => {
          if (res.data.error === false) {
            // window.swal('info', res.data.message, 'info');
            const datak = res.data.payload;
            const save = { user: datak, token: datak.jwt };
            thisVue.$http.defaults.headers.common['Jwt'] = datak.jwt;
            thisVue.$store.commit('setAuth', save);
            thisVue.$store.commit('setReadyLoad', true);
          } else {
            window.swal('error', res.data.message, 'error');
            console.error(res.data.line);
          }
        })
        .catch(err => {
          console.error(err.toString());
          window.swal('error', err.toString(), 'error');
        });
    },
    firstLoad () {},
    hideShow () {
      this.drawer = !this.drawer;
    },
    changeLocale (to) {
      global.helper.ls.set('locale', to);
      this.$i18n.locale = to;
    },
    fetchMenu () {
      // fetch menu from server
      // this.$http.get('menu').then(({data}) => this.$store.commit('setMenu', data))
    }
  },

  created () {
    this.fetchMenu();
  },
  mounted () {
    // const thisVue = this;
    // thisVue.firstLoad();
  }
};
</script>
