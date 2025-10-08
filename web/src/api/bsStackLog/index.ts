import { http, jumpExport } from '@/utils/http/axios';

// 获取库存流水列表
export function List(params) {
  return http.request({
    url: '/bsStackLog/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除库存流水
export function Delete(params) {
  return http.request({
    url: '/bsStackLog/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑库存流水
export function Edit(params) {
  return http.request({
    url: '/bsStackLog/edit',
    method: 'POST',
    params,
  });
}

// 获取库存流水指定详情
export function View(params) {
  return http.request({
    url: '/bsStackLog/view',
    method: 'GET',
    params,
  });
}

// 导出库存流水
export function Export(params) {
  jumpExport('/bsStackLog/export', params);
}