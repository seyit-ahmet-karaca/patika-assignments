import VueRouter from "vue-router";
import Vue from "vue";
import VideoListingPage from "../views/VideoListingPage";
import FavoritesPage from "../views/FavoritesPage";
import VideoPlayPage from "../views/VideoPlayPage";

Vue.use(VueRouter);

export const routes = [
    {path: "/", component: VideoListingPage},
    {name: "favorites", path: "/Favorites/:userid", component: FavoritesPage},
    {path: "/watch", component: VideoPlayPage}
]

const router = new VueRouter({
    mode: "history",
    routes
});

export default router;