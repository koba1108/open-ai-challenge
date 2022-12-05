<template>
  <div>
    <h1>ツイート</h1>
    <form @submit="onSubmit">
      <input type="text" v-model="text"/>
      <button>投稿</button>
    </form>
    <ul>
      <li v-for="t in tweets" :key="t.id">
        [{{ t.userId }}] {{ t.text }}
      </li>
    </ul>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'

interface Tweet {
  id?: number
  userId: number
  text: string
}

export default Vue.extend({
  data() {
    return {
      tweets: [] as Tweet[],
      text: '',
    }
  },
  methods: {
    async onSubmit(e: Event) {
      e.preventDefault()
      const res = await fetch('http://localhost:5001/timeline', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          userId: 1,
          text: this.text
        })
      })
      const data = await res.json() as Tweet
      this.tweets.push(data)
      this.text = ''
    }
  },
  async created() {
    const res = await fetch('http://localhost:5001/timeline')
    this.tweets = await res.json() as Tweet[]
  }
})
</script>
<style>
/*
  ・依頼コメント
  このhtmlにtwitterのようなcssをつけてくれない？
*/
form {
  background-color: #1da1f2;
  border-radius: 5px;
  padding: 10px;
  margin: 10px;
}

input {
  border: none;
  border-radius: 5px;
  background: #fafafa;
  padding: 10px;
  margin: 5px;
  font-size: 16px;
  color: #1da1f2;
}

button {
  background-color: #1da1f2;
  border-radius: 5px;
  padding: 10px;
  font-size: 16px;
  color: #fafafa;
  border: none;
}

ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

li {
  padding: 10px;
  font-size: 16px;
  margin-top: 10px;
  color: #000000;
  border-bottom: 1px solid #1da1f2;
}
</style>
