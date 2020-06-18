<template>
  <v-app>
    <v-app>
      <v-navigation-drawer v-model="drawer" app clipped>
        <v-list dense nav>
          <v-list-item link>
            <v-list-item-action>
              <v-icon>mdi-calendar</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Calendar</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item link>
            <v-list-item-action>
              <v-icon>mdi-file-upload</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Upload</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-app-bar app color="indigo" dark clipped-left>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
        <v-toolbar-title>Upload Scheduler</v-toolbar-title>
      </v-app-bar>

      <v-main>
        <v-container class="fill-height" fluid>
          <v-btn color="primary">Howdy</v-btn>
        </v-container>
      </v-main>
    </v-app>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: null,
    };
  },
};
</script>
