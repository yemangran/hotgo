import { http, jumpExport } from '@/utils/http/axios';

// 获取服务项目列表
export function List(params) {
  return http.request({
    url: '/bsService/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除服务项目
export function Delete(params) {
  return http.request({
    url: '/bsService/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑服务项目
export function Edit(params) {
  return http.request({
    url: '/bsService/edit',
    method: 'POST',
    params,
  });
}

// 获取服务项目指定详情
export function View(params) {
  return http.request({
    url: '/bsService/view',
    method: 'GET',
    params,
  });
}

// 获取服务项目关系树选项
export function TreeOption() {
  return http.request({
    url: '/bsService/treeOption',
    method: 'GET',
  });
}