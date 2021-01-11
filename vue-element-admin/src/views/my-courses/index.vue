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
          <el-form-item label="成绩">
            <span>{{ props.row.studentResult }}</span>
          </el-form-item>
          <el-form-item label="绩点">
            <span>{{ props.row.studentPoint }}</span>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column align="center" label="课程名">
      <template slot-scope="scope">
        <el-link
          type="primary"
          :href="'/#/student/index?courseId=' + scope.row.id"
          :underline="false"
        >{{ scope.row.name }}</el-link>
      </template>
    </el-table-column>
    <el-table-column align="center" label="上课周数">
      <template slot-scope="{ row }">
        <p v-for="baseInfo in row.baseInfos" :key="baseInfo.id">
          <el-tag>{{ baseInfo.startWeek }}-{{ baseInfo.endWeek }}周</el-tag>
        </p>
      </template>
    </el-table-column>

    <el-table-column align="center" label="上课节数">
      <template slot-scope="{ row }">
        <p v-for="baseInfo in row.baseInfos" :key="baseInfo.id">
          <el-tag>{{ baseInfo.startCount }}-{{ baseInfo.endCount }}节</el-tag>
        </p>
      </template>
    </el-table-column>

    <el-table-column align="center" label="上课地点">
      <template slot-scope="{ row }">
        <p v-for="baseInfo in row.baseInfos" :key="baseInfo.id">
          {{ baseInfo.address }}
        </p>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { getCourse } from '@/api/user/index'
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
      getCourse()
        .then(courses => {
          this.courses = courses
          console.log(courses)
        })
        .catch(err => {
          console.log(err)
        })
    },
    lookStudent(index) {
      console.log(index)
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
