<template>
  <div class="app-container">
    <el-col :span="6">
      <el-input placeholder="输入关键字进行过滤" v-model="filterText">
      </el-input>
      <br> <br>
      <el-tree class="filter-tree" @node-click="handleNodeClick" :data="data2" :props="defaultProps" :filter-node-method="filterNode" ref="tree2">
        <span slot-scope="{ node, data }">
        <i v-if="data.leaf == false" class="el-icon-menu">{{ node.label}}</i>
        <i v-if="data.leaf == true" class="el-icon-document">{{ node.label}}</i>
        </span>
      </el-tree>
    </el-col>
    <el-col :span="16" :offset="1">
      <el-table :data="tableData">
        <el-table-column label="name" prop="name"></el-table-column>
        <el-table-column label="age" prop="age"></el-table-column>
      </el-table>
    </el-col>
  </div>
</template>
<script>
// import Crosstab from 'vue-crosstab'
import treeTable from '@/components/TreeTable'
import * as report from '@/api/report'

export default {
  // OR register locally
  components: { treeTable },
  watch: {
    filterText(val) {
      this.$refs.tree2.filter(val)
    }
  },

  data() {
    return {
      filterText: '',
      tableData: [
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 },
        { name: 'zhangsan', age: 12 }
      ],
      data2: [],
      defaultProps: {
        children: 'children',
        label: 'label',
        isLeaf: 'leaf'
      }
    }
  },
  methods: {
    handleNodeClick(data, node, tree) {
      if (data.leaf) {
        report.findWidget({ fieldName: 'widget_name', keyWord: data.label }).then(res => {
          console.log(res.data.data[0])
        })
      }
    },
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    }
  },
  mounted() {
    report.getCategoryTree().then(res => {
      this.data2 = res.data.data
    })
  },
  beforeMount() {}
}
</script>
