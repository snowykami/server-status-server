import {createRouter, createWebHistory} from "vue-router"
import Home from "../views/Home.vue"
import Test from "../views/Test.vue"

const routes = [
    {
        path: '/',
        name: '主页',
        component: Home
    },
    {
        path: '/test',
        name: '测试',
        component: Test
    }
]

const router = createRouter({
    history: createWebHistory("/"),
    routes
})

export default router