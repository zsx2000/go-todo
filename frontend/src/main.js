import Vue from 'vue'
import { BootstrapVue, BootstrapVueIcons } from "bootstrap-vue"
import App from './App.vue'
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap-vue/dist/bootstrap-vue.css"
import VueProgressBar from 'vue-progressbar'

Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)


Vue.use(VueProgressBar)

new Vue({
    render: h => h(App),
}).$mount('#app')