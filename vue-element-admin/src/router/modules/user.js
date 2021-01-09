/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/layout'

const visitorRouter = {
  path: '/user',
  component: Layout,
  redirect: 'noRedirect',
  name: '用户',
  meta: {
    title: '用户',
    icon: 'eye-open'
  },
  children: [
    {
      path: 'index',
      component: () => import('@/views/visitor/index'),
      name: 'index',
      meta: { title: '游客', noCache: true }
    }
  ]
}

export default visitorRouter
