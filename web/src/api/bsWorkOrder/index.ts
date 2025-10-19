import { http, jumpExport } from '@/utils/http/axios';

// 获取工单管理列表
export function List(params) {
  return http.request({
    url: '/bsWorkOrder/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除工单管理
export function Delete(params) {
  return http.request({
    url: '/bsWorkOrder/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑工单管理
export function Edit(params) {
  return http.request({
    url: '/bsWorkOrder/edit',
    method: 'POST',
    params,
  });
}

// 修改工单管理状态
export function Status(params) {
  return http.request({
    url: '/bsWorkOrder/status',
    method: 'POST',
    params,
  });
}

// 获取工单管理指定详情
export function View(params) {
  return http.request({
    url: '/bsWorkOrder/view',
    method: 'GET',
    params,
  });
}

// 导出工单管理
export function Export(params) {
  jumpExport('/bsWorkOrder/export', params);
}

// 处理工单
export function Process(params) {
  return http.request({
    url: '/bsWorkOrder/process',
    method: 'POST',
    params,
  });
}

// 生成工单报告
export function GenerateReport(params) {
  return http.request({
    url: '/bsWorkOrder/generateReport',
    method: 'GET',
    params,
  });
}
