<script setup>
import { useRoute } from 'vue-router'
import axios from "axios";
import {onMounted, ref} from "vue";

const route = useRoute()
const courseSlug = route.params.slug

const loading = ref(true)
const error = ref(null)
const courseDetail = ref([])
const chapters = ref([])

async function fetchCourseDetail() {
  loading.value = true
  try {
    const res = await axios.get('/api/course/'+courseSlug)
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
  <h2>{{ courseDetail.name }}</h2>
  <p>{{ courseDetail.description }}</p>
  <p>Chapters: {{ courseDetail.chapter_count }}</p>
  <hr>

  <div v-if="loading">Loading...</div>
  <div v-else-if="error">Error: {{ error }}</div>

  <div v-else>
    <div
      v-for="(chapter, index) in chapters"
      :key="index"
    >
      <h2>{{ chapter.title }}</h2>
      <p>{{ chapter.id }}</p>

    </div>
  </div>
</template>
