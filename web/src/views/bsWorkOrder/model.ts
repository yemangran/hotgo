import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defShortcuts } from '@/utils/dateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';
import { CascaderOption, NButton, NCascader, NInput, NInputNumber, NTreeSelect } from 'naive-ui';

const dict = useDictStore();

export class State {
  public id = 0; //主键
  public userId = 0; //用户Id
  public createTime = ''; //创建时间
  public customerName = ''; //客户名称
  public customerAddress = ''; //客户地址
  public customerContact = ''; //客户联系方式
  public customerPerson = ''; //客户对接人
  public problemDescription = ''; //问题描述
  public defectDescription = ''; //检测描述
  public dispatchId = 0; //分配处理人Id
  public sendType = 0; //发起类型
  public acceptType = 0; //收取方式
  public productType = 0; //产品类型
  public status = 1; //工单状态
  public suggestSolution = ''; //建议处理方案(多选)
  public actualSolution = ''; //实际处理方案(多选)
  public totalMoney = 0; //总金额
  public invoiceType = 1; //开票类型
  public remark = ''; //备注
  public dispatchName = ''; //
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
  defectDescription: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入检测描述',
  },
  customerName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入客户名称',
  },
  customerAddress: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入客户地址',
  },
  customerContact: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入客户联系方式',
  },
  sendType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入发起类型',
  },
  acceptType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入收取方式',
  },
  productType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入产品类型',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入工单状态',
  },
  invoiceType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入开票类型',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'createTime',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetime',
      clearable: true,
      shortcuts: defShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'customerName',
    component: 'NInput',
    label: '客户名称',
    componentProps: {
      placeholder: '请输入客户名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'customerAddress',
    component: 'NInput',
    label: '客户地址',
    componentProps: {
      placeholder: '请输入客户地址',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'customerContact',
    component: 'NInput',
    label: '客户联系方式',
    componentProps: {
      placeholder: '请输入客户联系方式',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'customerPerson',
    component: 'NInput',
    label: '客户对接人',
    componentProps: {
      placeholder: '请输入客户对接人',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'sendType',
    component: 'NSelect',
    label: '发起类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择发起类型',
      options: dict.getOption('send_type'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'productType',
    component: 'NSelect',
    label: '产品类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择产品类型',
      options: dict.getOption('product_type'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '工单状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择工单状态',
      options: dict.getOption('work_status'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'invoiceType',
    component: 'NSelect',
    label: '开票类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择开票类型',
      options: dict.getOption('invoice_type'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);
export interface RowData {
  id: number;
  key: number;
  serviceId: number;
  handleType: string;
  price: number;
  remark: string;
  status: string;
}
// interface OnUpdateValue {
//   (value: string | number | undefined): void;
// }
// const ShowOrEdit = defineComponent({
//   props: {
//     value: [String, Number],
//     onUpdateValue: [Function, Array] as PropType<OnUpdateValue>,
//     placeholder: String,
//     data: Array as PropType<RowData[]>,
//     component: {
//       type: Object as PropType<Component>,
//       default: () => NInput,
//     },
//   },
//   setup(props) {
//     const isEdit = ref(false);
//     const insRef = ref<InputInst | null>(null);
//     const inputValue = ref(props.value);
//     function handleChange() {
//       props.onUpdateValue?.(inputValue.value);
//       isEdit.value = false;
//     }
//     return () =>
//       h(
//         'div',
//         {
//           style: 'min-height: 22px',
//         },
//         h(props.component as Component, {
//           ref: insRef,
//           value: inputValue.value,
//           onUpdateValue: (v: string | number | undefined) => {
//             inputValue.value = v;
//           },
//           onBlur: handleChange,
//           placeholder: props.placeholder,
//         })
//       );
//   },
// });
export function createProcessColumns(
  onDelete: (row: RowData, index: number) => void,
  serviceOptions: any[]
) {
  return [
    {
      title: '服务项目',
      key: 'name',
      align: 'center',
      width: 200,
      render(row: RowData) {
        // const index = getDataIndex(row.key);
        return h(NCascader, {
          value: row.serviceId,
          onUpdateValue(v: string | number | undefined, option: CascaderOption) {
            row.serviceId = v as number;
            row.price = option.price as number;
            row.status = '' as string;
          },
          status: row.status,
          options: serviceOptions,
          labelField: 'name',
          valueField: 'id',
          showPath: false,
          checkStrategy: 'child',
          placeholder: '选择服务项目',
        });
      },
    },
    {
      title: '处理方式',
      key: 'handleType',
      align: 'center',
      width: -1,
      render(row: RowData) {
        // const index = getDataIndex(row.key);
        return h(NInput, {
          value: row.handleType,
          onUpdateValue(v: string | number | undefined) {
            row.handleType = v as string;
          },
          placeholder: '请输入处理方式',
        });
      },
    },
    {
      title: '金额',
      key: 'price',
      align: 'center',
      width: -1,
      render(row: RowData) {
        console.log('row', row);
        // const index = getDataIndex(row.key);
        return h(NInputNumber, {
          min: 0,
          value: row.price,
          precision: 2,
          onUpdateValue(v: string | number | undefined) {
            row.price = v as number;
          },
        });
      },
    },
    {
      title: '备注',
      key: 'remark',
      align: 'center',
      width: -1,
      render(row: RowData) {
        // const index = getDataIndex(row.key);
        return h(NInput, {
          value: row.remark,
          onUpdateValue(v: string | number | undefined) {
            row.remark = v as string;
          },
        });
      },
    },
    {
      title: '操作',
      key: 'action',
      align: 'center',
      width: -1,
      render(row: RowData, index: number) {
        return h(
          NButton,
          {
            tertiary: true,
            type: 'error',
            onClick: () => onDelete(row, index),
          },
          () => '删除'
        );
      },
    },
  ];
}

// 表格列
export const columns = [
  {
    title: '工单号',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '客户名称',
    key: 'customerName',
    align: 'left',
    width: -1,
  },
  {
    title: '产品类型',
    key: 'productType',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('product_type', row.productType);
    },
  },
  {
    title: '工单状态',
    key: 'status',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('work_status', row.status);
    },
  },
  {
    title: '问题描述',
    key: 'problemDescription',
    align: 'left',
    width: -1,
  },
  {
    title: '对接人',
    key: 'customerPerson',
    align: 'left',
    width: -1,
  },
  {
    title: '发起类型',
    key: 'sendType',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('send_type', row.sendType);
    },
  },
  {
    title: '收取方式',
    key: 'acceptType',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('accept_type', row.acceptType);
    },
  },
  {
    title: '总金额',
    key: 'totalMoney',
    align: 'left',
    width: -1,
  },
  {
    title: '开票类型',
    key: 'invoiceType',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('invoice_type', row.invoiceType);
    },
  },
  {
    title: '创建时间',
    key: 'createTime',
    align: 'left',
    width: 160,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions([
    'accept_type',
    'work_status',
    'invoice_type',
    'send_type',
    'product_type',
    'biz_solution',
  ]);
}
