import request from '@/utils/request_go'

export function getCourse() {
  return request({
    url: '/user/v1/course/show',
    method: 'get'
  })
}

export function chooseCourseList() {
  return request({
    url: '/user/v1/course/choose',
    method: 'get'
  })
}

export function getStudent(courseId) {
  return request({
    url: '/user/v1/course/grade',
    method: 'get',
    params: { courseId }
  })
}

export function logout() {
  return request({
    url: '/user/v1/logout',
    method: 'get'
  })
}
