<template>
  <div>
    <div style="width: 400px; margin: 0 auto;">
      <v-file-input
        label="Select a file"
        v-model="file"
        outlined
        :prepend-icon="null"
        prepend-inner-icon="mdi-paperclip"
        show-size
      />
      <v-text-field label="File Name" outlined v-model="name" />
      <v-btn @click="submit" color="primary" :disabled="progress.active"
        ><v-icon left>mdi-upload</v-icon> Upload</v-btn
      >
      <v-progress-linear :value="progress.value" :active="progress.active" />
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      file: null,
      name: "",
      progress: {
        active: false,
        value: 0,
      },
    };
  },
  watch: {
    file(v) {
      console.log(v);
      this.dialog = true;
    },
  },
  methods: {
    async submit() {
      let data = new FormData();
      data.append("name", this.name);
      data.append("file", this.file);

      this.progress.active = true;
      let resp = await axios("http://localhost:3001/api/upload", {
        method: "POST",
        onUploadProgress: (e) => {
          console.log(e);
          this.progress.value = (e.loaded / e.total) * 100;
        },
        data,
      });

      this.progress.value = 100;
    },
  },
};
</script>
