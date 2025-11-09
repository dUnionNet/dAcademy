<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import Nav from '@/views/components/Nav.vue'

const courses = ref([])
const loading = ref(true)
const error = ref(null)

async function fetchCourses() {
  loading.value = true
  try {
    const res = await axios.get('/api/course/list')
    courses.value = res.data.courses || []
  } catch (err) {
    error.value = err.message || 'Failed to load courses'
  } finally {
    loading.value = false
  }
}

async function scanCourses() {
  try {
    const res = await axios.get('/api/course/scan')
    alert(res.data.message || 'Scan complete.')
    // Refetch courses after scanning
    await fetchCourses()
  } catch (err) {
    alert('Scan failed: ' + (err.message || 'Unknown error'))
  }
}

onMounted(fetchCourses)
</script>

<template>
  <Nav />
  <div class="container mx-auto p-6 space-y-6">
    <h1 class="text-3xl font-bold text-center text-primary mb-6">Courses</h1>
    <button
      @click="scanCourses"
      class="btn btn-accent"
    >
      Scan Courses
    </button>

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

    <div v-else class="space-y-4">
      <div
        v-for="(course, index) in courses"
        :key="index"
        class="card bg-base-100 shadow-xl p-4 space-y-4"
      >
        <h2 class="text-xl font-semibold text-primary">{{ course.name }}</h2>
        <p class="text-gray-700">{{ course.description }}</p>
        <p class="text-sm text-gray-600">Chapters: {{ course.chapter_count }}</p>
        <RouterLink
          :to="`/c/${course.slug}`"
          class="btn btn-primary w-full"
        >
          Start Course
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 800px;
}
.card {
  border-radius: 12px;
  transition: transform 0.3s ease;
}
.card:hover {
  transform: translateY(-5px);
}
</style>
