
<template>
  <div class="create" id="create">
    <b-container md="6">
      <b-card>
        <b-form @submit="onSubmit" @reset="onReset">
          <b-row class="justify-content-md-center">
            <b-form-group id="input-group-1" label="Name:" label-for="input-1">
              <b-form-input
                id="input-1"
                v-model="form.name"
                required
                placeholder="Enter recipe name"
              ></b-form-input>
            </b-form-group>
          </b-row>
          <b-row class="justify-content-md-center">
            <b-form-group id="input-group-2" label="Ingredients:" label-for="input-2">
              <b-form-input
                id="input-2"
                v-model="form.ingredients"
                required
                placeholder="Enter ingredients"
              ></b-form-input>
            </b-form-group>
          </b-row>
          <b-row class="justify-content-md-center">
            <b-form-group id="input-group-3" label="Preparation:" label-for="input-3">
              <b-form-input
                id="input-3"
                v-model="form.preparation"
                required
                placeholder="Enter preparation"
              ></b-form-input>
            </b-form-group>
          </b-row>
          <b-row class="justify-content-md-center">
            <b-form-group id="input-group-4" label="Time:" label-for="input-4">
              <b-form-input
                id="input-4"
                type="number"
                min="0"
                default="0"
                v-model="form.time"
                required
                placeholder="Enter time"
              ></b-form-input>
            </b-form-group>
          </b-row>
          <b-row md="4" class="justify-content-md-center">
            <b-col md="1">
              <b-button type="submit" variant="primary">Submit</b-button>
            </b-col>
            <b-col md="1">
              <b-button type="reset" variant="danger">Reset</b-button>
            </b-col>
          </b-row>
        </b-form>
      </b-card>
    </b-container>
  </div>
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

    reset() {
      this.form.name = "";
      this.form.ingredients = "";
      this.form.preparation = "";
      this.form.time = 0;
    },
    createRecipe() {
      let context = this;
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
        data: this.form
      })
        .then(function(response) {
          if (response.statusText == "Created") {
            alert("successful creation");
            context.reset();
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
