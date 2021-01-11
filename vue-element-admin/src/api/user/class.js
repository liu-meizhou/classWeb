import request from '@/utils/request_go'

export function getClass(classId) {
  return request({
    url: '/user/v1/class/info',
    method: 'get',
    params: { classId }
  })
}

export function editClass(data) {
  return request({
    url: '/user/v1/class/info',
    method: 'post',
    data
  })
}

export function deleteClass(classId) {
  return request({
    url: '/user/v1/class/delete',
    method: 'get',
    params: { classId }
  })
}

export function createClass(data) {
  return request({
    url: '/user/v1/class/create',
    method: 'post',
    data
  })
}

export function getClassList() {
  return request({
    url: '/user/v1/class/list',
    method: 'get'
  })
}
