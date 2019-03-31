<template>
  <div class="wrapper">
    <div class="tweet-wrapper">
      <div class="header">つぶやき</div>
      <div class="body">
        <textarea v-model="tweet" class="text-area"></textarea>
        <button class="tweet-button">ツイートする</button>
      </div>
    </div>
    <!-- 5件ほど表示させる。Apiを自動で叩く？ -->
    <div class="tweet-histroy"></div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class Front extends Vue {
  tweet: string = '';
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.tweet-wrapper {
  float: left;
}

</style>

