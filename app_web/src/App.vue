
<template>
  <transition name="slide-left" mode="out-in" >
    <div >
      <router-view class="router-view"  />

    </div>
  </transition>
</template>

<style scoped>
.slide-right-enter-active,.slide-right-leave-active,.slide-left-enter-active,.slide-left-leave-active {
  will-change: transform;
  transition: all .5s;
}
.slide-right-enter {
  opacity: 0;
  transform: translate(-100%);
}
.slide-right-leave-active {
  position: absolute;

  opacity: 0;
  transform: translateX(100%);
}
.slide-left-enter {
  opacity: 0;
  transform: translateX(100%);
}
.slide-left-leave-active {
  opacity: 0;
  transform: translateX(-100%);
}
</style>
<script setup>
</script>