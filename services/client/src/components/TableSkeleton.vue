<script lang="ts" setup>
export interface Props {
  rowCount?: number
}

const props = withDefaults(defineProps<Props>(), {
  rowCount: 5,
})
</script>

<template lang="pug">
q-markup-table(style="overflow: initial")
  thead
    tr
      th.text-left(style="width: 150px")
        q-skeleton(
          animation="wave"
          type="text"
        )

      th.text-right(v-for="n in 5" :key="n")
        q-skeleton(
          animation="wave"
          type="text"
        )

    tbody
      tr(v-for="n in props.rowCount" :key="n")
        td.text-left
          q-skeleton(
            animation="wave"
            type="text"
            width="85px"
          )
        td.text-right
          q-skeleton(
            animation="wave"
            type="text"
            width="50px"
          )
        td.text-right
          q-skeleton(
            animation="wave"
            type="text"
            width="35px"
          )
        td.text-right
          q-skeleton(
            animation="wave"
            type="text"
            width="65px"
          )
        td.text-right
          q-skeleton(
            animation="wave"
            type="text"
            width="25px"
          )
        td.text-right
          q-skeleton(
            animation="wave"
            type="text"
            width="85px"
          )
</template>
