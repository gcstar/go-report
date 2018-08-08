<template>
  <div>
    <el-card>
      <div slot="header" class="clearfix">
        <span>报表:{{report.name}}</span>
      </div>
      <el-row>
        <el-col v-for="item in queryColumn" :key="item.name" :span="6">
          <el-date-picker v-if="item.formElement=='date'" value-format="yyyy-MM-dd" v-model="item.defaultValue" :placeholder="item.text"></el-date-picker>
          <el-input v-if="item.formElement == 'string'" v-model="item.defaultValue" :placeholder="item.text"></el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="searchReportData">查 询</el-button>
        </el-col>
        <el-col :span="4" :offset="2">
          <el-button style='margin:0 0 20px 20px;' type="primary" @click="handleDownload" :loading="downloadLoading">{{$t('excel.export')}} excel</el-button>

        </el-col>
      </el-row>
    </el-card>
    <el-card>
      <el-table :data="reportData" border stripe>
        <el-table-column v-for="item in tableColumn" :key="item.name" :prop="item.name" :label="item.text"></el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination :page-sizes="[10, 20,50]" :page-size="pagination.size" background layout="sizes,total,prev, pager, next" :total="tableData.length" @size-change="handleSizeChange" @current-change="handlePageChange">
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
import * as report from '@/api/report'

export default {
  name: 'reportPreview',
  data() {
    return {
      pagination: {
        page: 1,
        size: 10
      },
      uid: -1,
      report: {},
      queryColumn: [],
      queryForm: {},
      tableColumn: [],
      tableData: [],
      downloadLoading: false
    }
  },
  computed: {
    reportData: function() {
      if (this.tableData.length === 0) {
        return []
      } else {
        return this.tableData.slice((this.pagination.page - 1) * this.pagination.size,
          this.pagination.page * this.pagination.size)
      }
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
        if (res.data.data !== null) {
          this.queryColumn = res.data.data
          if (this.queryColumn.length !== 0) {
            this.searchReportData()
          }
        }
      })
    })
  },
  methods: {
    handleDownload() {
      if (this.tableData.length === 0) {
        // $.message()
      }
      this.downloadLoading = true
      import('@/vendor/Export2Excel').then(excel => {
        const tHeader = []
        this.tableColumn.forEach(e => {
          tHeader.push(e.name)
        })
        const data = []
        this.tableData.forEach(row => {
          const r = []
          tHeader.forEach(header => {
            for (const key in row) {
              if (key === header) {
                r.push(row[key])
              }
            }
          })
          data.push(r)
        })
        excel.export_json_to_excel({
          header: tHeader,
          data,
          filename: 'excel-list',
          autoWidth: true
        })
        this.downloadLoading = false
      })
    },
    handleSizeChange(size) {
      this.pagination.page = 1
      this.pagination.size = size
    },
    handlePageChange(page) {
      this.pagination.page = page
    },
    searchReportData() {
      report.getReportTableData(this.report.uid, this.queryColumn).then(res => {
        const { data, metadata } = res.data
        this.tableColumn = metadata
        if (data !== null) {
          this.tableData = data
        }
      })
    }
  }
}
</script>

<style scoped>
.tab-container {
  margin: 30px;
}

.pagination {
  text-align: center;
  margin-top: 30px;
}
</style>
