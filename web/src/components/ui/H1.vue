<script setup lang="ts">
import { cn } from './shadcn';
import { Primitive, type PrimitiveProps } from 'radix-vue';

interface Props extends PrimitiveProps {
  class?: string
}

const props = withDefaults(defineProps<Props>(), {
  as: 'h1'
})

</script>

<template>
  <Primitive :class="cn('text-5xl font-bold tracking-tighter', props.class)">
    <slot/>
  </Primitive>
</template>