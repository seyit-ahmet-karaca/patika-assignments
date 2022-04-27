import {createLocalVue, shallowMount} from "@vue/test-utils";
import App from "../App";
import Vuex from 'vuex';
import {getters, state} from "../store";


function mountComponent(stateParam) {
    const localVue = createLocalVue();
    localVue.use(Vuex);

    return shallowMount(App, {
        store: new Vuex.Store({
            getters,
            state: stateParam === undefined ? state : stateParam
        })
    })
}

function getRandomNumber(min, max) {
    return Math.floor(Math.random() * (max - min)) + min;
}

describe("App.vue", () => {

    test("h1 exists", () => {
        const wrapper = mountComponent()

        expect(wrapper.find("h1").exists()).toBeTruthy()
    })

    test("h1 text equals to Daily Corona Cases in Turkey check", () => {
        const wrapper = mountComponent()

        const h1ElementText = wrapper.find("h1").text();
        expect(h1ElementText).toStrictEqual("Daily Corona Cases in Turkey")
    })

    describe("notificationArea class check", () => {
        test("getCount value is less than 5", () => {

            const wrapper = mountComponent({count: getRandomNumber(0, 5)})

            const notificationAreaDiv = wrapper.find(".notificationArea");

            expect(notificationAreaDiv.attributes("class")).toContain("safe")
        })

        test("getCount value is between 5 and 10", () => {
            const wrapper = mountComponent({count: getRandomNumber(5, 10)})

            const notificationAreaDiv = wrapper.find(".notificationArea");

            expect(notificationAreaDiv.attributes("class")).toContain("normal")
        })

        test("getCount value is higher than 10", () => {
            const wrapper = mountComponent({count: getRandomNumber(10, Number.MAX_VALUE)})

            const notificationAreaDiv = wrapper.find(".notificationArea");

            expect(notificationAreaDiv.attributes("class")).toContain("danger")
        })
    })

    describe("notificationArea text message check", () => {
        it('notificationArea text is Danger!!! Case count is ${Count}k ', () => {
            const localThis = {
                $store: {
                    state: {
                        count: getRandomNumber(0, 5)
                    }
                }
            };

            const actualSafe = App.computed.message.call(localThis);
            expect(actualSafe).toStrictEqual(`So safe. Case count is ${localThis.$store.state.count}k`)

            localThis.$store.state.count = getRandomNumber(5, 10)
            const actualNormal = App.computed.message.call(localThis);
            expect(actualNormal).toStrictEqual(`Life is normal. Case count is ${localThis.$store.state.count}k`)

            localThis.$store.state.count = getRandomNumber(10, Number.MAX_VALUE)
            const actualDanger = App.computed.message.call(localThis);
            expect(actualDanger).toStrictEqual(`Danger!!! Case count is ${localThis.$store.state.count}k`)

        });


    })

})