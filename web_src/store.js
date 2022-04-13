import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        logoText: "柏晓技术",
        logoMiniText: "柏晓",
        serverInfo: {},
        userInfo: null,
        menus: [
            {
                path: '/',
                title: "首页",
                icon: 'dashboard'
            },
            {
                path: '/tradideskconf/1',
                title: "桌面配置",
                icon: "video-camera"
            }, 
            {
                path: '/cloudesks/1',
                title: "桌面信息",
                icon: "cloud"
            },
            {
                path: '/deskcontrols',
                title: "桌面控制",
                icon: "cube"
            },
            {
                path: '/users/1',
                title: "用户管理",
                icon: "cube"
            },
            {
                path: "/about",
                icon: "support",
                title: "版本信息"                
            }
        ]
    },
    mutations: {
        updateServerInfo(state, serverInfo) {
            state.serverInfo = serverInfo;
        },
        updateUserInfo(state, userInfo) {
            state.userInfo = userInfo;
        }
    },
    actions : {
        getServerInfo({commit}){
            return new Promise((resolve, reject) => {
                $.get('/api/v1/serverinfo').then(serverInfo => {
                    commit('updateServerInfo', serverInfo);
                    resolve(serverInfo);
                }).fail(() => {
                    resolve(null);
                });
            })
        },
        getUserInfo({ commit, state }) {
            return new Promise((resolve, reject) => {
                $.get("/api/v1/userinfo").then(userInfo => {
                    commit('updateUserInfo', userInfo);
                    resolve(userInfo);
                }).fail(() => {
                    resolve(null);
                })
            })
        },   
        logout({ commit, state }) {
            return new Promise((resolve, reject) => {
                $.get('/api/v1/logout').then(data => {
                    commit('updateUserInfo', null);
                    resolve(null);
                }).fail(() => {
                    resolve(null);
                })
            })
        }
    }
})

export default store;
