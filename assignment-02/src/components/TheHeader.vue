<template>
  <div class="header-container">
    <button class="favorite-button" @click="navigateFavorites">Favorites</button>
  </div>
</template>

<script>
export default {
  props: {
    favorites: {
      type: Array,
    }
  },
  data() {
    return {
      userId: "bootcamp",
    }
  },
  methods: {
    navigateFavorites() {
      this.$router.push({name: "favorites", params: {"userid": this.userId}});
    }
  }
}
</script>

<style scoped>
.header-container {
  width: 1540px;
  display: flex;
  justify-content: flex-end;
  margin: 20px 0;
}

.favorite-button {
  height: 40px;
  width: 150px;
  font-size: 23px;
}
</style>