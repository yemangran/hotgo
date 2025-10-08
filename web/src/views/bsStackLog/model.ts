import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';

export class State {

  public id = 0;//主键
  public itemId = null;//物品Id
  public qty = 0;//数量
  public workOrderId = null;//工单Id
  public createTime = "";//创建时间
  public remark = "";//备注
  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则
export const rules = {
  itemId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入物品Id',
  },
  qty: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入数量',
  },
  workOrderId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入工单Id',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'bsItemName',
    component: 'NInput',
    label: '物品名称',
    componentProps: {
      placeholder: '请输入物品名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '数量',
    key: 'qty',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createTime',
    align: 'left',
    width: -1,
  },
  {
    title: '物品名称',
    key: 'bsItemName',
    align: 'left',
    width: -1,
  },
  {
    title: '物品类别',
    key: 'bsItemCategory',
    align: 'left',
    width: -1,
  },
  {
    title: '规格型号',
    key: 'bsItemSpec',
    align: 'left',
    width: -1,
  },
  {
    title: '工单',
    key: 'bsWorkOrderId',
    align: 'left',
    width: -1,
  },
];