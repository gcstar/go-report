<template>
  <div>
    <el-row style="margin:30px">
      <el-col :span="6">
        <el-select v-model="category" placeholder="报表目录" @change="handleCategoryChange">
          <el-option v-for="item in categoryList" :key="item.id" :label="item.name" :value="item.id">
          </el-option>
        </el-select>
      </el-col>
      <el-col :span="10">
        <el-input placeholder="请输入内容" v-model="input5">
          <el-select v-model="select" slot="prepend" placeholder="请选择查询项">
            <el-option label="报表名" value="1"></el-option>
            <el-option label="创建者" value="2"></el-option>
          </el-select>
        </el-input>
      </el-col>
      <el-col :span="6" :offset="1">
        <el-button type="primary" icon="el-icon-search">查找</el-button>
      </el-col>
    </el-row>
    <el-table :data="reportList" stripe border>
      <el-table-column v-for="item in reportColumns" :key="item.id" :prop="item.prop" :label="item.text" v-if="item.visiable">
      </el-table-column>
      <el-table-column label="操作" width="300">
        <template slot-scope="scope">
          <el-button-group>
            <el-button size="mini" type="primary" icon="el-icon-edit" @click="alert(scope.row.id)">编辑</el-button>
            <el-button size="mini" type="danger" icon="el-icon-delete" @click="removeReport(scope.row.id)">删除</el-button>
            <el-button size="mini" type="warning" icon="el-icon-view" @click="alert(scope.row.id)">预览</el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>

  </div>
</template>

<script>
import {
  listAllReports,
  listAllCategory,
  listReportsByCategory
} from '@/api/report'

export default {
  name: 'reportCategory',
  components: {},
  data() {
    return {
      select: '',
      input5: '',
      category: '',
      reportList: [],
      categoryList: [],
      reportColumns: [
        { prop: 'id', text: '报表id', visiable: true },
        { prop: 'categoryId', text: 'categoryId', visiable: false },
        { prop: 'categoryName', text: 'categoryName', visiable: false },
        { prop: 'name', text: '报表名称', visiable: true },
        { prop: 'sqlText', text: 'sql', visiable: false },
        { prop: 'createUser', text: '创建者', visiable: true },
        { prop: 'gmtCreated', text: '创建时间', visiable: true },
        { prop: 'gmtModified', text: '修改时间', visiable: true }
      ]
    }
  },
  methods: {
    handleCategoryChange() {
      listReportsByCategory(this.category).then(res => {
        this.reportList = res.data.data
      })
    },
    alert(id) {},
    removeReport(id) {},
    fetchReportList() {
      listAllReports().then(res => {
        this.reportList = res.data.data
      })
    },
    fetchCategoryList() {
      listAllCategory().then(res => {
        this.categoryList = res.data.data
      })
    }
  },
  mounted() {},
  beforeMount() {
    this.fetchReportList()
    this.fetchCategoryList()
  }
}
</script>

<style>
.el-select .el-input {
  width: 130px;
}
.input-with-select .el-input-group__prepend {
  background-color: #fff;
}
</style>

