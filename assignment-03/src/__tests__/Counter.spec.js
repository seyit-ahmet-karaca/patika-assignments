import {createLocalVue, shallowMount} from "@vue/test-utils";
import Counter from "../Counter";
import Vuex from 'vuex';
import {actions, getters, mutations, state} from "../store";

function mountComponent(_state) {
    const localVue = createLocalVue();
    localVue.use(Vuex)

    return shallowMount(Counter, {
        store: new Vuex.Store({
            getters,
            state: _state === undefined ? state : _state,
            mutations,
            actions
        })
    })
}

describe("Counter.vue", () => {
    it('Component Exist Check', () => {
        const wrapper = mountComponent()

        expect(wrapper.exists()).toBeTruthy()
    });

    it('Increase button exist check', () => {
        const wrapper = mountComponent()

        const buttons = wrapper.findAll("button");
        expect(buttons).toHaveLength(2)

        expect(buttons.at(1).text()).toStrictEqual("Increase")
    });

    it('Decrease button exist check', () => {
        const wrapper = mountComponent()

        const buttons = wrapper.findAll("button");
        expect(buttons).toHaveLength(2)

        expect(buttons.at(0).text()).toStrictEqual("Decrease")
    });

    it('Increase button functionality check', () => {
        const localThis = {
            $store: {
                dispatch: jest.fn()
            }
        }
        Counter.methods.increase.call(localThis)

        expect(localThis.$store.dispatch).toHaveBeenCalled()
        expect(localThis.$store.dispatch).toHaveBeenCalledWith("increment")
    });

    it('Decrease button functionality check', () => {
        const localThis = {
            $store: {
                dispatch: jest.fn()
            }
        }
        Counter.methods.decrease.call(localThis)

        expect(localThis.$store.dispatch).toHaveBeenCalled()
        expect(localThis.$store.dispatch).toHaveBeenCalledWith("decrement")
    });

    it('2 increase + decrease functionality check together', async () => {
        const wrapper = mountComponent()
        const buttons = wrapper.findAll("button");
        const decrease = buttons.at(0);
        const increase = buttons.at(1);

        await increase.trigger("click")
        await increase.trigger("click")
        await decrease.trigger("click")

        expect(wrapper.find("span").text()).toStrictEqual("1k")
    });

    it('Count text show check', () => {
        const wrapper = mountComponent({count : 0})

        expect(wrapper.find("span").text()).toStrictEqual("0k")
    });
})