<template>
  <div class="tab-container">
    <h3>报表id:{{reportId}}</h3>
  </div>
</template>

<script>
import {
  findReports
} from '@/api/report'

export default {
  name: 'reportPreview',
  data() {
    return {
      reportId: -1
    }
  },
  beforeMount() {
    this.reportId = this.$route.params.id
    const param = { fieldName: 'id', keyWord: this.reportId }
    findReports(param).then(res => {
      if (res.data.data.length === 0) {
        this.$router.push('/404')
      }
    })
  },
  methods: {
  }
}
</script>

<style scoped>
  .tab-container{
    margin: 30px;
  }
</style>
