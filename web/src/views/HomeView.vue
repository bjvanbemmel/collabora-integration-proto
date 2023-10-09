<template>
    <main
        class="flex flex-col justify-start items-start"
    >
        <button
                class="py-1.5 px-2 border-black border bg-gray-300 m-8 hover:bg-gray-400"
                @click="createNewFile"
                >
                Create new document
        </button>

        <table class="table-fixed my-16 mx-64">
            <thead>
                <tr>
                    <th class="w-96 text-left">ID</th>
                    <th class="w-96 text-left">Name</th>
                    <th class="w-96 text-left">Path</th>
                    <th class="w-96 text-left">Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr
                    v-for="file in files"
                >
                    <td class="truncate w-32">{{file.ID}}</td>
                    <td class="truncate w-32">{{file.Name}}</td>
                    <td class="truncate w-32">{{file.Path}}</td>
                    <td class="truncate w-32 flex gap-8">
                        <RouterLink :to="`/document/${file.ID}`" class="hover:cursor-pointer text-blue-600 underline hover:text-blue-800">
                            Edit
                        </RouterLink>
                        <a
                            @click="deleteFile(file.ID)"
                            class="hover:cursor-pointer text-blue-600 underline hover:text-blue-800"
                            >
                            Delete
                        </a>
                    </td>
                </tr>
            </tbody>
        </table>
    </main>
</template>

<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue'
import axios from 'axios'
import type File from '@/types/file'
import router from '@/router';

const files: Ref<Array<File>> = ref([])

onMounted(() => {
    document.title = 'Collabora stress test'

    axios.get("/files")
        .then((res) => {
            files.value = res.data
        })
})

function createNewFile() {
    axios.post("/files")
        .then((res) => {
            router.push(`/document/${res.data.ID}`)
        })
}

function deleteFile(id: string) {
    axios.delete(`/files/${id}`)
        .then(() => {
            axios.get("/files")
                .then((res) => {
                    files.value = res.data
                })
        })
}
</script>
