<template>
  <div>
    <p>2020-2021年第1学期</p>
    <p>计科182课表</p>
    <p>专业: 计算机科学与技术</p>
    <el-button type="primary" style="margin-left: 16px;" @click="drawer = true">
      查看时间
    </el-button>
    <el-drawer title="课程具体时间" :visible.sync="drawer" direction="ltr">
      <span>我来啦!</span>
    </el-drawer>
    <el-table
      :data="tableData"
      :span-method="courseTable"
      border
      style="width: 100%"
    >
      <el-table-column prop="id" align="center" label="时间" />
      <el-table-column prop="name" align="center" label="节次" />
      <el-table-column prop="amount1" align="center" label="星期一" />
      <el-table-column prop="amount2" align="center" label="星期二" />
      <el-table-column prop="amount3" align="center" label="星期三" />
      <el-table-column prop="amount1" align="center" label="星期四" />
      <el-table-column prop="amount2" align="center" label="星期五" />
      <el-table-column prop="amount3" align="center" label="星期六" />
      <el-table-column prop="amount1" align="center" label="星期日" />
    </el-table>

    <el-table
      :data="tableData"
      :span-method="arraySpanMethod"
      border
      style="width: 100%"
    >
      <el-table-column prop="id" label="时间" />
      <el-table-column prop="name" label="节次" />
      <el-table-column prop="amount1" label="星期一" />
      <el-table-column prop="amount2" label="星期二" />
      <el-table-column prop="amount3" label="星期三" />
    </el-table>

    <el-table
      :data="tableData"
      :span-method="objectSpanMethod"
      border
      style="width: 100%; margin-top: 20px"
    >
      <el-table-column prop="id" label="ID" width="180" />
      <el-table-column prop="name" label="姓名" />
      <el-table-column prop="amount1" label="数值 1（元）" />
      <el-table-column prop="amount2" label="数值 2（元）" />
      <el-table-column prop="amount3" label="数值 3（元）" />
    </el-table>
  </div>
</template>

<script>
import { getCourse } from '@/api/user/index'
export default {
  data() {
    return {
      drawer: false,
      tableData: [
        {
          id: '12987122',
          name: '王小虎',
          amount1: '234',
          amount2: '3.2',
          amount3: 10
        },
        {
          id: '12987123',
          name: '王小虎',
          amount1: '165',
          amount2: '4.43',
          amount3: 12
        },
        {
          id: '12987124',
          name: '王小虎',
          amount1: '324',
          amount2: '1.9',
          amount3: 9
        },
        {
          id: '12987125',
          name: '王小虎',
          amount1: '621',
          amount2: '2.2',
          amount3: 17
        },
        {
          id: '12987126',
          name: '王小虎',
          amount1: '539',
          amount2: '4.1',
          amount3: 15
        }
      ]
    }
  },
  created() {
    this.getCourse()
  },
  methods: {
    getCourse() {
      getCourse()
        .then(courses => {
          console.log(courses)
        })
        .catch(err => {
          console.log(err)
        })
    },

    courseTable({ row, column, rowIndex, columnIndex }) {
      if (columnIndex === 0) {
        if (rowIndex % 2 === 0) {
          return {
            rowspan: 2,
            colspan: 1
          }
        } else {
          return {
            rowspan: 0,
            colspan: 0
          }
        }
      }
    },

    arraySpanMethod({ row, column, rowIndex, columnIndex }) {
      if (rowIndex % 2 === 0) {
        if (columnIndex === 0) {
          return [1, 2]
        } else if (columnIndex === 1) {
          return [0, 0]
        }
      }
    },

    objectSpanMethod({ row, column, rowIndex, columnIndex }) {
      if (columnIndex === 0) {
        if (rowIndex % 2 === 0) {
          return {
            rowspan: 2,
            colspan: 1
          }
        } else {
          return {
            rowspan: 0,
            colspan: 0
          }
        }
      }
    }
  }
}
</script>
