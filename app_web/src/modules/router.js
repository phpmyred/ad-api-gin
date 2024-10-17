import {createRouter,createWebHashHistory } from "vue-router"
import { setupLayouts } from 'virtual:generated-layouts'
import generatedRoutes from '~pages'
import Index from "@/pages/index.vue";


// const routes = setupLayouts(generatedRoutes)

const routes = [
    { path: '/', component: Index,
        meta: {
            title: "首页"

        }
    },


];


const router = createRouter({
    routes,
    history: createWebHashHistory()

})


export default router;