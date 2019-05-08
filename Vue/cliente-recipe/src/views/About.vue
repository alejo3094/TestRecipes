
<template>
  <div class="listSearch" id="listSearch"></div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      hola: "hola munndo!!",
      time: 0,
      form: {
        name: "",
        ingredients: "",
        preparation: "",
        time: 0
      }
    };
  },

  methods: {
    onSubmit(evt) {
      evt.preventDefault();
      //alert(JSON.stringify(this.form))
      this.createRecipe();
    },
    onReset(evt) {
      evt.preventDefault();
      this.form.name = "";
      this.form.ingredients = "";
      this.form.preparation = "";
      this.form.time = 0;
    },
    createRecipe() {
      let objectTemp = this.form;
      this.form.name = "";
      this.form.ingredients = "";
      this.form.preparation = "";
      this.form.time = 0;

      console.log("iniciandooo");
      axios({
        headers: {
          "Access-Control-Allow-Headers":
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token",
          "Access-Control-Allow-Methods": "POST, GET, PUT, DELETE",
          "Access-Control-Allow-Origin": "*",
          "Content-Type": "application/json"
        },
        url: "http://localhost:8090/api/recipe/",
        method: "POST",
        data: objectTemp
      })
        .then(function(response) {
          if (response.statusText == "Created") {
            alert("successful creation");
          } else {
            alert("error creation");
          }
          console.log(response);
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.margin {
  margin: 5px;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

