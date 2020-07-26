<template>
  <div>
    <ul>
      <li v-for="(file, index) in files" :key="index">{{ file }}</li>
    </ul>
    <h2 v-if="loading">Loading...</h2>
  </div>
</template>

<script>
import axios from "axios"

export default {
  data() {
    return {
      files: [],
      loading: false,
    };
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.loading = true
      let resp = await axios("http://localhost:3001/api/files")
      this.loading = false
      this.files = resp.data
    }
  }
};
</script>
