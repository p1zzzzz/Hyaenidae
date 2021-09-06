<template>
  <div>
    <iframe src="https://cyberchef.p1zzzzz.work" id="mobsf" scrolling="no" frameborder="0"></iframe>
  </div>
</template>
 
 
<script>
  export default {
    data () {
      return {
      }
    },
    mounted(){
      /**
      * iframe-宽高自适应显示   
      */
      function changeMobsfIframe(){
        const mobsf = document.getElementById('mobsf');
        const deviceWidth = document.body.clientWidth;
        const deviceHeight = document.body.clientHeight;
        mobsf.style.width = (Number(deviceWidth)-270) + 'px'; //数字是页面布局宽度差值
        mobsf.style.height = (Number(deviceHeight)-140) + 'px'; //数字是页面布局高度差
      }
 
      changeMobsfIframe()
 
      window.onresize = function(){
        changeMobsfIframe()
      }
    },     
  }
</script>