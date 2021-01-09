<template>
  <el-table :data="students" style="width: 100%">
    <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="学生学号">
            <span>{{ props.row.id }}</span>
          </el-form-item>
          <el-form-item label="姓名">
            <span>{{ props.row.name }}</span>
          </el-form-item>
          <el-form-item label="成绩">
            <span>{{ props.row.courseResult }}</span>
          </el-form-item>
          <el-form-item label="绩点">
            <span>{{ props.row.coursePoint }}</span>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column align="center" label="姓名" prop="name" />
    <el-table-column align="center" label="班级" prop="class.name" />
    <el-table-column align="center" label="成绩">
      <template slot-scope="{ row }">
        <template v-if="row.edit">
          <el-input v-model="row.courseResult" class="edit-input" size="small" />
          <el-button
            class="cancel-btn"
            size="small"
            icon="el-icon-refresh"
            type="warning"
            @click="cancelEdit(row)"
          >
            cancel
          </el-button>
        </template>
        <el-button v-else type="text" @click="row.edit=!row.edit">{{
          row.courseResult
        }}</el-button>
      </template>
    </el-table-column>
    <el-table-column align="center" label="绩点" prop="coursePoint" />
  </el-table>
</template>

<script>
import { getStudent } from '@/api/user/index'
export default {
  data() {
    return {
      students: []
    }
  },
  created() {
    const query = this.$route.query
    if (query) {
      this.getStudent(query.courseId)
    }
  },
  methods: {
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
