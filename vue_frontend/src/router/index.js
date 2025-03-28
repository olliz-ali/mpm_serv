import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '@/view/HomeView.vue';
import LoginView from '@/view/LoginView.vue';
import SignupView from '@/view/SignupView.vue';


const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/', 
            name: 'Home',
            component: HomeView
        },
        {
            path: '/login',
            name: 'Login',
            component: LoginView
        },
        {
            path: '/signup',
            name: 'Signup',
            component: SignupView
        }
    ]
});


export default router;