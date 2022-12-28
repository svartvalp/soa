<script>
import {defineComponent, ref} from 'vue';
import router from "@/router";
import {useRoute} from "vue-router";

export default defineComponent({
  setup() {
    const route = useRoute()
    const cats = ref([
      {value: '1', label: 'Электроника'},
      {value: '2', label: 'Алкоголь'},
      {value: '3', label: 'Игрушки'},
    ]);
    const brands = ref([
      {value: '1', label: 'OZON'},
      {value: '2', label: 'Apple'},
      {value: '3', label: 'ODIN'},
      {value: '4', label: 'DVA'},
    ])
    const handleChange = (value) => {
      router.push({path:"/", query: {'brand': brandVal.value, 'category': catVal.value}})
      console.log(`selected ${value}`);
    };
    let catVal = ref(route.query['category'])
    let brandVal = ref(route.query['brand'])
    return {
      catVal,
      handleChange,
      cats,
      brands,
      brandVal,
    };
  },
});
</script>

<template>
  <a-layout :style="{width: '100%', height: '100%'}">
    <a-layout-header>
      <a-typography-title class="ssss" keyboard type="success">ИЩЕМ ИЩЕМ</a-typography-title>
    </a-layout-header>
    <a-layout-content :style="{ marginTop: '64px' }">
      <a-space>
        <a-space direction="vertical">
        <h4>Категории</h4>
        <a-select
            v-model:value="catVal"
            show-search
            placeholder="Выберите категорию"
            style="width: 200px"
            :options="cats "
            @change="handleChange"
        ></a-select>
        </a-space>
          <a-space direction="vertical">
        <h4>Бренды</h4>
        <a-select
            v-model:value="brandVal"
            show-search
            placeholder="Выберите категорию"
            style="width: 200px"
            :options="brands"
            @change="handleChange"
        ></a-select>
          </a-space>
      </a-space>
        <div :style="{ background: '#fff', padding: '24px', minHeight: '380px' }">Content</div>
    </a-layout-content>
    <a-layout-footer :style="{ textAlign: 'center' }">
    </a-layout-footer>
  </a-layout>
</template>

<style scoped>
.ssss {
  margin-top: 10px;
}
</style>
