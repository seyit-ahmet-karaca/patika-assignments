<template>
  <div class="video-card">
    <video-image :video="filteredVideo"></video-image>
    <div class="video-information">
      <strong class="title">{{ filteredVideo.title }}</strong>
      <div class="display-details">
        <span class="view-count">{{ filteredVideo.viewCount }} görüntülenme </span>
        <span class="publish-date-in-month">{{ filteredVideo.publishDateInMonth }} ay önce</span>
      </div>
      <div class="owner-details">
        <img :src="filteredVideo.ownerImage" alt="" class="owner-image">
        <span class="owner-name">{{ filteredVideo.ownerName }}</span>
      </div>
      <div class="description">{{ filteredVideo.description }}</div>
    </div>
  </div>
</template>

<script>

import VideoImage from "./VideoImage";

export default {
  name: "VideoCard",
  props: {
    filteredVideo: Object
  },
  components: {
    VideoImage
  }

}
</script>

<style scoped>
.video-card {
  display: flex;
  margin: 10px 0;
}

.owner-details {
  display: flex;
}

.owner-image {
  height: 24px;
  width: 24px;
  border-radius: 50px;
}

.video-image {
  width: 360px;
  height: 200px;
}

.video-information {
  width: 100%;
  height: 100px;
  margin-left: 10px;
  margin-top: 5px;
}

.owner-name {
  margin-left: 10px;
  display: flex;
  align-items: center;
}

.video-information > * {
  margin-bottom: 10px;
}

.display-details {
  display: flex;
}

.publish-date-in-month {
  margin-left: 10px;
}
</style>