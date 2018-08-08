<template>
  <div>
    <el-row>
      <Chart span="12" chartType="line" :chartData="chartData" title="折线图"></Chart>
      <Chart span="12" chartType="pie" :chartData="chartData" title="饼图"></Chart>
    </el-row>
    <el-row>
      <Chart span="12" chartType="column" :chartData="chartData" title="柱状图"></Chart>
      <Chart span="12" chartType="bar" :chartData="chartData" title="bar图"></Chart>
    </el-row>
  </div>

</template>
<script>
import Crosstab from 'vue-crosstab'
import { getSalaryData } from './salary.js'
import Chart from './chart/DashBoardChart'

export default {
  // OR register locally
  components: { Crosstab, Chart },
  data() {
    return {
      title: '折线图',
      chartData: {
        columns: ['日期', '访问用户', '下单用户'],
        rows: [
          { 日期: '2018-05-22', 访问用户: 32371, 下单用户: 19810 },
          { 日期: '2018-05-23', 访问用户: 12328, 下单用户: 4398 },
          { 日期: '2018-05-24', 访问用户: 92381, 下单用户: 52910 },
          { 日期: '2018-05-25', 访问用户: 32371, 下单用户: 19810 },
          { 日期: '2018-05-26', 访问用户: 12328, 下单用户: 4398 },
          { 日期: '2018-05-27', 访问用户: 92381, 下单用户: 52910 },
          { 日期: '2018-05-28', 访问用户: 32371, 下单用户: 19810 },
          { 日期: '2018-05-29', 访问用户: 12328, 下单用户: 4398 },
          { 日期: '2018-05-30', 访问用户: 92381, 下单用户: 52910 }
        ]
      },
      salary: [],
      data: [
        {
          date: '2012/01/01',
          precipitation: 0.0,
          temp_max: 12.8,
          temp_min: 5,
          wind: 4.7,
          weather: 'drizzle'
        },
        {
          date: '2012/01/02',
          precipitation: 10.9,
          temp_max: 10.6,
          temp_min: 2.8,
          wind: 4.5,
          weather: 'rain'
        },
        {
          date: '2012/02/06',
          precipitation: 0,
          temp_max: 16.1,
          temp_min: 1.7,
          wind: 5,
          weather: 'sun'
        }
      ],
      rows: [
        {
          field: 'city',
          label: '城市'
        },
        {
          field: 'age',
          label: '年龄'
        },

        {
          field: 'salary_per_month',
          label: '薪水',
          aggregate: 'min',
          formatter: 'currency'
        }
      ],
      theme: {
        name: 'Light Blue'
      }
    }
  },
  beforeMount() {
    this.salary = getSalaryData()
  }
}
</script>

