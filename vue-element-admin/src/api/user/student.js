import request from '@/utils/request_go'

export function editStudentClass(studentId, classId) {
  return request({
    url: '/user/v1/student/editClass',
    method: 'get',
    params: {
      studentId: studentId,
      classId: classId
    }
  })
}

export function getStudentList() {
  return request({
    url: '/user/v1/student/list',
    method: 'post'
  })
}
