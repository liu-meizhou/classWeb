<template>
  <el-table :data="courses" style="width: 100%">
    <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="课程号">
            <span>{{ props.row.id }}</span>
          </el-form-item>
          <el-form-item label="课程名">
            <span>{{ props.row.name }}</span>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column align="center" label="课程名" prop="name" />
    <el-table-column
      align="center"
      label="上课周数"
      prop="baseInfos[0].startWeek"
    />
    <el-table-column
      align="center"
      label="上课时间"
      prop="baseInfos[0].startCount"
    />
    <el-table-column
      align="center"
      label="上课地点"
      prop="baseInfos[0].address"
    />
  </el-table>
</template>

<script>
import { chooseCourseList } from '@/api/user/index'
export default {
  data() {
    return {
      courses: []
    }
  },
  created() {
    this.getCourse()
  },
  methods: {
    getCourse() {
      chooseCourseList()
        .then(courses => {
          this.courses = courses
          console.log(courses)
        })
        .catch(err => {
          console.log(err)
        })
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
</style>
