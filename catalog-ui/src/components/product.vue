<script>
import {defineComponent, toRefs, defineProps} from "vue";

export default defineComponent({
  props:['product'],
  setup(props) {
    let prod = props.product
    prod.img = "http://localhost:7003/api/v1/catalog/image/" + prod.image
    return {
      product: props.product
    }
  }
})

</script>
<template>
  <a-card hoverable :style="{height: '400', width: '300px'}">
    <template #cover>
      <img width="300" height="300" alt="example" :src="product.img" />
    </template>
    <a-card-meta :title='product.name'>
      <template #description>{{product.description}}</template>
    </a-card-meta>
  </a-card>
</template>
<style scoped>
</style>