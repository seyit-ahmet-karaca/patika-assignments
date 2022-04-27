<template>
  <h1>VİDEO ID {{ videoId }}</h1>
</template>

<script>
export default {
  name: "VideoPlayPage",
  data() {
    return {
      videoId: undefined
    }
  },
  mounted() {
    this.videoId = this.$route.query.id;
  }
}
</script>

<style scoped>

</style>