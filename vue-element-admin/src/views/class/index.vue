<template>
  <div>
    <el-select v-model="classId" placeholder="请选择" @change="getClass">
      <el-option
        v-for="item in classList"
        :key="item.id"
        :label="item.name"
        :value="item.id"
      />
    </el-select>
    <el-button
      type="primary"
      icon="el-icon-plus"
      @click="addClass"
    >添加班级信息</el-button>
    <div v-if="clazz">
      <div>
        <p>
          班号: <el-tag>{{ clazz.id }}</el-tag>
        </p>
        <p>
          班名: <el-tag>{{ clazz.name }}</el-tag>
        </p>
        <p v-if="clazz.teacher">
          班主任: <el-tag>{{ clazz.teacher.name }}</el-tag>
        </p>
        <p v-else>班主任: <el-tag>暂无班主任</el-tag></p>

        <p v-if="clazz.students">
          班级学生数量: <el-tag>{{ clazz.students.length }}</el-tag>
        </p>
        <p v-else>班级学生数量: <el-tag>该班级无学生</el-tag></p>
        <p v-if="clazz.courses">
          班级统一选课数量: <el-tag>{{ clazz.courses.length }}</el-tag>
        </p>
        <p v-else>班级统一选课数量: <el-tag>班级无统一选课</el-tag></p>
      </div>
      <el-button
        type="primary"
        icon="el-icon-edit"
        @click="editClass"
      >修改班级信息</el-button>
      <el-divider />
      <div>
        <span>班级统一选择课程基本信息:</span>
        <el-button
          type="primary"
          icon="el-icon-plus"
          @click="addCourseClass"
        >给该班级统一选课</el-button>
        <el-table
          v-if="clazz.courses"
          :data="clazz.courses"
          style="width: 100%"
        >
          <el-table-column align="center" label="课程号" prop="id" />
          <el-table-column align="center" label="课程名">
            <template slot-scope="{ row }">
              <el-link
                type="primary"
                :href="'/#/course/index?courseId=' + row.id"
                :underline="false"
              >{{ row.name }}</el-link>
            </template>
          </el-table-column>
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button
                size="mini"
                type="danger"
                @click="handleDeleteCourse(scope.$index, scope.row)"
              >取消课程</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <el-divider />
      <div>
        <span>学生基本信息:</span>
        <el-button
          type="primary"
          icon="el-icon-plus"
          @click="addStudentToClass"
        >添加学生到该班级</el-button>
        <el-table
          v-if="clazz.students"
          :data="clazz.students"
          style="width: 100%"
        >
          <el-table-column align="center" label="学号" prop="id" />
          <el-table-column align="center" label="姓名">
            <template slot-scope="{ row }">
              <el-link
                type="primary"
                :href="'/#/student/index?studentId=' + row.id"
                :underline="false"
              >{{ row.name }}</el-link>
            </template>
          </el-table-column>
          <el-table-column align="center" label="性别" prop="sex" />
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button
                size="mini"
                type="danger"
                @click="handleDeleteStudent(scope.$index, scope.row)"
              >移除班级</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <el-divider />
    </div>

    <el-dialog :title="dialogName" :visible.sync="dialogFormVisible">
      <el-form :model="bufClass" label-position="left" label-width="80px">
        <el-form-item label="班级号">
          <el-input v-model="bufClass.id" autocomplete="off" />
        </el-form-item>
        <el-form-item label="班级名">
          <el-input v-model="bufClass.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="班主任">
          <el-select
            v-model="bufClass.teacher"
            value-key="id"
            placeholder="请选择老师"
          >
            <el-option
              v-for="item in teacherList"
              :key="item.id"
              :label="item.name"
              :value="item"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogEditClass">{{
          dialogEditName
        }}</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getClass,
  getClassList,
  editClass,
  createClass
} from '@/api/user/class'
import { getTeachers } from '@/api/user/teacher'
export default {
  data() {
    return {
      classId: null,
      classList: null,
      teacherList: null,
      clazz: null,
      bufClass: {
        id: '',
        name: '',
        teacher: {
          id: '',
          name: ''
        }
      },
      dialogName: '',
      dialogEditName: '添 加',
      dialogFormVisible: false
    }
  },
  created() {
    const query = this.$route.query
    console.log(query)
    if (query && query.classId) {
      this.classId = query.classId
      this.getClass()
    }
    this.getClasses()
  },
  methods: {
    getClasses() {
      getClassList().then(classList => {
        this.classList = classList
      })
    },
    getClass() {
      getClass(this.classId)
        .then(clazz => {
          this.clazz = clazz
          console.log(clazz)
        })
        .catch(err => {
          console.log(err)
        })
    },
    initClass() {
      return {
        id: '',
        name: '',
        teacher: {
          id: '',
          name: ''
        }
      }
    },
    addClass() {
      if (!this.teacherList) {
        this.getTeachers()
      }
      this.bufClass = this.initClass()
      this.dialogEditName = '添 加'
      this.dialogName = '添加班级'
      this.dialogFormVisible = true
    },
    editClass() {
      if (!this.teacherList) {
        this.getTeachers()
      }
      this.bufClass = JSON.parse(JSON.stringify(this.clazz))
      this.dialogEditName = '修 改'
      this.dialogName = '修改班级: ' + this.clazz.name
      this.dialogFormVisible = true
    },
    dialogEditClass() {
      if (this.dialogEditName === '添 加') {
        createClass(this.bufClass)
          .then(res => {
            this.dialogFormVisible = false
            this.clazz = this.bufClass
            this.getClasses()
          })
          .catch(err => {
            console.log(err)
          })
      } else if (this.dialogEditName === '修 改') {
        editClass(this.bufClass)
          .then(res => {
            this.dialogFormVisible = false
            this.clazz = this.bufClass
          })
          .catch(err => {
            console.log(err)
          })
      }
    },
    getTeachers() {
      getTeachers()
        .then(teachers => {
          this.teacherList = teachers
        })
        .catch(err => {
          console.log(err)
        })
    },
    // 给班级统一选课
    addCourseClass() {
      this.$router.push({
        path: '/courseChoose/index',
        params: {
          type: 'AddCourseClass',
          data: this.clazz
        }
      })
    },
    addStudentToClass() {
      this.$message({
        type: 'success',
        message: '到学生查询中...,修改学生信息中的班级即可!'
      })
    },
    handleDeleteCourse(index, row) {
      this.$confirm('即将把课程:' + row.name + '移出班级, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    handleDeleteStudent(index, row) {
      this.$confirm('即将把学生:' + row.name + '移出班级, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
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

.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
