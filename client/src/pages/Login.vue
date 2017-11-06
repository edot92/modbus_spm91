<style scoped>
.ui-button {
  width: 100%;
}
/* @media screen and (min-device-width: 1200px) and (max-device-width: 1600px) and (-webkit-min-device-pixel-ratio: 2) and (min-resolution: 192dpi) {
  .dialog {
    width: "40%" !important;
  }
}

@media only screen and (max-width: 451px) {
  .dialog {
    width: "80%" !important;
  }
} */
</style>

<template>
  <div>

    <v-container grid-list-md text-xs-center>
      <br>
      <!-- <v-flex xs12> -->
      <v-layout row justify-center>

        <v-dialog style="margin-top:-50px;" class="form-uhuy" :value="true" persistent  width= "40%"  >
          <v-card hover="" style="background:white">
            <v-card-text class="pt-12">
              <v-flex xs12>
                <v-card class=" white--text">
                  <v-container fluid grid-list-lg>
                    <v-layout row>
                      <v-flex xs7>
                        <!-- <div>
                          <div class="headline">
                            <h1>Login</h1>
                            </div>
                        </div> -->
                      </v-flex>
                      <v-flex xs5>
                        <v-card-media src="https://previews.123rf.com/images/bitontawan02/bitontawan021611/bitontawan02161100013/66162960-Cloud-IOT-Internet-of-Things-Icon-and-symbol-Stock-Vector.jpg" height="155px" contain></v-card-media>
                      </v-flex>
                    </v-layout>
                  </v-container>
                </v-card>
              </v-flex>

              <v-layout row>
                <v-flex xs4>
                  <v-subheader style="margin-top:10px;">Username:</v-subheader>
                </v-flex>
                <v-flex xs8>
                  <v-text-field v-model="form.username" name="input-11-2" label="Username Anda" hint="At least 8 characters" min="8"></v-text-field>
                </v-flex>
              </v-layout>
              <v-layout row>
                <v-flex xs4>
                  <v-subheader style="margin-top:10px;">Password:</v-subheader>
                </v-flex>
                <v-flex xs8>
                  <v-text-field v-model="form.password" name="input-11-2" label="Password Anda" hint="At least 8 characters" min="8" :append-icon="form.password ? 'visibility' : 'visibility_off'" :append-icon-cb="() => (form.hidepasspassword = !form.hidepasspassword)" value="wqfasds" type="password" class="input-group--focused" :type="form.hidepasspassword ? 'password' : 'text'"></v-text-field>
                </v-flex>
              </v-layout>

              <v-layout row>

                <v-btn
                style="height:5em"
      color="primary"
      :loading="hideBtn"
      @click="cekLogin()"
      :disabled="hideBtn"
block
    >

     Login
             <v-icon  >arrow_forward
</v-icon>
    </v-btn>
              </v-layout>

              <!-- </v-form> -->
            </v-card-text>
            </v-card-row>
          </v-card>
        </v-dialog>
      </v-layout>

      <!-- </v-flex> -->

    </v-container>

  </div>
</template>
<script>
export default {
  data () {
    return {
      hideBtn: false,
      model: '',
      form: {
        hidepasspassword: false,
        username: '',
        password: ''
      }
    };
  },

  methods: {
    cekLogin () {
      const thisVue = this;
      this.hideBtn = true;
      setTimeout(function () {
        thisVue.hideBtn = false;
      }, 2000);
      if (this.form.username === '' || this.form.password === '') {
        window.swal('Info', 'harap isi parameter', 'error');
        return;
      }
      const paramBody = {
        username: this.form.username,
        password: this.form.password
      };
      this.$http({
        method: 'post',
        url: '/user/login',
        data: paramBody
      })
        .then(response => {
          thisVue.$store.commit('setLoading', false);
          if (response.data.error === true) {
            const strError = response.data.message;
            if (strError === 'not found') {
              window.swal('Info', 'username / password tidak sesuai', 'error');
            } else {
              window.swal('Info', response.data.message, 'error');
            }
          } else {
            window.swal('Info', response.data.message, 'success');
            const datak = response.data.payload;
            const save = { user: datak, token: datak.jwt };
            thisVue.$http.defaults.headers.common['Jwt'] = datak.jwt;
            thisVue.$store.commit('setAuth', save);
            thisVue.$store.commit('setReadyLoad', true);
            thisVue.goto('/');
          }
        })
        .catch(error => {
          thisVue.$store.commit('setLoading', false);
          window.swal('Error jaringan', error.toString(), 'error');
        });
    },
    onSuccess () {
      alert('sukses');
    },
    goto (param) {
      this.$router.push(param);
    }
  }
};
</script>
