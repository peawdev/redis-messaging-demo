<template>
  <div>
    <h2>Redis Chat Messages</h2>
    <ul>
      <li v-for="(msg, index) in messages" :key="index">{{ msg }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      messages: [],
      ws: null,
    };
  },
  mounted() {
    this.ws = new WebSocket("ws://localhost:3000/ws");

    this.ws.onmessage = (event) => {
      this.messages.push(event.data);
    };

    this.ws.onopen = () => {
      console.log("Connected to WebSocket");
    };

    this.ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };
  },
  beforeUnmount() {
    if (this.ws) this.ws.close();
  },
};
</script>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
