import request from '@/utils/request_go'

export function getTeacherById(classId) {
  return request({
    url: '/user/v1/class/info',
    method: 'get',
    params: { classId }
  })
}

export function getTeachers() {
  return request({
    url: '/user/v1/teacher/list',
    method: 'get'
  })
}
