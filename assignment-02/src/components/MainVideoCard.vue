<template>
  <div class="video-card">
    <video-image :video="video" :isListing=true></video-image>
    <div class="video-details">
      <img :src="video.ownerImage" alt="" class="owner-image">
      <div class="owner-information">
        <strong class="title">{{ video.title }}</strong>
        <span class="owner-name">{{ video.ownerName }}</span>
        <div class="display-details">
          <span class="view-count">
            {{ video.viewCount }} görüntüleme
          </span>
          <span class="publish-date-in-month">{{ video.publishDateInMonth }} ay önce</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import VideoImage from "./VideoImage";

export default {
  name: "ListedVideoCard",
  props: {
    video: Object
  },
  components: {
    VideoImage
  },
}
</script>

<style scoped>
.video-card {
  background-color: #ffffff;
  margin: 10px 0;
  display: flex;
  flex-direction: column;
  width: 360px;
}

.favorite {
  background-color: green;
}

.video-details {
  width: 100%;
  display: flex;
  margin: 10px 10px 0 0;
}

.owner-image {
  width: 36px;
  height: 36px;
  border-radius: 50px;
  margin-right: 10px;
}

.title {
  display: flex;
  justify-content: flex-start;
  width: 100%;
}

.owner-name {
  display: flex;
  justify-content: flex-start;
}

.display-details {
  display: flex;
  justify-content: flex-start;
}

.owner-information {
  display: flex;
  flex-direction: column;
  width: 100%;
  justify-content: flex-start;
}

.owner-information > * {
  margin-bottom: 3px;
}

.publish-date-in-month {
  margin-left: 10px;
}
</style>