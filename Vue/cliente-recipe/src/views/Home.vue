<template>
  <div class="home">
    <img style="max-width:10rem;" alt="Vue logo" src="../assets/recipe.png">
    <HelloWorld/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "home",
  components: {
    HelloWorld
  }
};
</script>
