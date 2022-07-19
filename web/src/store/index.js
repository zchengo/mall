import { createStore } from 'vuex'

const store = createStore({
    state() {
        return {
            
            // 默认激活
            defaultActive: '/index',

            // 动态导航
            navigation: [{
                name: '我的主页',
                index: '/index',
            }],

            // 结果页标题状态
            pageTitle: ''
        }
    },
    mutations: {
        setMenu(state, index){
            state.defaultActive = index
        },
        addNav(state, item) {
            state.navigation = item
        },
        setPageTitle(state, title) {
            state.pageTitle = title
        }
    },
    actions: {},
    modules: {}
})

export default store;