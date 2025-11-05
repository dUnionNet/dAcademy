<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

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
  <div class="p-6">
    <h1 class="text-2xl font-bold mb-4">Courses</h1>
    <button @click="scanCourses">Scan Courses</button>

    <div v-if="loading">Loading...</div>
    <div v-else-if="error">Error: {{ error }}</div>

    <div v-else>
      <div
        v-for="(course, index) in courses"
        :key="index"
      >
        <h2>{{ course.name }}</h2>
        <p>{{ course.description }}</p>
        <p>Chapters: {{ course.chapter_count }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
