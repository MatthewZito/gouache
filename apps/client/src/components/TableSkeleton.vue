<script lang="ts" setup>
const props = defineProps({
  rowCount: {
    type: Number,
    default: 5,
  },
})
</script>

<template>
  <q-markup-table style="overflow: initial">
    <thead>
      <tr>
        <th class="text-left" style="width: 150px">
          <q-skeleton animation="wave" type="text" />
        </th>
        <th v-for="n in 5" :key="n" class="text-right">
          <q-skeleton animation="wave" type="text" />
        </th>
      </tr>
    </thead>

    <tbody>
      <tr v-for="n in props.rowCount" :key="n">
        <td class="text-left">
          <q-skeleton animation="wave" type="text" width="85px" />
        </td>
        <td class="text-right">
          <q-skeleton animation="wave" type="text" width="50px" />
        </td>
        <td class="text-right">
          <q-skeleton animation="wave" type="text" width="35px" />
        </td>
        <td class="text-right">
          <q-skeleton animation="wave" type="text" width="65px" />
        </td>
        <td class="text-right">
          <q-skeleton animation="wave" type="text" width="25px" />
        </td>
        <td class="text-right">
          <q-skeleton animation="wave" type="text" width="85px" />
        </td>
      </tr>
    </tbody>
  </q-markup-table>
</template>
