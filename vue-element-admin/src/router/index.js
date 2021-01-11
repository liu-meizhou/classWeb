import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

import visitorRouter from './modules/visitor'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/course'
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [

  visitorRouter,

  {
    path: '/my-courses',
    component: Layout,
    redirect: '/my-courses/index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/my-courses/index'),
        name: 'my-courses',
        meta: { title: '我的课表', icon: 'el-icon-notebook-2', noCache: true }
      }
    ]
  },

  {
    path: '/class',
    component: Layout,
    redirect: '/class/index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/class/index'),
        name: 'class',
        meta: { title: '班级查询', icon: 'peoples', noCache: true }
      }
    ]
  },

  {
    path: '/student',
    component: Layout,
    redirect: '/student/index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/student/index'),
        name: 'student',
        meta: { title: '学生查询', icon: 'user', noCache: true }
      }
    ]
  },

  {
    path: '/teacher',
    component: Layout,
    redirect: '/teacher/index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/teacher/index'),
        name: 'teacher',
        meta: { title: '老师查询', icon: 'user', noCache: true }
      }
    ]
  },

  {
    path: '/course',
    component: Layout,
    redirect: '/course/index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/course/index'),
        name: 'course',
        meta: { title: '课程查询', icon: 'education', noCache: true }
      }
    ]
  },

  {
    path: '/courseChoose',
    component: Layout,
    redirect: '/courseChoose/index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/course-choose/index'),
        name: 'courseChoose',
        meta: { title: '选课查询', icon: 'el-icon-circle-plus-outline', noCache: true }
      }
    ]
  },

  {
    path: '/icon',
    component: Layout,
    children: [
      {
        path: 'index',
        component: () => import('@/views/icons/index'),
        name: 'Icons',
        meta: { title: 'Icons', icon: 'icon', noCache: true }
      }
    ]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
