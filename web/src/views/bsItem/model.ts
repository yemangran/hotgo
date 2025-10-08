import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { validate } from '@/utils/validateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {

  public id = 0;//主键
  public name = "";//物品名称
  public category = null;//物品类别
  public spec = "";//规格型号
  public price = null;//单价
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
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入物品名称',
  },
  category: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入物品类别',
  },
  price: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'number',
    validator: validate.amount,
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '物品名称',
    componentProps: {
      placeholder: '请输入物品名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'category',
    component: 'NSelect',
    label: '物品类别',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择物品类别',
      options: dict.getOption('item_category'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'spec',
    component: 'NInput',
    label: '规格型号',
    componentProps: {
      placeholder: '请输入规格型号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '物品名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '物品类别',
    key: 'category',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('item_category', row.category);
    },
  },
  {
    title: '规格型号',
    key: 'spec',
    align: 'left',
    width: -1,
  },
  {
    title: '单价',
    key: 'price',
    align: 'left',
    width: -1,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['item_category']);
}