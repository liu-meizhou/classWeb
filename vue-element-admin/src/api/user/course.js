import request from '@/utils/request_go'

export function getCourse() {
  return request({
    url: '/user/v1/course/show',
    method: 'get'
  })
}

export function chooseCourse(courseId) {
  return request({
    url: '/user/v1/course/choose',
    method: 'get',
    params: { courseId }
  })
}

export function chooseCourseList(data) {
  return request({
    url: '/user/v1/course/choose',
    method: 'post',
    data
  })
}

export function getStudent(courseId) {
  return request({
    url: '/user/v1/course/grade',
    method: 'get',
    params: { courseId }
  })
}
