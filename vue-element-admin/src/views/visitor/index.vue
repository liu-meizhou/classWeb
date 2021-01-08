<template>
  <div class="mixin-components-container">
    <el-row class="row">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span> {{ name }} </span>
          <span> 点击结果: {{ result }} </span>
        </div>
        <div style="margin-bottom: 50px">
          <el-col :span="4" class="text-center">
            <el-button type="primary" @click="get"> Get </el-button>
          </el-col>
          <el-col :span="4" class="text-center">
            <el-button type="success" @click="getVersion"> version </el-button>
          </el-col>
          <el-col :span="4" class="text-center">
            <el-button type="info" @click="show"> show </el-button>
          </el-col>
        </div>
      </el-card>
    </el-row>

    <el-row>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>主题按钮</span>
        </div>
        <div style="margin-bottom: 50px">
          <el-col :span="4" class="text-center">
            <el-button type="primary"> Primary </el-button>
          </el-col>
          <el-col :span="4" class="text-center">
            <el-button type="success"> Success </el-button>
          </el-col>
          <el-col :span="4" class="text-center">
            <el-button type="info"> Info </el-button>
          </el-col>
          <el-col :span="4" class="text-center">
            <el-button type="warning"> Warning </el-button>
          </el-col>
          <el-col :span="4" class="text-center">
            <el-button type="danger"> Danger </el-button>
          </el-col>
        </div>
      </el-card>
    </el-row>
  </div>
</template>

<script>
import { get, version, show } from '@/api/visitor/index'
export default {
  data() {
    return {
      name: '测试访问后端按钮',
      result: ''
    }
  },
  methods: {
    get() {
      const that = this
      get().then(res => {
        that.result = res.data
      })
    },
    getVersion() {
      const that = this
      version().then(res => {
        that.result = res.data
      })
    },
    show() {
      const that = this
      show().then(res => {
        that.result = res.data
      })
    }
  }
}
</script>

<style scoped>
.mixin-components-container {
  background-color: #f0f2f5;
  padding: 30px;
  min-height: calc(100vh - 84px);
}
.row {
  margin-bottom: 50px;
}
.component-item {
  min-height: 100px;
}
</style>
