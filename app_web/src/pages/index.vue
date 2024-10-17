<template>
  <Verify
      ref="verify"
      :mode="'pop'"
      :captcha-type="'blockPuzzle'"
      :img-size="{ width: '330px', height: '155px' }"
  />

  <Verify
      ref="verify1"
      :mode="'pop'"
      :captcha-type="'clickWord'"
      :img-size="{ width: '330px', height: '155px' }"
  />
  <button @click="useVerify">点击调用滑块验证组件</button>

  <button @click="useVerify1">点击调用点文字验证组件</button>

</template>

****注: mode为"pop"时,使用组件需要给组件添加ref值,并且手动调用show方法 例: this.$refs.verify.show()****
****注: mode为"fixed"时,无需添加ref值,无需调用show()方法****

<script>
//引入组件
import Verify from "@/components/verifition/Verify.vue";
export default {
  name: 'app',
  components: {
    Verify
  },
  methods: {
    success(params) {
      // params 返回的二次验证参数, 和登录参数一起回传给登录接口，方便后台进行二次验证
    },
    useVerify() {
      this.$refs.verify.show()
    },
    useVerify1() {
      this.$refs.verify1.show()
    }
  }
}
</script>