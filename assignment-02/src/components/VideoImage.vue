<template>
  <div class="container">
    <font-awesome-icon icon="heart" v-if="isListing" @click="addFavorite" id="fav-button"
                       :class="{favorite: isFavorite, 'not-favorite': !isFavorite, 'fa-icon':true}"/>
    <router-link :to="{path: '/watch' , query:{id : video.id}}">
      <img :src="coverImageChange" alt="" :class="{'listing-image': isListing, 'filtered-image': !isListing}"
           @mouseover="isCover = false" @mouseleave="isCover = true">
    </router-link>
  </div>
</template>

<script>
import {mapMutations} from 'vuex';
import {mapState} from 'vuex';

export default {
  name: "VideoImage",
  props: {
    video: Object,
    isListing: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      isCover: true,
      isFavorite: false
    }
  },
  methods: {
    addFavorite() {
      this.isFavorite = !this.isFavorite;
      this.manageFavoriteVideoIds(this.video.id);
    },
    ...mapMutations(["manageFavoriteVideoIds"]),
    ...mapState(["favoriteVideoIds"]),
  },
  computed: {
    coverImageChange() {
      if (this.isCover) {
        return this.video.coverImage;
      } else {
        return this.video.hoverImage;
      }
    }
  },
  mounted() {
    this.isFavorite = this.favoriteVideoIds().includes(this.video.id);
  }
}
</script>

<style scoped>
.listing-image {
  width: 100%;
  height: 200px;
}

.filtered-image {
  width: 360px;
  height: 200px;
}

.fa-icon {
  margin: 5px;
  border-radius: 100px;
  box-shadow: 0 0 8px 4px white;
  font-size: 20px;
  background-color: white;
}

.not-favorite {
  transition: 0.3s;
  color: rgba(220, 20, 60, 0.2);
}

.favorite {
  box-shadow: 0 0 8px 4px white;
  color: rgba(220, 20, 60);
  font-size: 30px;
  transition: 0.3s;
}

.container {
  position: relative;
}

#fav-button {
  position: absolute;
  right: 0;
  top: 0;
  z-index: 1;
  padding: 5px;
  font-weight: bold;
}
</style>