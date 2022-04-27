<template>
  <div id="app">
    <the-header></the-header>
    <router-view></router-view>
  </div>
</template>

<script>
import { mapActions} from 'vuex';
import TheHeader from "./components/TheHeader";
import api_url from "./constants";

export default {
  name: 'App',
  components:{
    TheHeader
  },
  methods:{
    ...mapActions(['initVideos']),
  },
  mounted() {
    console.log(api_url)
    this.initVideos();
  }
}
</script>

<style>
#app {
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 100%;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin-top: 60px;
}
</style>
