<template>
  <div>
    <div style="width: 400px; margin: 0 auto;">
      <v-file-input
        label="Select a file"
        v-model="file"
        outlined
        :prepend-icon="null"
        prepend-inner-icon="mdi-paperclip"
      />
      <v-text-field label="File Name" outlined />
      <v-time-picker />
    </div>
    <v-dialog v-model="dialog" max-width="290">
      <v-card>
        <v-card-title class="headline"
          >Use Google's location service?</v-card-title
        >

        <v-card-text>
          You're about to upload <b>{{ file ? file.name : "" }}</b>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="green darken-1" text @click="dialog = false">
            Disagree
          </v-btn>

          <v-btn color="green darken-1" text @click="dialog = false">
            Agree
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      file: null,
      dialog: false,
    };
  },
  watch: {
    file(v) {
      console.log(v);
      this.dialog = true;
    },
  },
};
</script>
