<template>
  <div>
    <vue-progress-bar></vue-progress-bar>
    <b-row class="justify-content-center">
      <b-col md="6">
        <b-input-group prepend="Zadanie" class="mt-3">
          <b-form-input v-model="task" v-on:keyup.enter="addTask"></b-form-input>
          <b-input-group-append>
            <b-button v-on:click="addTask" variant="success">
              <b-icon-plus-circle font-scale="0.97"></b-icon-plus-circle>Dodaj
            </b-button>
          </b-input-group-append>
        </b-input-group>
      </b-col>
    </b-row>
    <h4 class="mt-3">Do zrobienia</h4>
    <b-row class="justify-content-center">
      <b-col md="6">
        <b-list-group>
          <Task v-for="task in taskTodo" v-bind:key="task.id" v-bind:task="task" />
        </b-list-group>
      </b-col>
    </b-row>
    <h4 class="mt-3">Gotowe</h4>
    <b-row class="justify-content-center">
      <b-col md="6">
        <b-list-group>
          <Task v-for="task in taskDone" v-bind:key="task.id" v-bind:task="task" />
        </b-list-group>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import axios from "axios";
import Task from "./Task";

export default {
  components: {
    Task,
  },
  data() {
    return {
      tasks: [],
      task: "",
    };
  },
  computed: {
    taskTodo() {
      if (this.tasks != null) {
        return this.tasks.filter(function (task) {
          return !task.done;
        });
      } else return [];
    },
    taskDone() {
      if (this.tasks != null) {
        return this.tasks.filter(function (task) {
          return task.done;
        });
      } else return [];
    },
  },
  methods: {
    async addTask() {
      try {
        this.$Progress.start();
        await axios
          .post("http://127.0.0.1:1323/api/task/add", {
            name: this.task,
          })
          .then(() => {
            this.getTasks();
            this.task = "";
            this.$Progress.finish();
          });
      } catch (err) {
        this.$Progress.fail();
        console.log(err);
      }
    },
    async getTasks() {
      try {
        this.$Progress.start();
        let response = await axios.get("http://127.0.0.1:1323/api/tasks");
        this.tasks = response.data;
        this.$Progress.finish();
      } catch (err) {
        this.$Progress.fail();
        console.log(err);
      }
    },
    async deleteTask(task) {
      try {
        this.$Progress.start();
        await axios
          .delete(`http://127.0.0.1:1323/api/task/delete/${task.id}`)
          .then(() => {
            this.getTasks();
            this.$Progress.finish();
          });
      } catch (err) {
        this.$Progress.fail();
        console.log(err);
      }
    },
    async markTaskDone(task) {
      try {
        this.$Progress.start();
        await axios.put(`http://127.0.0.1:1323/api/task/markdone/${task.id}`).then(() => {
          this.getTasks();
          this.$Progress.finish();
        });
      } catch (err) {
        this.$Progress.fail();
        console.log(err);
      }
    },
  },
  beforeMount() {
    this.getTasks();
  },
  mounted() {
    //  [App.vue specific] When App.vue is finish loading finish the progress bar
    this.$Progress.finish();
  },
  created() {
    //  [App.vue specific] When App.vue is first loaded start the progress bar
    this.$Progress.start();
    //  hook the progress bar to start before we move router-view
  },
};
</script>
