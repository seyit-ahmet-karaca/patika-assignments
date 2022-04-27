<template>
  <div class="filtered-video-container">
    <filtered-video-card
        v-for="video in $store.getters.favoriteVideos" :key="video.id"
        :filteredVideo="video"
    ></filtered-video-card>
  </div>
</template>

<script>
import FilteredVideoCard from "../components/FavoriteVideoCard";

export default {
  name: "FavoritesPage",
  components: {
    FilteredVideoCard
  }
}
</script>

<style scoped>

.filtered-video-container {
  display: flex;
  flex-direction: column;
  height: 200px;
  width: 1540px;
  margin-bottom: 20px;
}
</style>