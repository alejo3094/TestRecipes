
<template>
  <div class="list margin" id="list">
    <b-input-group class="margin" >
      <b-form-select style="max-width:7rem;" v-model="selected" :options="options"></b-form-select>

      <b-form-input v-model="value" v-if="selected == 'ID'" type="number" min="1" default="1"></b-form-input>
      <b-form-input v-model="value" v-if="selected == 'Name'"></b-form-input>

      <b-input-group-append>
        <b-button @click="search" variant="outline-secondary">&#x1f50d;</b-button>
        <b-button @click="getAllRecips" variant="outline-secondary">List All</b-button>
      </b-input-group-append>
    </b-input-group>
    <b-row md="2" class="justify-content-md-center">
           
      <b-col v-for="i in numPages" :key="i" md="2" @click="getRecipePage(i)">
        <span v-if="actualPage==i" style="color: blue;"><u><b>{{i}}</b></u></span>
        <span v-if="actualPage!=i" >{{i}}</span>
      </b-col>
    </b-row>

    <b-table :items="recipes" :fields="fields" striped class="margin">
      <template slot="edit" slot-scope="row">
        <b-button
          variant="primary"
          size="sm"
          @click="row.toggleDetails"
          class="mr-2"
        >{{ row.detailsShowing ? 'Hide' : 'Update'}}</b-button>
      </template>
      <template slot="delete" slot-scope="row">
        <b-button
          variant="danger"
          size="sm"
          class="mr-2"
          @click="confirmDelete(row.item.id, row.item.name)"
        >{{ row.detailsShowing ? 'delete' : 'Delete'}}</b-button>
      </template>
      <template slot="row-details" slot-scope="row">
        <b-card>
          <b-container md="6">
            <b-card>
              <b-form ref="form">
                <b-row class="justify-content-md-center">
                  <b-form-group id="input-group-1" label="Name:" label-for="input-1">
                    <b-form-input
                      id="input-1"
                      v-model="row.item.name"
                      required
                      placeholder="Enter recipe name"
                    ></b-form-input>
                  </b-form-group>
                </b-row>
                <b-row class="justify-content-md-center">
                  <b-form-group id="input-group-2" label="Ingredients:" label-for="input-2">
                    <b-form-input
                      id="input-2"
                      v-model="row.item.ingredients"
                      required
                      placeholder="Enter ingredients"
                    ></b-form-input>
                  </b-form-group>
                </b-row>
                <b-row class="justify-content-md-center">
                  <b-form-group id="input-group-3" label="Preparation:" label-for="input-3">
                    <b-form-input
                      id="input-3"
                      v-model="row.item.preparation"
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
                      v-model="row.item.time"
                      required
                      placeholder="Enter time"
                    ></b-form-input>
                  </b-form-group>
                </b-row>
                <b-row md="4" class="justify-content-md-center">
                  <b-col md="2">
                    <b-button
                      variant="danger"
                      @click="updateRecipe(row.item.id,row.item.name,row.item.ingredients,row.item.preparation,row.item.time)"
                    >Save</b-button>
                  </b-col>
                  <b-col md="2">
                    <b-button variant="primary" @click="row.toggleDetails">Hide</b-button>
                  </b-col>
                </b-row>
              </b-form>
            </b-card>
          </b-container>
        </b-card>
      </template>
    </b-table>
    <b-modal ref="my-modal" hide-footer id="modal-1" title="Delete recipe">
      <p class="my-4">Are you sure delete {{tempName}}?</p>
      <b-button class="mt-3" variant="outline-primary" block @click="hideModal">Cancel</b-button>
      <b-button class="mt-2" variant="outline-danger" block @click="deleteRecipe">Delete</b-button>
    </b-modal>
    <b-modal ref="my-modal-success" id="modal-2" title="Success">
      <b-alert variant="success" show>
        <p class="my-4">{{message}}</p>
      </b-alert>
    </b-modal>
    <b-modal ref="my-modal-failed" id="modal-3" title="Failed">
      <b-alert variant="danger" show>
        <p class="my-4">{{message}}</p>
      </b-alert>
    </b-modal>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      numRegPerPages:10,
      numRgs: 0,
      actualPage:1,
      numPages:0,
      value: "",
      message: "",
      selected: "Name",
      options: ["ID", "Name"],
      tempName: "",
      tempId: 0,
      fields: [
        "id",
        "name",
        "ingredients",
        "preparation",
        "time",
        "edit",
        "delete"
      ],
      recipes: [],
      hola: "hola munndo!!",
      time: 0,
      form: {
        name: "",
        ingredients: "",
        preparation: "",
        time: 0
      },
      back: 11,
      next: 20,
    };
  },
  beforeCreate() {},

  created() {
    this.getCount()
    
    //console.log("Recipes: ", this.recipes)
    this.getAllRecips();
  },

  methods: {
    getRecipePage(i){
      this.actualPage = i;
      let context = this;
      
      axios({
          url: "http://localhost:8090/api/recipe/?count="+this.numRegPerPages+" &start="+((i-1)*this.numRegPerPages),
          method: "GET"
        })
          .then(function(response) {
            context.recipes = response.data;
          })
          .catch(function(error) {
            console.log("Error get count")
          });
      
    },
    getCount(){
      let context = this;
      axios({
          url: "http://localhost:8090/api/recipe/count/",
          method: "GET"
        })
          .then(function(response) {
            console.log("Get countasas", response.data.count)
            context.numRgs = response.data.count;
            context.numPages = Math.ceil(context.numRgs/context.numRegPerPages)
          })
          .catch(function(error) {
            console.log("Error get count")
          });

    },
    search() {
      let context = this;

      if (this.value) {
        axios({
          url: "http://localhost:8090/api/recipe/" + this.value.toLowerCase(),
          method: "GET"
        })
          .then(function(response) {
            context.recipes = response.data;
          })
          .catch(function(error) {
            context.message = "Recipe not found."
            context.showModal("my-modal-failed");
            contexto.callBack();
          });
      } else {
        context.message = "Insert search values"
        context.showModal("my-modal-failed");
      }
    },

    hideModal() {
      this.$refs["my-modal"].hide();
      this.tempName = "";
      this.tempId = 0;
    },
    showModal(nameModal) {
      this.$refs[nameModal].show();
    },
    confirmDelete(id, name) {
      this.tempName = name;
      this.tempId = id;

      this.$refs["my-modal"].show();
    },

    callBack() {
      this.hideModal();
      this.getAllRecips();
    },
    updateRecipe(id, name, ingredients, preparation, time) {
      let context = this;
      let recipe = {
        name: name,
        ingredients: ingredients,
        preparation: preparation,
        time: time
      };
      axios({
        headers: {
          "Access-Control-Allow-Headers":
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token",
          "Access-Control-Allow-Methods": "POST, GET, PUT, DELETE",
          "Access-Control-Allow-Origin": "*",
          "Content-Type": "application/json"
        },
        url: "http://localhost:8090/api/recipe/" + id,
        method: "PUT",
        data: recipe
      })
        .then(function(response) {
          context.message = "The edition of the recipe was success."
          context.showModal("my-modal-success");
          context.callBack();
          context.hideModal();
          console.log(response);
        })
        .catch(function(error) {
          context.message = "The edition of the recipe failed."
          context.showModal("my-modal-failed");
          context.callBack();
          console.log(error);
        });
    },
    deleteRecipe() {
      let context = this;
      axios({
        headers: {
          "Access-Control-Allow-Headers":
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token",
          "Access-Control-Allow-Methods": "POST, GET, PUT, DELETE",
          "Access-Control-Allow-Origin": "*",
          "Content-Type": "application/json"
        },
        url: "http://localhost:8090/api/recipe/" + this.tempId,
        method: "DELETE"
      })
        .then(function(response) {
          context.callBack();
        })
        .catch(function(error) {
          console.log(error);
        });
    },
    getAllRecips() {
      let currentObject = this;
      axios({
        url: "http://localhost:8090/api/recipe/",
        method: "GET"
      })
        .then(function(response) {
          currentObject.recipes = response.data;
          console.log(response);
        })
        .catch(function(error) {
          console.log(error);
        });

      console.log("resultadoooo", this.recipes);
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
