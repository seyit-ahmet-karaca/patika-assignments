import {createLocalVue, shallowMount} from "@vue/test-utils";
import VideoImage from "../VideoImage";
import Vuex from 'vuex';
import {getters, mutations, state} from "../../store/store";
import VueRouter from "vue-router";
import {routes} from "../../router/router";

describe("video image", () => {
    it('should exists', () => {
        const localVue = createLocalVue();
        localVue.use(Vuex)

        const wrapper = shallowMount(VideoImage, {
            propsData: {
                video: {}
            },
            store: new Vuex.Store({
                getters,
                state,
                mutations
            })
        })
        expect(wrapper.exists())
    });

    it('should mouseover', () => {
        const localThis = {
            isCover: true,
            video: {
                coverImage: 'coverImage',
                hoverImage: 'hoverImage'
            }
        }
        const retCover = VideoImage.computed.coverImageChange.call(localThis);
        expect(retCover).toStrictEqual(localThis.video.coverImage)

        localThis.isCover = false;
        const retHover = VideoImage.computed.coverImageChange.call(localThis);
        expect(retHover).toStrictEqual(localThis.video.hoverImage)

    });

    it('should render correctly', () => {
        // todo : sor
        state.favoriteVideoIds.push(...[1,2,3])
        const localVue = createLocalVue();
        localVue.use(Vuex)

        const wrapper = shallowMount(VideoImage, {
            propsData: {
                video: {
                    id: 1
                },
                isListing: true
            },
            store: new Vuex.Store({
                getters,
                state,
                mutations
            })
        })
        const favButton = wrapper.find('#fav-button');
        expect(favButton.exists()).toBeTruthy()

        const classAttribute = favButton.attributes("class");
        expect(classAttribute).toStrictEqual("favorite fa-icon")

    });

    it('should clicked', async function () {
        state.favoriteVideoIds.push(...[1,2,3])
        const localVue = createLocalVue();
        localVue.use(Vuex)
        localVue.use(VueRouter)

        const wrapper = shallowMount(VideoImage, {
            propsData: {
                video: {
                    id: 1
                },
                isListing: true
            },
            store: new Vuex.Store({
                getters,
                state,
                mutations
            }),
            router: new VueRouter({
                routes
            })
        })
        const mockMethod = jest.fn();
        wrapper.setMethods({'addFavorite': mockMethod})

        await wrapper.find("#fav-button").trigger("click")

        expect(mockMethod).toHaveBeenCalled()
    });

    it('should addFavorite', function () {
        const localThis = {
            isFavorite: true,
            manageFavoriteVideoIds : jest.fn(),
            video: {
                id: 1
            }
        }
        VideoImage.methods.addFavorite.call(localThis)
        expect(localThis.manageFavoriteVideoIds).toHaveBeenCalled()
        expect(localThis.manageFavoriteVideoIds).toHaveBeenCalledWith(localThis.video.id)
    });
})