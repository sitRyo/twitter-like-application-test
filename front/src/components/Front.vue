<template>
  <b-container>
    <div class="header">
      <div class="pb-2 header-message">twitter like app</div>
      <b-form-textarea v-model="tweet" class="tweet-area"></b-form-textarea>
      <b-row class="pt-2">
        <b-button id="get-button" class="ml-4 bg-info">show</b-button>
        <b-button id="post-button" class="ml-2 bg-info" @click="postTweet">post</b-button>
      </b-row>
      <hr>
    </div>
    <!--show tweet histroies here-->
    <div id="tweet-history"></div>
  </b-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import axios from 'axios';

@Component
export default class Front extends Vue {
  public tweet: string = '';

  public postTweet() {
    axios.post('/api/upload', {
      body: this.tweet,
    })
    .then(() => {
      console.log(this.tweet);
    })
    .catch((err) => {
      alert('server internal error');
      console.log(err);
    });
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>

