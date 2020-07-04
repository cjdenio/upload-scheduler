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
      <v-btn @click="submit" color="primary"
        ><v-icon left>mdi-upload</v-icon> Upload</v-btn
      >
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
      dialog: false,
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

      let resp = await axios.post("http://localhost:3001/api/upload", data, {});

      alert("success");
      this.$router.push("/");
    },
  },
};
</script>
