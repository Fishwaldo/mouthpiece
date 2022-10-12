<template>
  <CFooter>
    <div>
      <a href="https://github.com/Fishwaldo/mouthpiece" target="_blank">MouthPiece</a>
      <span class="ms-1"
        >&copy; {{ new Date().getFullYear() }} Justin Hammond.</span
      >
    </div>
    <div class="ms-auto">
      <span class="me-1">Local Cache Size {{size}}</span>
    </div>
  </CFooter>
</template>

<script>
export default {
  name: 'AppFooter',
  data () {
    return {
      size: this.$store.cache.state().size
    }
  },
  methods: {
    updateSize () {
      this.$store.cache.state().forEach(this.checkCache)
      this.size = this.$store.cache.state().size
    },
    checkCache(item, key) {
      if (item.expiresIn < Date.now()) {
        const tokens = key.split(':')
        this.$store.cache.delete(tokens[0], tokens[1])
      }
    }
  },
  mounted () {
    setInterval(this.updateSize, 2500);
  }
}
</script>
