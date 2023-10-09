<template>
    <div
        class="flex flex-col h-full"
    >
        <nav
            class="flex gap-12 bg-black p-8 text-white w-full text-lg"
        >
            <RouterLink
                to="/"
                class="hover:underline"
            >Home</RouterLink>
            <a
                href="https://192.168.50.216:9443/#!/2/docker/containers"
                target="_blank"
                rel="noreferrer noopener"
                class="hover:underline"
            >Portainer</a>
            <a
                @click="createNewFile"
                class="hover:cursor-pointer hover:underline"
            >Create new document</a>
        </nav>
        <RouterView />
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { RouterView } from 'vue-router'
import router from './router';

function createNewFile() {
    axios.post("/files")
        .then((res) => {
            router.push(`/document/${res.data.ID}`)
        })
}
</script>
