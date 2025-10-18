import { http, jumpExport } from '@/utils/http/axios';

// 获取服务项目列表
export function List(params) {
  return http.request({
    url: '/bsService/list',
    method: 'get',
    params,
  });
}
