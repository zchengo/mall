import {createRouter, createWebHashHistory} from 'vue-router';
import Login from "../view/Login";
import Home from "../view/Home";
import Index from "../view/Index";
import Goods from "../view/Goods";
import Category from "../view/Category";
import Order from "../view/Order";
import Market from "../view/Market";

const routes = [
    {
        path: '/',
        component: Login,
    },
    {
        path: '/home',
        component: Home,
        redirect: '/index',
        children: [
            {
                path: '/index',
                name: 'index',
                component: Index,
            },
            {
                path: '/goods',
                name: 'goods',
                component: Goods,
            },
            {
                path: '/category',
                name: 'category',
                component: Category,
            },
            {
                path: '/order',
                name: 'order',
                component: Order,
            },
            {
                path: '/market',
                name: 'market',
                component: Market,
            }
        ],
    },
]

const router = createRouter({
    history: createWebHashHistory(), routes
})

export default router;