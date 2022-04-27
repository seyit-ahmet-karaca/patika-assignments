import {actions, mutations, state} from "../store";
import VideosAPI from "../../services/VideosAPI";
import flushPromises from "flush-promises";

jest.mock('../../services/VideosAPI')


describe("foo", () => {
    describe("mutations", () => {
        it('manageFavoriteVideoIds remove', () => {
            state.favoriteVideoIds.push([1, 2, 3]);
            const favoriteVideoIdWantToRemove = 1

            mutations.manageFavoriteVideoIds(state, favoriteVideoIdWantToRemove)

            expect(state.favoriteVideoIds).toHaveLength(2)
        });

        it('manageFavoriteVideoIds add', () => {
            const favoriteVideoIdWantToAdd = 1
            mutations.manageFavoriteVideoIds(state, favoriteVideoIdWantToAdd)

            expect(state.favoriteVideoIds).toHaveLength(1)
        });
    })

    describe("actions", () => {
        it('should ', async () => {
            const mockVideos = [
                {id: 1},
                {id: 2},
                {id: 3},
                {id: 4}
            ]

            const context = {
                commit: jest.fn()
            }

            VideosAPI.getVideoList.mockResolvedValue(mockVideos)
            actions.initVideos(context);
            await flushPromises();

            expect(context.commit).toHaveBeenCalledWith('setVideos', mockVideos)
        });
    })
})
