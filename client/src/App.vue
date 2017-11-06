<template>
  <div>
    <transition mode="out-in">
      <router-view></router-view>
    </transition>
  </div>
</template>

<script>
export default {
  data () {
    return {

    }
  },
  methods: {

  },
  created () {
    // this.$http.get('/users').then(({
    //     data
    //   }) => console.log(data))
  },
  mounted () {
    // window.socket.on('event', function (data) {});
  }
}

</script>
