<template>
  <div class="tab-container">
    <el-card>
      <el-form :inline="true" :model="reportForm" class="demo-form-inline">
        <el-form-item label="名称:">
          <el-input v-model="reportForm.name" placeholder="名称"></el-input>
        </el-form-item>
        <el-form-item label="数据源:">
          <el-select v-model="reportForm.dsId">
            <el-option v-for="item in datasources" :key="item.id" :value="item.id" :label="item.name"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="启用:">
          <el-select v-model="reportForm.status">
            <el-option label="启用" value="1"></el-option>
            <el-option label="禁用" value="0"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="显示几天数据:">
          <el-input v-model="reportForm.dataRange"></el-input>
        </el-form-item>
        <el-form-item label="sql语句:">
          <codemirror ref="myEditor" :options="editorOptions" v-model="reportForm.sqlText"></codemirror>
          <br>
          <el-button style="text-align:center;" type="primary" @click="execSql">执行sql</el-button>
        </el-form-item>
        <el-form-item label="元数据列配置:">
          <br>
          <el-table :data="reportForm.metaColumns" border stripe>
            <el-table-column label="列名" prop="name" width="200"></el-table-column>
            <el-table-column label="标题" width="200">
              <template slot-scope="scope">
                <el-input v-model="scope.row.text" placeholder=""></el-input>
              </template>
            </el-table-column>
            <el-table-column label="类型" width="140">
              <template slot-scope="scope">
                <el-select v-model="scope.row.type" placeholder="">
                  <el-option v-for="option in colmunType" :value="option.value" :key="option.value" :label="option.label"></el-option>
                </el-select>
              </template>
            </el-table-column>
            <el-table-column label="数据类型" prop="dataType"></el-table-column>
            <el-table-column label="宽度" prop="width"></el-table-column>
            <el-table-column label="精度" prop="decimals"></el-table-column>
          </el-table>
        </el-form-item>

        <!-- <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item> -->
      </el-form>
    </el-card>
  </div>

</template>

<script>
import * as report from '@/api/report'
import { codemirror } from 'vue-codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/dracula.css'

export default {
  name: 'reportEdit',
  components: {
    codemirror
  },
  data() {
    return {
      reportId: -1,
      reportForm: {
        name: '',
        status: 0,
        dataRange: '',
        categoryName: '',
        dsId: '',
        metaColumns: []

      },
      datasources: [],
      editorOptions: {
        tabSize: 4,
        mode: 'text/x-mysql',
        theme: 'dracula',
        lineNumbers: true,
        line: true,
        foldGutter: true,
        gutters: ['CodeMirror-linenumbers', 'CodeMirror-foldgutter'],
        styleSelectedText: true,
        highlightSelectionMatches: { showToken: /\w/, annotateScrollbar: true },
        extraKeys: { Ctrl: 'autocomplete' }
      },
      colmunType: [
        { label: '布局列', value: '1' },
        { label: '维度列', value: '2' },
        { label: '统计列', value: '3' },
        { label: '计算列', value: '4' }
      ]
    }
  },
  beforeCreate() {
    report.getAllDatasource().then(res => {
      const { data } = res.data
      this.datasources = data
    })
  },
  mounted() {
    this.reportId = this.$route.params.id
    const param = { fieldName: 'uid', keyWord: this.reportId }
    report.findReports(param).then(res => {
      const reports = res.data.data
      if (reports.length === 0) {
        this.$router.push('/404')
      }
      const report = reports[0]
      this.reportForm.name = report.name
      this.reportForm.categoryName = report.categoryName
      this.reportForm.status = report.status.toString()
      const options = JSON.parse(report.options)
      // const queryParam = JSON.parse(report.queryParams)
      const metaColumns = JSON.parse(report.metaColumns)
      this.reportForm.dataRange = options.dataRange
      this.reportForm.dsId = report.dsId
      this.reportForm.sqlText = report.sqlText
      this.reportForm.metaColumns = metaColumns
      console.log(this.reportForm.metaColumns)
    })
  },
  methods: {
    execSql() {
      alert('exec sql')
    }
  }
}
</script>

<style scoped>
.tab-container {
  margin: 30px;
}
</style>
