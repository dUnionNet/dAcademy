<script setup>
import { useRoute } from 'vue-router'
import axios from 'axios'
import { onMounted, ref } from 'vue'

import Nav from '@/views/components/Nav.vue'

const route = useRoute()
const courseSlug = route.params.slug

const loading = ref(true)
const error = ref(null)
const courseDetail = ref([])
const chapters = ref([])

async function fetchCourseDetail() {
  loading.value = true
  try {
    const res = await axios.get(`/api/course/${courseSlug}`)
    courseDetail.value = res.data.course || []
    chapters.value = res.data.chapters || []
  } catch (err) {
    error.value = err.message || 'Failed to load course detail'
  } finally {
    loading.value = false
  }
}

onMounted(fetchCourseDetail)
</script>

<template>
  <Nav />
  <div class="container mx-auto p-6 space-y-6">
    <div v-if="loading" role="alert" class="alert">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info h-6 w-6 shrink-0">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
      </svg>
      <span>Loading...</span>
    </div>
    <div v-else-if="error" role="alert" class="alert alert-error">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <span>{{ error }}</span>
    </div>
    <div v-else>
      <div class="card bg-base-100 shadow-xl p-6 mb-6">
        <h2 class="text-3xl font-bold text-primary">{{ courseDetail.name }}</h2>
        <p class="text-lg text-gray-700 mt-2">{{ courseDetail.description }}</p>
        <p class="text-sm text-gray-600 mt-4">Chapters: {{ courseDetail.chapter_count }}</p>
      </div>

      <ul class="timeline timeline-snap-icon max-md:timeline-compact timeline-vertical">
        <li v-for="(chapter, index) in chapters" :key="index">
          <div class="timeline-middle">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
              class="h-5 w-5"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div :class="['mb-10 space-y-2', (chapter.id % 2) === 0 ? 'timeline-end' : 'timeline-start md:text-end']">
            <time class="font-mono italic">#{{ chapter.id }}</time>
            <div class="text-lg font-black">{{ chapter.title }}</div>
            <RouterLink
              :to="`/c/${courseDetail.slug}/chapter/${chapter.id}`"
              class="btn btn-primary"
            >
              Enter
            </RouterLink>
          </div>
          <hr />
        </li>
      </ul>


    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 1024px;
}

.card {
  border-radius: 12px;
  transition: transform 0.3s ease;
}

.card:hover {
  transform: translateY(-5px);
}
</style>
