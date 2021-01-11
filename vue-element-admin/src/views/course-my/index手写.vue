<template>
  <div>
    <div>信息介绍</div>
    <div>时间表头</div>
    <div>
      <div>
        <div>上午</div>
        <div>下午</div>
        <div>晚上</div>
      </div>
      <div>
        <div>第一节</div>
        <div>第二节</div>
        <div>第三节</div>
        <div>第四节</div>
        <div>第五节</div>
        <div>第六节</div>
        <div>第七节</div>
        <div>第八节</div>
        <div>第九节</div>
        <div>第十节</div>
        <div>第十一节</div>
      </div>
    </div>
    <div>实践课程</div>
    <div>其它课程</div>
  </div>
</template>

<script>
import { getCourse } from '@/api/user/index'
export default {
  data() {
    return {}
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
    }
  }
}
</script>

<style lang="scss">

</style>
