import { http, jumpExport } from '@/utils/http/axios';

// 获取仓库物品列表
export function List(params) {
  return http.request({
    url: '/bsItem/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除仓库物品
export function Delete(params) {
  return http.request({
    url: '/bsItem/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑仓库物品
export function Edit(params) {
  return http.request({
    url: '/bsItem/edit',
    method: 'POST',
    params,
  });
}

// 获取仓库物品指定详情
export function View(params) {
  return http.request({
    url: '/bsItem/view',
    method: 'GET',
    params,
  });
}

// 导出仓库物品
export function Export(params) {
  jumpExport('/bsItem/export', params);
}