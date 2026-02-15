import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0, behavior: 'smooth' }
    }
  },
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue')
      // Removed requiresAuth to make landing page public
    },
    {
      path: '/home',
      redirect: '/'
    },
    {
      path: '/donors/:id',
      name: 'donor-profile',
      component: () => import('../views/PublicDonorProfileView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('../views/DonorProfileView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/search',
      name: 'search',
      component: () => import('../views/DonorSearchView.vue')
    },
    // Admin Routes
    {
      path: '/admin',
      component: () => import('../views/admin/AdminLayout.vue'),
      meta: { requiresAuth: true, requiresAdmin: true },
      children: [
        {
          path: '',
          name: 'admin-dashboard',
          component: () => import('../views/admin/DashboardView.vue')
        },
        {
          path: 'donors',
          name: 'admin-donors',
          component: () => import('../views/admin/DonorsView.vue')
        },
        {
          path: 'donors/:id/edit',
          name: 'admin-donor-edit',
          component: () => import('@/views/admin/AdminUserEditView.vue')
        }
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  // Hydrate auth state if needed (e.g. on refresh)
  if (!authStore.user && authStore.token) {
      authStore.hydrate();
  }

  const isAuthenticated = authStore.isAuthenticated;

  // Prevent logged in users from visiting login/register
  if ((to.name === 'login' || to.name === 'register') && isAuthenticated) {
    return next('/');
  }

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else if (to.meta.requiresAdmin && (!authStore.user || authStore.user.role !== 'admin')) {
    // Redirect non-admins to home
    next('/')
  } else {
    next()
  }
})

export default router
