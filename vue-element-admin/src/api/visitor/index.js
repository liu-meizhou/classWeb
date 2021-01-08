import request from '@/utils/request_go'

export function get() {
  return request({
    url: '/visitor/v1/get',
    method: 'get'
  })
}

export function version() {
  return request({
    url: '/visitor/v1/version',
    method: 'get'
  })
}

export function show() {
  return request({
    url: '/show/1',
    method: 'get'
  })
}

export function login(data) {
  return request({
    url: '/visitor/v1/login',
    method: 'post',
    data
  })
}
