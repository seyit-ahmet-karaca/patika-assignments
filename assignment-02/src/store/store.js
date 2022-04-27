import Vuex from 'vuex';
import Vue from 'vue';
import VideosAPI from "../services/VideosAPI";

Vue.use(Vuex);

export const state = {
    videos: [],
    favoriteVideoIds: []
};

export const getters = {
    favoriteVideos(state) {
        return state.videos.filter(video => state.favoriteVideoIds.includes(video.id))
    },
    getVideos(state) {
        return state.videos;
    }
};

export const mutations = {
    setVideos(state, payload) {
        state.videos.push(...payload);
    },
    manageFavoriteVideoIds(state, payload) {
        const isVideoIn = state.favoriteVideoIds.some(id => id === payload);
        if (isVideoIn) {
            state.favoriteVideoIds.splice(state.favoriteVideoIds.indexOf(payload), 1);
        } else {
            state.favoriteVideoIds.push(payload);
        }
    }
};

export const actions = {
    async initVideos({commit}) {
        const videos = await VideosAPI.getVideoList();
        commit('setVideos', videos);
    }
};

const store = new Vuex.Store({
    state,
    getters,
    mutations,
    actions
});

export default store;