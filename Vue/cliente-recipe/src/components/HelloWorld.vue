<template>
  <div class="hello" id="main">
    <b-container>
      <b-row class="justify-content-md-center">
        <b-col md="3">
          <b-card title="Create" tag="create">
            <b-link :to="{ path: 'create'}">
              <img style="max-width:10rem;" alt="search" src="../assets/create.png">
            </b-link>
          </b-card>
        </b-col>
        <b-col md="3">
          <b-card title="List, update delete" tag="list">
            <b-link :to="{ path: 'list'}">
              <img style="max-width:10rem;" alt="search" src="../assets/listSearch.png">
            </b-link>
          </b-card>
        </b-col>
      </b-row>
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
      this.form.time = "";
    },

    createRecipe() {
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
        data: this.form
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
    },

    getByid() {
      axios({
        url: "http://localhost:8090/api/recipe/4",
        method: "GET"
      })
        .then(function(response) {
          console.log(response);
        })
        .catch(function(error) {
          console.log(error);
        });
    },

    getByName() {
      axios({
        url: "http://localhost:8090/api/recipe/cho",
        method: "GET"
      })
        .then(function(response) {
          console.log(response);
        })
        .catch(function(error) {
          console.log(error);
        });
    },

    deleteRecipe() {
      axios({
        headers: {
          "Access-Control-Allow-Headers":
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token",
          "Access-Control-Allow-Methods": "POST, GET, PUT, DELETE",
          "Access-Control-Allow-Origin": "*",
          "Content-Type": "application/json"
        },
        url: "http://localhost:8090/api/recipe/2",
        method: "DELETE"
      })
        .then(function(response) {
          console.log(response);
        })
        .catch(function(error) {
          console.log(error);
        });
    },
    getAllRecips() {
      axios({
        url: "http://localhost:8090/api/recipe/",
        method: "GET"
      })
        .then(function(response) {
          console.log(response);
        })
        .catch(function(error) {
          console.log(error);
        });
    },

    updateRecipe() {
      console.log("Yeahhh");
      axios({
        headers: {
          "Access-Control-Allow-Headers":
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token",
          "Access-Control-Allow-Methods": "POST, GET, PUT, DELETE",
          "Access-Control-Allow-Origin": "*",
          "Content-Type": "application/json"
        },
        url: "http://localhost:8090/api/recipe/4",
        method: "PUT",
        data: {
          name: "Alejo",
          ingredients: "Azucar Flores y muchos colores",
          preparation: "-----",
          time: "8 meses"
        }
      })
        .then(function(response) {
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
#main {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

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
