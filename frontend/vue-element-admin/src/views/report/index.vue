<template>
  <div>
    <el-row style="margin:30px">
      <el-col :span="7">
        <el-select clearable filterable v-model="category" placeholder="报表目录" @change="handleCategoryChange">
          <el-option v-for="item in categoryList" :key="item.id" :label="item.name" :value="item.id">
          </el-option>
        </el-select>
      </el-col>
      <el-col :span="10">
        <el-input placeholder="请输入内容" v-model="keyWord">
          <el-select v-model="fieldName" slot="prepend" placeholder="请选择查询项">
            <el-option label="报表名" value="name"></el-option>
            <el-option label="创建者" value="create_user"></el-option>
          </el-select>
        </el-input>
      </el-col>
      <el-col :span="6" :offset="1">
        <el-button type="primary" icon="el-icon-search" @click="handleSearchReports">查 找</el-button>
      </el-col>
    </el-row>
    <el-table style="margin:30px" :data="reportList.slice((pagination.page-1)*pagination.size,pagination.page*pagination.size)" stripe border>
      <el-table-column v-for="item in reportColumns" :key="item.id" :prop="item.prop" :label="item.text" v-if="item.visiable">
      </el-table-column>
      <el-table-column label="操作" width="300">
        <template slot-scope="scope">
          <el-button-group>
            <el-button size="mini" type="primary" icon="el-icon-edit" @click="editReport(scope.row.uid)">编辑</el-button>
            <el-button size="mini" type="danger" icon="el-icon-delete" @click="removeReport(scope.row.id)">删除</el-button>
            <el-button size="mini" type="warning" icon="el-icon-view" @click="previewReport(scope.row.uid)">查看</el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination">
      <el-pagination :page-sizes="[10, 20,50]" :page-size="pagination.size" background layout="sizes,total,prev, pager, next" :total="reportList.length" @size-change="handleSizeChange" @current-change="handlePageChange">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import * as report from '@/api/report'

export default {
  name: 'reportCategory',
  components: {},
  data() {
    return {
      pagination: {
        page: 1,
        size: 10
      },
      fieldName: 'name',
      keyWord: '',
      category: '',
      reportList: [],
      categoryList: [],
      reportColumns: [
        { prop: 'id', text: '报表id', visiable: true },
        { prop: 'uid', text: '报表id', visiable: false },
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
    handleSizeChange(size) {
      this.pagination.page = 1
      this.pagination.size = size
    },
    handlePageChange(page) {
      this.pagination.page = page
    },
    handleSearchReports() {
      if (this.fieldName === '') {
        this.fetchReportList()
      } else {
        const param = { fieldName: this.fieldName, keyWord: this.keyWord }
        report.findReports(param).then(res => {
          this.reportList = res.data.data
          this.keyWord = ''
          this.fieldName = ''
        })
      }
    },
    handleCategoryChange() {
      report.listReportsByCategory(this.category).then(res => {
        this.reportList = res.data.data
      })
    },
    previewReport(id) {
      this.$router.push({ name: 'preview', params: { id: id }})
    },
    editReport(id) {
      this.$router.push({
        name: 'reportedit',
        params: { id: id }
      })
    },
    removeReport(id) {},
    fetchReportList() {
      report.listAllReports().then(res => {
        this.reportList = res.data.data
      })
    },
    fetchCategoryList() {
      report.listAllCategory().then(res => {
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
.pagination {
  text-align: center;
  margin-top: 30px;
}
</style>

