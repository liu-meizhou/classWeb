<template>
  <el-table :data="students" style="width: 100%">
    <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="成绩">
            <span>{{ props.row.courseResult }}</span>
          </el-form-item>
          <el-form-item label="绩点">
            <span>{{ props.row.coursePoint }}</span>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column align="center" label="学号" prop="id" />
    <el-table-column align="center" label="姓名" prop="name" />
    <el-table-column align="center" label="性别" prop="sex" />
    <el-table-column align="center" label="班级" prop="class.name" />
    <el-table-column align="center" label="学院" prop="college" />
    <el-table-column label="操作">
      <template slot-scope="scope">
        <el-button
          size="mini"
          @click="handleEditStudent(scope.$index, scope.row)"
        >编辑</el-button>
        <el-button
          size="mini"
          type="danger"
          @click="handleDeleteStudent(scope.$index, scope.row)"
        >删除</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { getStudent } from '@/api/user/course'
import { getStudentList } from '@/api/user/student'
export default {
  data() {
    return {
      students: []
    }
  },
  created() {
    const query = this.$route.query
    if (query.courseId) {
      this.getStudent(query.courseId)
    } else {
      this.getAllStudent()
    }
  },
  methods: {
    getAllStudent() {
      getStudentList()
        .then(pageInfo => {
          this.students = pageInfo.lists
          console.log(this.students)
        })
        .catch(err => {
          console.log(err)
        })
    },
    getStudent(courseId) {
      getStudent(courseId)
        .then(students => {
          this.students = students
          console.log(students)
        })
        .catch(err => {
          console.log(err)
        })
    },
    editGrade(index, row) {
      console.log(index)
      row.edit = !row.edit

      console.log(row)
    }
  }
}
</script>

<style>
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}

.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
