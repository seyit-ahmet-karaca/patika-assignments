import {createLocalVue, shallowMount} from "@vue/test-utils";
import VideoListingPage from "../VideoListingPage";
import Vuex from 'vuex';
import {actions, getters, mutations, state} from "../../store/store";
import videosAPI from "../../services/VideosAPI";
import FilteredVideoCard from "../../components/FavoriteVideoCard";

jest.mock("../../services/VideosAPI");

describe("foo", () => {
    it('video listing page exists ', () => {
        const localVue = createLocalVue();
        localVue.use(Vuex);
        videosAPI.getVideoList.mockResolvedValue([]);

        const wrapper = shallowMount(VideoListingPage,
            {
                store : new Vuex.Store({
                    state,
                    getters,
                    mutations,
                    actions
                })
            })

        expect(wrapper.exists()).toBeTruthy()
    });

    it('video listing page exists ', () => {
        const localVue = createLocalVue();
        localVue.use(Vuex);
        videosAPI.getVideoList.mockResolvedValue([]);

        const wrapper = shallowMount(VideoListingPage,
            {
                store : new Vuex.Store({
                    state,
                    getters,
                    mutations,
                    actions
                })
            })
        const vueWrapperArray = wrapper.findAllComponents(FilteredVideoCard);
        expect(vueWrapperArray).toHaveLength(0)
    });

})