<script>
import {defineComponent, ref} from 'vue';
import router from "@/router";
import {useRoute} from "vue-router";
import axios from 'axios'
import Product from "@/components/product.vue";


export default defineComponent({
  components: {Product},
  setup() {
    const route = useRoute()
    const cats = ref([]);

    let selCat = route.query['category']
    let selBrand = route.query['brand']
    let query = route.query['query']
    let pfrom = route.query['price_from']
    let pto = route.query['price_to']
    let dis = true

    if (!pfrom) {
      pfrom = 0
      dis = false
    }
    if (!pto) {
      pto = 100000
      dis = false
    }


    let catVal = ref(undefined)
    let brandVal = ref(undefined)
    const brands = ref([])
    let search = ref(query)
    let queried = ref(false)
    let rowsVal = ref([])
    let rng = ref([pfrom, pto])
    let enabled = ref(dis)


    axios.get('/api/api/v1/catalog/category/list').then((res) => {
      cats.value = res.data.map(a => {
        return {
          value: a.ID,
          label: a.name
        }
      })
      catVal.value = res.data.filter(a => a.ID == selCat)[0].ID
    })
    axios.get('/api/api/v1/catalog/brand/list').then((res) => {
      brands.value = res.data.map(a => {
        return {
          value: a,
        }
      })
      brandVal.value = selBrand
    })
    let req = {
      query: query,
      brand: selBrand,
      cat_id: Number(selCat)
    }
    if (enabled.value) {
      req.price_from = Number(pfrom)
      req.price_to = Number(pto)
    }
    axios.post("/api/api/v1/catalog/product/list", req).then((res) => {
      let count = 4
      let rowsL = 0
      let rows = []
      if (res.data) {
        rowsL = Math.ceil(res.data.length / count)
      }
      for (let i = 0; i < rowsL; i++) {
        rows.push([])
      }
      if (res.data) {
        res.data.forEach((value, index) => {
          let rowI = Math.floor(index / count)
          console.log(rowI)
          rows[rowI].push(value)
        })
      }
      console.log(rows)
      rowsVal.value = rows
      queried.value = true
    })
    const direct = (query) => {
      if (enabled.value) {
        router.push({path: "/", query: { 'category': catVal.value, 'query': query, 'price_from': rng.value[0], 'price_to': rng.value[1]}}).then(() => {
          router.go(0)
        })
      } else {
        router.push({path: "/", query: { 'category': catVal.value, 'query': query}}).then(() => {
          router.go(0)
        })
      }
    }
    const handleChange = (value) => {
      direct('')
    };
    const onSearch = (value) => {
      direct(value)
    }
    const onAfterChange = (value) => {
      direct('')
    }
    return {
      catVal,
      handleChange,
      cats,
      brands,
      brandVal,
      search,
      onSearch,
      queried,
      rowsVal,
      rng,
      enabled,
      onAfterChange
    };
  },
});
</script>

<template>
  <a-layout :style="{minWidth: '100%', minHeight: '100%'}">
    <a-layout-header>
      <a-space align="end">
        <a-typography-title class="ssss" keyboard type="success">ИЩЕМ ИЩЕМ</a-typography-title>
        <a-input-search :style="{alignSelf: 'flex-end'}"
                        v-model:value="search"
                        placeholder="ПИШИ И ИЩИ"
                        style="width: 200px"
                        @search="onSearch"
        />
      </a-space>
    </a-layout-header>
    <a-space>
      <a-space direction="vertical">
        <h4>Категории</h4>
        <a-select
            allowClear
            v-model:value="catVal"
            show-search
            placeholder="Выберите категорию"
            style="width: 200px"
            :options="cats"
            @change="handleChange"
        ></a-select>
      </a-space>
<!--      <a-space direction="vertical">-->
<!--        <h4>Бренды</h4>-->
<!--        <a-select-->
<!--            allowClear-->
<!--            v-model:value="brandVal"-->
<!--            show-search-->
<!--            placeholder="Выберите категорию"-->
<!--            style="width: 200px"-->
<!--            :options="brands"-->
<!--            @change="handleChange"-->
<!--        ></a-select>-->
<!--      </a-space>-->
      <a-space direction="vertical" :style="{width: '200px', backgroundColor: 'white', marginTop: '10%', color: 'black'}">
        Цена
      <a-slider @afterChange="onAfterChange" v-model:value="rng" :min="0" :max="100000" :step="100" range :disabled="!enabled" :style="{width: '100%'}" />
      <a-switch v-model:checked="enabled" size="small" @change="onAfterChange" />
      </a-space>
    </a-space>
    <a-spin v-if="!rowsVal && !queried" size="large"/>
    <a-alert
        v-if="(!rowsVal || rowsVal.length == 0) && queried"
        message="Товары не найдены"
        description="Попробуйте другой поисковый запрос"
        type="warning"
        show-icon
    />
    <a-layout-content  :style="{ marginTop: '32px', marginLeft: '10%'}">
      <a-row v-for="row in rowsVal">
        <a-col v-for="pr in row">
          <Product :product="pr"></Product>
        </a-col>
      </a-row>
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
