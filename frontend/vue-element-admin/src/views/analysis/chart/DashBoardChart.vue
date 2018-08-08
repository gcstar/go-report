<template>
<el-col :span="span">
    <el-card class="card">
        <div class="title">{{this.title}}
        </div>
        <div class="pull-right">
            <i class="el-icon-refresh"></i>
            <i class="el-icon-edit"></i>
            <i :class="foldUpIcon" @click="foldUp"></i>
            <i class="el-icon-view" @click="view"></i>
        </div>
        <div :class="className">
            <ve-line v-if="chartType=='line'" :data="chartData" :width="width"></ve-line>
            <ve-histogram v-if="chartType=='column'" :data="chartData" :width="width"></ve-histogram>
            <ve-bar v-if="chartType=='bar'" :data="chartData" :width="width"></ve-bar>
            <ve-pie v-if="chartType=='pie'" :data="chartData" :width="width"></ve-pie>
        </div>
    </el-card>
</el-col>
</template>

<script>
export default {
  props: {
    chartType: {
      type: String,
      default: 'line'
    },
    chartData: {
      type: Object,
      default: null
    },
    width: {
      type: String,
      default: 'auto'
    },
    title: {
      type: String,
      default: ''
    },
    span: {
      type: String,
      default: '12'
    }
  },
  data() {
    return {
      className: 'grid-content',
      foldUpIcon: 'el-icon-arrow-up'
    }
  },
  mounted() {},
  beforeDestroy() {},
  methods: {
    view() {
      alert('view')
    },
    foldUp() {
      if (this.className === 'grid-content') {
        this.className = 'grid-content-foldup'
        this.foldUpIcon = 'el-icon-arrow-down'
      } else {
        this.className = 'grid-content'
        this.foldUpIcon = 'el-icon-arrow-up'
      }
    }
  }
}
</script>

<style scoped>
.grid-content {
  border-radius: 12px;
}
.grid-content-foldup {
  border-radius: 12px;
  display: none;
}
.bg-purple-dark {
  background: #99a9bf;
}
.card {
  margin: 10px;
}
.title {
  float: left;
}
.pull-right {
  float: right;
}
i {
  margin-right: 10px;
  cursor: pointer;
}
</style>

