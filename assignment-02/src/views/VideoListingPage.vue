<template>
  <div class="main-listing-container">
    <div class="video-listing-container">
      <listed-video-card
          v-for="video in $store.state.videos" :key="video.id"
          :video="video"
      ></listed-video-card>
    </div>
  </div>
</template>

<script>
import ListedVideoCard from "../components/MainVideoCard";

export default {
  name: "VideoListPage",
  components: {
    ListedVideoCard
  }
}
</script>

<style scoped>
.main-listing-container {
  display: flex;
  width: 1540px;
  flex-direction: column;
}

.video-listing-container {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

</style>