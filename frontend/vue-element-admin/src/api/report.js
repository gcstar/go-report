import request from '@/utils/request'
import Qs from 'qs'

export function listAllReports() {
  return request({
    url: '/metadata/report/listAll',
    method: 'get'
  })
}

export function listAllCategory() {
  return request({
    url: '/metadata/categories',
    method: 'get'
  })
}

export function listReportsByCategory(id) {
  return request({
    url: '/metadata/report/list?categoryId=' + id,
    method: 'get'
  })
}

export function findReports(param) {
  return request({
    url: '/metadata/report/findReport?' + Qs.stringify(param),
    method: 'get'
  })
}

export function getQueryColumn(uid) {
  return request({
    url: '/metadata/report/getQueryColumn?uid=' + uid,
    method: 'get'
  })
}

export function getReportTableData(uid, params) {
  return request({
    url: '/metadata/report/table/getData.json?uid=' + uid,
    method: 'post',
    data: params
  })
}

export function getAllDatasource() {
  return request({
    url: '/metadata/datasource/list',
    method: 'get'
  })
}

export function getCategoryTree() {
  return request({
    url: '/metadata/categorytree',
    method: 'get'
  })
}
export function findWidget(params) {
  return request({
    url: '/metadata/widget/findWidget?' + Qs.stringify(params),
    method: 'get'
  })
}

export function getDataset(id) {
  return request({
    url: '/metadata/dataset?id=' + id,
    method: 'get'
  })
}

export function getAggSql(params) {
  return request({
    url: '/metadata/dataprovider/queryAggSql',
    method: 'post',
    data: params
  })
}
