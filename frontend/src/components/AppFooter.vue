<template>
  <CFooter>
    <div>
      <a href="https://github.com/Fishwaldo/mouthpiece" target="_blank">MouthPiece</a>
      <span class="ms-1"
        >&copy; {{ new Date().getFullYear() }} Justin Hammond.</span
      >
    </div>
    <div class="ms-auto">
      <span class="me-1" target="_blank">Powered by</span>
      <a href="https://coreui.io/vue">CoreUI for Vue</a>
    </div>
  </CFooter>
</template>

<script>
export default {
  name: 'AppFooter',
}
</script>
