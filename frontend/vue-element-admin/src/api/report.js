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
