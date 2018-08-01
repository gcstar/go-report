<template>
  <div>
    <el-card>
      <el-row>
        <el-col v-for="item in queryColumn" :key="item.name" :span="6">
          <el-date-picker v-if="item.formElement=='date'" value-format="yyyy-MM-dd" v-model="item.defaultValue" :placeholder="item.text"></el-date-picker>
          <el-input v-if="item.formElement == 'string'" v-model="item.defaultValue" :placeholder="item.text"></el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="searchReportData">查 询</el-button>
        </el-col>
      </el-row>
    </el-card>
    <el-card>
      <el-table :data="tableData" border stripe>
        <el-table-column v-for="item in tableColumn" :key="item.name" :prop="item.name" :label="item.text"></el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
import * as report from '@/api/report'

export default {
  name: 'reportPreview',
  data() {
    return {
      uid: -1,
      report: {},
      queryColumn: [],
      queryForm: {},
      tableColumn: [],
      tableData: []
    }
  },
  beforeMount() {
    this.uid = this.$route.params.id
    const param = { fieldName: 'uid', keyWord: this.uid }
    report.findReports(param).then(res => {
      if (res.data.data.length === 0) {
        this.$router.push('/404')
      }
      this.report = res.data.data[0]
      report.getQueryColumn(this.report.uid).then(res => {
        if (res.data.data != null) {
          this.queryColumn = res.data.data
        }
      })
    })
  },
  methods: {
    searchReportData() {
      report.getReportTableData(this.report.uid, this.queryColumn).then(res => {
        const { data, metadata } = res.data
        this.tableColumn = metadata
        this.tableData = data
      })
    }
  }
}
</script>

<style scoped>
.tab-container {
  margin: 30px;
}
</style>
