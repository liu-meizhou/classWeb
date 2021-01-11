import request from '@/utils/request_go'

export function logout() {
  return request({
    url: '/user/v1/logout',
    method: 'get'
  })
}
