import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { validate } from '@/utils/validateUtil';
import { TreeOption } from '@/api/bsService';

export class State {

  public id = 0;//主键
  public pid = 0;//父键
  public level = 0;//关系树级别
  public tree = 0;//关系树
  public name = "";//服务名称
  public price = null;//价格
  public orderNum = 0;//序号
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
    message: '请输入服务名称',
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
    label: '服务名称',
    componentProps: {
      placeholder: '请输入服务名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '服务名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '价格',
    key: 'price',
    align: 'left',
    width: -1,
  },
  {
    title: '序号',
    key: 'orderNum',
    align: 'left',
    width: -1,
  },
];

// 关系树选项
export const treeOption = ref([]);

// 加载关系树选项
export function loadTreeOption() {
  TreeOption().then((res) => {
    treeOption.value = res;
  });
}