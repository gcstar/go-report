<template>
  <div class="app-container">
    <el-col :span="8">
      <el-input placeholder="输入关键字进行过滤" v-model="filterText">
      </el-input>
      <br> <br>
      <el-tree class="filter-tree" @node-click="handleNodeClick" :data="widgetData" :props="defaultProps" :filter-node-method="filterNode" ref="widgetTree">
        <span slot-scope="{ node, data }">
          <div v-if="data.leaf == false">
            <svg-icon icon-class="chart_file"></svg-icon>{{ node.label}}
          </div>
          <div v-if="data.leaf == true">
            <svg-icon icon-class="chart_doc"></svg-icon>{{ node.label}}
          </div>
        </span>
      </el-tree>
    </el-col>
    <el-col :span="15" :offset="1">
      <el-card>
        <!-- <svg-icon  icon-class="chart_line" /> -->
        <el-form ref="widgetForm" :model="widgetForm" label-width="120px">
          <el-form-item label="图表名称:">
            <el-input v-model="widgetForm.name" placeholder=""></el-input>
          </el-form-item>
          <el-form-item label="图表类型:">
            <el-radio-group v-model="widgetForm.chartType" size="medium" @change="handleChartTypeChange">
              <el-radio class="radio" v-for="(t,index) in chartTypes" :key="index" :label="t">
                <svg-icon :icon-class="t" />
              </el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item v-for="(option,index) in widgetForm.options" :key="option.key" :prop="'options.'+ index + '.value'" :label="option.label">
            <el-input v-if="option.type == 'input'" v-model="option.value" placeholder=""></el-input>
            <draggable v-if="option.type == 'draggable'" v-model="option.value">
              <transition-group>
                <div style="display:inline;" v-for="(r,i) in option.value" :key="i">
                  <el-button type="primary" plain>
                    {{r}}
                  </el-button>
                </div>
              </transition-group>
            </draggable>
          </el-form-item>

        </el-form>

      </el-card>
    </el-col>

  </div>
</template>
<script>
// import Crosstab from 'vue-crosstab'
import treeTable from '@/components/TreeTable'
import draggable from 'vuedraggable'
import * as report from '@/api/report'
const chartTypes = [
  'chart_table',
  'chart_column',
  'chart_line',
  'chart_pie',
  'chart_kpi',
  'chart_bubble',
  'chart_funnel',
  'chart_radar',
  // "chart_boxplot",
  'chart_sankey'
]

const chartTypeConfig = {
  chart_table: { columns: [], rows: [], filters: [], values: [] },
  chart_column: {},
  chart_line: {},
  chart_pie: {},
  chart_kpi: {},
  chart_bubble: {},
  chart_funnel: {},
  chart_radar: {},
  chart_sankey: {}
}

export default {
  components: { treeTable, draggable },
  watch: {
    filterText(val) {
      this.$refs.widgetTree.filter(val)
    }
  },

  data() {
    return {
      chartTypeConfig: chartTypeConfig,
      chartTypes: chartTypes,
      filterText: '',
      widgetData: [],
      defaultProps: {
        children: 'children',
        label: 'label',
        isLeaf: 'leaf'
      },
      widgetOptions: {},
      widgetForm: {
        options: [],
        name: '',
        chartType: ''
      },
      myArray: [
        { name: 1, id: 1 },
        { name: 2, id: 2 },
        { name: 3, id: 3 },
        { name: 4, id: 1 }
      ]
    }
  },
  methods: {
    handleChartTypeChange(value) {
      this.widgetForm.chartType = value
    },
    handleNodeClick(data, node, tree) {
      if (data.leaf) {
        this.widgetForm = {
          options: [],
          name: '',
          chartType: ''
        }
        report
          .findWidget({ fieldName: 'widget_name', keyWord: data.label })
          .then(res => {
            const widget = res.data.data
            this.widgetOptions.name =
              widget.categoryName + '/' + widget.widgetName
            const dataJson = JSON.parse(widget.dataJson)
            const config = dataJson.config
            this.widgetForm.name = this.widgetOptions.name
            if (dataJson.datasetId != null) {
              report.getDataset(dataJson.datasetId).then(res => {
                const dataset = res.data.data
                this.widgetForm.options.push({
                  key: 'datasetName',
                  value: dataset.categoryName + '/' + dataset.datasetName,
                  label: '数据集名称:',
                  type: 'input'
                })
                const chart_type = config.chart_type
                const datassetJson = JSON.parse(dataset.dataJson)
                this.widgetForm.chartType = 'chart_' + chart_type
                this.generateChartOptions(
                  chart_type,
                  config,
                  datassetJson.query.sql
                )
              })
            }
          })
      }
    },
    generateChartOptions(chartType, config, sql) {
      var cfg = {
        rows: [],
        columns: [],
        filters: [],
        values: []
      }
      switch (chartType) {
        case 'table':
          var rows = []
          var columns = []
          var filters = []
          var values = []
          config.keys.forEach(i => {
            if (i.alias !== null && i.alias !== '') {
              rows.push(i.col + '-->' + i.alias)
            } else {
              rows.push(i.col)
            }
            cfg.rows.push({
              columnName: i.col,
              filterType: i.type,
              values: i.values
            })
          })
          config.groups.forEach(i => {
            if (i.alias !== null && i.alias !== '') {
              columns.push(i.col + '-->' + i.alias)
            } else {
              columns.push(i.col)
            }
            cfg.columns.push({
              columnName: i.col,
              filterType: i.type,
              values: i.values
            })
          })
          config.filters.forEach(i => {
            if (i.alias !== null && i.alias !== '') {
              filters.push(i.col + '-->' + i.alias)
            } else {
              filters.push(i.col)
            }
            cfg.filters.push({
              columnName: i.col,
              filterType: i.type,
              values: i.values
            })
          })
          config.values[0].cols.forEach(i => {
            if (i.alias !== null && i.alias !== '') {
              values.push(i.col + '-->' + i.alias)
            } else {
              values.push(i.col)
            }
            cfg.values.push({
              columnName: i.col,
              aggType: i.aggregate_type,
              values: i.values
            })
          })
          var col = {
            key: 'columns',
            value: columns,
            label: '列维:',
            type: 'draggable'
          }
          var row = {
            key: 'rows',
            value: rows,
            label: '行维:',
            type: 'draggable'
          }
          var filter = {
            key: 'filters',
            value: filters,
            label: '过滤:',
            type: 'draggable'
          }
          var value = {
            key: 'values',
            value: values,
            label: '指标:',
            type: 'draggable'
          }
          this.widgetForm.options.push(col)
          this.widgetForm.options.push(row)
          this.widgetForm.options.push(filter)
          this.widgetForm.options.push(value)

          break
      }

      var params = {
        cfg: JSON.stringify(cfg),
        sql: sql
      }

      console.log(params)
      report.getAggSql(params).then(res => {
        const sql = res.data.data
        alert(sql)
      })
    },

    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    }
  },
  mounted() {
    report.getCategoryTree().then(res => {
      this.widgetData = res.data.data
    })
  },
  beforeMount() {}
}
</script>
