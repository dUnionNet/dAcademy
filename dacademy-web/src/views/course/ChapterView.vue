<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const courseSlug = route.params.courseSlug
const chapterID = route.params.chapterID

const respond = ref({})
const sections = ref([])
const current = ref(0)
const loading = ref(true)
const finished = ref(false)

onMounted(async () => {
  try {
    const res = await axios.get(`/api/chapter/${courseSlug}/${chapterID}`)
    respond.value = res.data || {}
    sections.value = res.data.sections || []
  } catch (err) {
    console.error('Failed to load chapter:', err)
  } finally {
    loading.value = false
  }
})

function nextSection() {
  if (current.value < sections.value.length - 1) {
    current.value++
  } else {
    finished.value = true
  }
}
</script>

<template>
  <div>
    <h2>📘 Chapter {{ respond.chapter_id }} : {{ respond.chapter_title }}</h2>

    <div v-if="loading">Loading...</div>

    <div v-else>
      <div v-if="!finished">
        <div v-if="sections.length">
          <div v-if="sections[current].type === 'text'">
            <p>{{ sections[current].text }}</p>
          </div>

          <div v-else-if="sections[current].type === 'markdown'">
            <!-- display markdown as HTML -->
            <div v-html="marked.parse(sections[current].text)"></div>
          </div>

          <button @click="nextSection">Next ➡️</button>
        </div>
        <div v-else>
          <p>No sections found.</p>
        </div>
      </div>

      <div v-else>
        <h3>🎉 Chapter Finished!</h3>
        <p>Well done, learner!</p>
        <RouterLink :to="`/c/${courseSlug}`">Back</RouterLink>
      </div>
    </div>
  </div>
</template>

<script>
// markdown support (lightweight)
import { marked } from 'marked'
</script>
