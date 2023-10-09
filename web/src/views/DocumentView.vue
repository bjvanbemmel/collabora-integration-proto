<template>
    <LibreOfficeEditor
        class="h-full"
    />
</template>

<script setup lang="ts">
import LibreOfficeEditor from '@/components/LibreOfficeEditor.vue'
import router from '@/router';

import { onMounted } from 'vue';

onMounted(() => {
    document.title = `${router.currentRoute.value.params.id}.docx`
})
</script>
