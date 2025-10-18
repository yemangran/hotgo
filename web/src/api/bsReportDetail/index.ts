import { http, jumpExport } from '@/utils/http/axios';

// 获取服务项目列表
export function List(params) {
  return http.request({
    url: '/bsReportDetail/list',
    method: 'get',
    params,
  });
}

// 删除服务项目
export function Delete(params) {
  return http.request({
    url: '/bsReportDetail/delete',
    method: 'POST',
    params,
  });
}
