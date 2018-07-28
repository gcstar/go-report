<template>
  <el-card>
    <el-row>
      <el-col v-for="item in queryColumn" :span="6">
        <el-date-picker v-if="item.formElement=='date'" v-model="item.name" :placeholder="item.text"></el-date-picker>
        <el-input v-if="item.formElement == 'string'" v-model="item.name" :placeholder="item.text"></el-input>
      </el-col>
      <el-col :span="4">
        <el-button type="primary">查 询</el-button>
      </el-col>
    </el-row>
  </el-card>

</template>

<script>
import { findReports, getQueryColumn } from '@/api/report'

export default {
  name: 'reportPreview',
  data() {
    return {
      uid: -1,
      report: {},
      queryColumn: [],
      queryForm: {}
    }
  },
  beforeMount() {
    this.uid = this.$route.params.id
    const param = { fieldName: 'uid', keyWord: this.uid }
    findReports(param).then(res => {
      if (res.data.data.length === 0) {
        this.$router.push('/404')
      }
      this.report = res.data.data[0]
      getQueryColumn(this.report.uid).then(res => {
        if (res.data.data != null) {
          this.queryColumn = res.data.data
        }
      })
    })
  },
  methods: {}
}
</script>

<style scoped>
.tab-container {
  margin: 30px;
}
</style>
