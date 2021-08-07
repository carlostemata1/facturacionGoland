import { createRouter, createWebHistory } from 'vue-router'
import Catalogo from '../views/Catalogo.vue'

const routes = [
  {
    path: '/',
    name: 'Catalogo',
    component: Catalogo
  },
  {
    path: '/promocion',
    name: 'Promocion',
    component: () => import('../views/Promocion.vue')
  },
  {
    path: '/factura',
    name: 'Factura',
    component: () => import('../views/Factura.vue')
  },
  {
    path: '/contacto',
    name: 'Contacto',
    component: () => import('../views/Contacto.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
