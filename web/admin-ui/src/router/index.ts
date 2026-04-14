import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import MainLayout from '../layouts/MainLayout.vue'

const placeholder = (title: string) => ({ template: `<div><h2>${title}</h2><p>MVP scaffold page.</p></div>` })

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: MainLayout,
    children: [
      { path: '', redirect: '/dashboard' },
      { path: 'dashboard', component: placeholder('Dashboard') },
      { path: 'hosts', component: placeholder('Hosts') },
      { path: 'hosts/:id', component: placeholder('Host Detail') },
      { path: 'packs', component: placeholder('Packs') },
      { path: 'packs/:id', component: placeholder('Pack Detail') },
      { path: 'projects', component: placeholder('Projects') },
      { path: 'projects/:id', component: placeholder('Project Detail') },
      { path: 'projects/:id/designer', component: placeholder('Topology Designer') },
      { path: 'projects/:id/install', component: placeholder('Install Config') },
      { path: 'projects/:id/generated', component: placeholder('Generated Output') },
      { path: 'deployments', component: placeholder('Deployments') },
      { path: 'deployments/:id', component: placeholder('Deployment Detail') },
      { path: 'runtime', component: placeholder('Runtime') },
      { path: 'runtime/:id', component: placeholder('Runtime Detail') },
      { path: 'runtime/:id/logs', component: placeholder('Logs') },
      { path: 'runtime/:id/terminal', component: placeholder('Terminal') },
      { path: 'metrics', component: placeholder('Metrics') },
      { path: 'settings', component: placeholder('Settings') }
    ]
  }
]

export const router = createRouter({
  history: createWebHistory(),
  routes
})
